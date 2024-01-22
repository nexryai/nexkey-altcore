package entities

import (
	"github.com/lib/pq"
	"time"
)

type User struct {
	Id                      string         `gorm:"primaryKey"`
	CreatedAt               time.Time      `gorm:"column:createdAt"`
	UpdatedAt               time.Time      `gorm:"column:updatedAt"`
	LastFetchedAt           time.Time      `gorm:"column:lastFetchedAt"`
	Username                string         `gorm:"column:username"`
	UsernameLower           string         `gorm:"column:usernameLower"`
	Name                    string         `gorm:"column:name"`
	FollowersCount          uint           `gorm:"column:followersCount"`
	FollowingCount          uint           `gorm:"column:followingCount"`
	NotesCount              uint           `gorm:"column:notesCount"`
	AvatarId                string         `gorm:"column:avatarId"`
	Avatar                  *DriveFile     `gorm:"foreignkey:Id;references:AvatarId"`
	BannerId                string         `gorm:"column:bannerId"`
	Banner                  *DriveFile     `gorm:"foreignkey:Id;references:BannerId"`
	Tags                    pq.StringArray `gorm:"column:tags;type:text[]"`
	IsSuspended             bool           `gorm:"column:isSuspended"`
	IsSilenced              bool           `gorm:"column:isSilenced"`
	IsLocked                bool           `gorm:"column:isLocked"`
	IsBot                   bool           `gorm:"column:isBot"`
	IsCat                   bool           `gorm:"column:isCat"`
	IsAdmin                 bool           `gorm:"column:isAdmin"`
	IsModerator             bool           `gorm:"column:isModerator"`
	Emojis                  pq.StringArray `gorm:"column:emojis;type:text[]"`
	Host                    string         `gorm:"host"`
	Inbox                   string         `gorm:"inbox"`
	SharedInbox             string         `gorm:"column:sharedInbox"`
	Featured                string         `gorm:"featured"`
	Uri                     string         `gorm:"uri"`
	IsExplorable            bool           `gorm:"column:isExplorable"`
	FollowersUri            string         `gorm:"column:followersUri"`
	LastActiveDate          time.Time      `gorm:"column:lastActiveDate"`
	HideOnlineStatus        bool           `gorm:"column:hideOnlineStatus"`
	IsDeleted               bool           `gorm:"column:isDeleted"`
	DriveCapacityOverrideMb string         `gorm:"column:driveCapacityOverrideMb"`
	Token                   string         `gorm:"token"`
}

func (User) TableName() string {
	return "user"
}

type UserProfileFields struct {
	Name  string `gorm:"column:name" json:"name"`
	Value string `gorm:"column:value" json:"value"`
}

type UserProfileMutedWord struct {
}

type UserProfileMutedInstance struct {
}

// ToDo: リモートユーザーには関係ないカラムをどこか別テーブルに移す
type UserProfile struct {
	Id                    string         `gorm:"column:userId;primaryKey"`
	Location              string         `gorm:"column:location;default:null"`
	Birthday              time.Time      `gorm:"column:birthday;default:null"`
	Description           string         `gorm:"column:description;default:null"`
	Url                   string         `gorm:"column:url"`
	Email                 string         `gorm:"column:email;default:null"`
	EmailVerifyCode       string         `gorm:"column:emailVerifyCode;default:null"`
	TwoFactorTempSecret   string         `gorm:"column:twoFactorTempSecret;default:null"`
	TwoFactorSecret       string         `gorm:"column:twoFactorSecret;default:null"`
	TwoFactorEnabled      bool           `gorm:"column:twoFactorEnabled"`
	PasswordHash          string         `gorm:"column:password;default:null"`
	AutoAcceptFollowed    bool           `gorm:"column:autoAcceptFollowed"`
	AlwaysMarkNsfw        bool           `gorm:"column:alwaysMarkNsfw"`
	CarefulBot            bool           `gorm:"column:carefulBot"`
	UserHost              string         `gorm:"column:userHost;default:null"`
	EnableWordMute        bool           `gorm:"column:enableWordMute"`
	MutedWords            pq.StringArray `gorm:"column:mutedWords;type:text[]"`
	NoCrawl               bool           `gorm:"column:noCrawl"`
	MutedInstances        pq.StringArray `gorm:"column:mutedInstances;type:text[]"`
	FfVisibility          string         `gorm:"column:ffVisibility"`
	ModerationNote        string         `gorm:"column:moderationNote;default:null"`
	TwoFactorBackupSecret string         `gorm:"column:twoFactorBackupSecret;default:null"`
	// Fields                   []UserProfileFields `gorm:"column:fields;type:jsonb;default:'[]';not null"`
}

type Following struct {
	Id         string    `xorm:"id"`
	CreatedAt  time.Time `xorm:"createdAt"`
	FolloweeId string    `xorm:"followeeId"`
	FollowerId string    `xorm:"followerId"`
}

type UserPublicKey struct {
	UserId string `xorm:"userId"`
	KeyId  string `xorm:"keyId"`
	KeyPem string `xorm:"keyPem"`
}
