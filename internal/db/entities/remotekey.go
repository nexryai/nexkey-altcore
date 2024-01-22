package entities

type RemoteKey struct {
	UserId       string `gorm:"column:userId;primary_key;unique;index;not null"`
	KeyId        string `gorm:"column:keyId;unique;not null"`
	PublicKeyPem string `gorm:"column:keyPem;not null"`
}

func (RemoteKey) TableName() string {
	return "user_publickey"
}
