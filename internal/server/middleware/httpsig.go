package middleware

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/go-fed/httpsig"
	"github.com/gofiber/fiber/v2"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"strings"
)

// WIP request-targetを読むようにする
func verifyHttpsigSignature(request *fiber.Ctx, publicKey *rsa.PublicKey) error {
	sigHeader := string(httpsig.Signature)
	authHeader := string(httpsig.Authorization)

	// 署名ヘッダーの取得
	signatureHeader := request.Get(sigHeader)
	if signatureHeader == "" {
		// フォールバック
		signatureHeader = request.Get(authHeader)
		if signatureHeader == "" {
			return fmt.Errorf("Signature header not found")
		}
	}

	// 署名ヘッダーの解析
	signatureParts := strings.Split(signatureHeader, ",")
	signatureParams := make(map[string]string)
	for _, part := range signatureParts {
		keyValue := strings.SplitN(part, "=", 2)
		if len(keyValue) == 2 {
			signatureParams[strings.TrimSpace(keyValue[0])] = strings.Trim(keyValue[1], ` "`)
		}
	}

	// 署名データの取得
	signatureData := signatureParams["signature"]

	// データを構築
	dataToSign := fmt.Sprintf("(request-target): %s %s\nhost: %s", request.Method(), request.Path(), request.Hostname())

	// 署名データをBase64でデコード
	decodedSignature, err := base64.StdEncoding.DecodeString(signatureData)
	if err != nil {
		return fmt.Errorf("Failed to decode signature")
	}

	// 署名検証
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, []byte(dataToSign), decodedSignature)
	if err != nil {
		// SHA512にフォールバック
		err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, []byte(dataToSign), decodedSignature)
		if err != nil {
			return fmt.Errorf("Signature verification failed: %v", err)
		}
	}

	return nil
}

func VerifySignature(ctx *fiber.Ctx) error {
	if ctx.Get("host") != config.Host {
		logger.Warn("invalid host header")
		return ctx.SendStatus(400)
	}

	ctx.Request()

	// デバッグ用
	// ここに本来ならDBから鍵を取ってくるかリモートから取ってくる処理を書く
	privateKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	pubKey := &privateKey.PublicKey

	err := verifyHttpsigSignature(ctx, pubKey)
	if err != nil {
		return ctx.SendStatus(401)
	}

	return ctx.Next()

}
