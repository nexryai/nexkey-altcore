package schema

type User struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	Host           string `json:"host"`
	AvatarUrl      string `json:"avatarUrl"`
	AvatarBlurhash string `json:"avatarBlurhash"`
	IsBot          bool   `json:"isBot"`
	IsCat          bool   `json:"isCat"`
}
