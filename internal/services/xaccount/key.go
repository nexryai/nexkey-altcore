package xaccount

import (
	"lab.sda1.net/nexryai/altcore/internal/core/system"
	"lab.sda1.net/nexryai/altcore/internal/db"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
)

type KeyringService struct {
	UserId string
	KeyId  string
}

func (param *KeyringService) GetPrivateKeyPem() (string, error) {
	if param.KeyId != "" {
		panic(system.InvalidParamsOnServiceCall)
	}

	var result entities.UserKeypair

	engine, err := db.GetEngine()
	if err != nil {
		return "", err
	}

	sql := engine.Table("user_keypair")
	sql.Where("\"userId\" = ?", param.UserId)

	_, err = sql.Get(&result)
	if err != nil {
		return "", err
	}

	return result.PrivateKey, nil
}

func (param *KeyringService) GetLocalPublicKeyPem() (string, error) {
	var resultPem string

	engine, err := db.GetEngine()
	if err != nil {
		return resultPem, err
	}

	if param.KeyId != "" {
		panic(system.InvalidParamsOnServiceCall)
	}

	var result entities.UserKeypair

	// ローカルユーザーの鍵
	sql := engine.Table("user_keypair")
	sql.Where("\"userId\" = ?", param.UserId)

	_, err = sql.Get(&result)
	if err != nil {
		return resultPem, err
	} else {
		resultPem = result.PublicKey
	}

	if resultPem == "" {
		panic(system.UnexpectedEmptyString)
	}

	return resultPem, nil
}

func (param *KeyringService) GetRemotePublicKeyPem() (string, error) {
	var resultPem string

	engine, err := db.GetEngine()
	if err != nil {
		return resultPem, err
	}

	// リモートユーザーの公開鍵
	var result entities.UserPublicKey

	sql := engine.Table("user_publickey")

	if param.UserId != "" {
		sql.Where("\"userId\" = ?", param.UserId)
	} else if param.KeyId != "" {
		sql.Where("\"keyId\" = ?", param.KeyId)
	} else {
		panic(system.InvalidParamsOnServiceCall)
	}

	_, err = sql.Get(&result)
	if err != nil {
		return resultPem, err
	} else {
		resultPem = result.KeyPem
	}

	if resultPem == "" {
		panic(system.UnexpectedEmptyString)
	}

	return resultPem, nil
}
