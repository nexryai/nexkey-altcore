package activitypub

import (
	"bytes"
	"crypto"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/go-fed/httpsig"
	"io"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/core/system"
	"net/http"
	"strings"
	"time"
)

const (
	expiry = 120
)

var (
	// http signer preferences
	prefs = []httpsig.Algorithm{httpsig.RSA_SHA256}
	// こっちからの署名に使うやつ
	digestAlgo = httpsig.DigestSha256
	// 連合先から受け入れるアルゴリズム
	acceptAlgorithms = []httpsig.Algorithm{
		httpsig.RSA_SHA256, // Prefer common RSA_SHA256.
		httpsig.RSA_SHA512, // Fall back to less common RSA_SHA512.
		httpsig.ED25519,    // Try ED25519 as a long shot.
	}
	getHeaders  = []string{httpsig.RequestTarget, "host", "date"}
	postHeaders = []string{httpsig.RequestTarget, "host", "date", "digest"}
)

func pemStringToPrivateKey(pemStr string) (crypto.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	var privateKey crypto.PrivateKey

	switch block.Type {
	case "RSA PRIVATE KEY":
		key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		privateKey = key
	case "PRIVATE KEY":
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		privateKey = key
	default:
		return nil, fmt.Errorf("unsupported key type: %s", block.Type)
	}

	return privateKey, nil
}

func pemStringToPublicKey(pemStr string) (crypto.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

type SignatureService struct {
	PrivateKeyPem string
	PublicKeyPem  string
	KeyId         string
	Request       *http.Request
}

func (params *SignatureService) Sign() error {
	privateKey, err := pemStringToPrivateKey(params.PrivateKeyPem)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if params.Request.Method == "GET" {
		getSigner, _, err := httpsig.NewSigner(prefs, digestAlgo, getHeaders, httpsig.Signature, expiry)
		err = getSigner.SignRequest(privateKey, params.KeyId, params.Request, nil)
		if err != nil {
			return err
		}
	} else if params.Request.Method == "POST" {
		requestBody, err := io.ReadAll(params.Request.Body)
		params.Request.Body = io.NopCloser(strings.NewReader(string(requestBody)))
		if err != nil {
			return err
		}

		postSigner, _, err := httpsig.NewSigner(prefs, digestAlgo, postHeaders, httpsig.Signature, expiry)
		err = postSigner.SignRequest(privateKey, params.KeyId, params.Request, requestBody)
		if err != nil {
			return err
		}
	} else {
		panic(system.InvalidParamsOnServiceCall)
	}

	return nil
}

func (params *SignatureService) Verify() bool {
	publicKey, err := pemStringToPublicKey(params.PublicKeyPem)
	if err != nil {
		logger.Error(err.Error())
		return false
	}

	if publicKey == nil {
		logger.Error("PublicKey is nil")
		return false
	}

	verifier, err := httpsig.NewVerifier(params.Request)

	// どのアルゴリズムで署名されてるか分からないのでループで回す
	for _, algo := range acceptAlgorithms {
		err := verifier.Verify(publicKey, algo)

		if err != nil {
			// だめなら次のアルゴリズムを試す
			logger.Debug("authentication NOT PASSED")
			continue
		} else {
			// 検証に成功した
			return true
		}
	}

	// ループ内で成功しないなら不正なアルゴリズムかダイジェスト
	return false
}

// Header represents a key-value pair in headers.
type Header struct {
	Name  string
	Value string
}

// Request represents an HTTP request.
type ActivityPubRequestService struct {
	Headers []Header
	Body    interface{}
	Method  string
	Url     string
}

func (params *ActivityPubRequestService) ToHttpRequest() *http.Request {
	var jsonActivity []byte
	var err error

	if params.Method == "POST" {
		jsonActivity, err = json.Marshal(params.Body)
		if err != nil {
			panic(err)
		}
	} else {
		jsonActivity = nil
	}

	var req *http.Request
	req, err = http.NewRequest(params.Method, params.Url, bytes.NewBuffer(jsonActivity))

	if err != nil {
		panic(err)
	}

	// Dateヘッダーを追加
	req.Header.Add("Date", time.Now().UTC().Format(http.TimeFormat))

	for _, header := range params.Headers {
		if strings.HasPrefix(header.Name, "(") {
			continue
		} else {
			req.Header.Add(header.Name, header.Value)
		}
	}

	req.Header.Add("accept", "application/activity+json")
	req.Header.Set("Content-Type", "application/json")

	return req
}
