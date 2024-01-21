package entities

type DriveFile struct {
	Id          string `gorm:"column:id;primary_key"`
	URL         string `gorm:"column:url"`
	URI         string `gorm:"column:uri"`
	Thumbnail   string `gorm:"column:thumbnailUrl"`
	Name        string `gorm:"column:name"`
	Type        string `gorm:"column:type"`
	IsSensitive bool   `gorm:"column:isSensitive"`
	BlurHash    string `gorm:"column:blurhash"`
}

func (DriveFile) TableName() string {
	return "drive_file"
}
