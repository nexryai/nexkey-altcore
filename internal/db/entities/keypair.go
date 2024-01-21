package entities

type UserKeypair struct {
	UserId     string `xorm:"userId"`
	PublicKey  string `xorm:"publicKey"`
	PrivateKey string `xorm:"privateKey"`
}
