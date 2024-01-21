package entities

type UserRegistryItem struct {
	Id    string `xorm:"id"`
	Key   string `xorm:"key"`
	Value string `xorm:"value jsonb"`
}
