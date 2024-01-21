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
	ShowTimelineReplies     bool           `gorm:"column:showTimelineReplies"`
	DriveCapacityOverrideMb string         `gorm:"column:driveCapacityOverrideMb"`
	Token                   string         `gorm:"token"`
}

func (User) TableName() string {
	return "user"
}

type UserProfileFields struct {
	Name  string `xorm:"name" json:"name"`
	Value string `xorm:"value" json:"value"`
}

type UserProfileMutedWord struct {
}

type UserProfileMutedInstance struct {
}

// ToDo: なんでUserとUserProfileが別テーブルに分かれてるか謎なのでもし理由がないようならそのうち統合する
// or プロフィールとかをuserテーブルに移してリモートユーザーには関係ないカラムのみ残してlocal_userテーブルとかに改名？
type UserProfile struct {
	Id                  string              `xorm:"userId"`
	Location            string              `xorm:"location"`
	Birthday            time.Time           `xorm:"birthday"`
	Description         string              `xorm:"description"`
	Fields              []UserProfileFields `xorm:"fields"`
	Url                 string              `xorm:"url"`
	Email               string              `xorm:"email"`
	EmailVerifyCode     string              `xorm:"emailVerifyCode"`
	TwoFactorTempSecret string              `xorm:"twoFactorTempSecret"`
	TwoFactorSecret     string              `xorm:"twoFactorSecret"`
	TwoFactorEnabled    bool                `xorm:"twoFactorEnabled"`
	PasswordHash        string              `xorm:"password"`
	AutoAcceptFollowed  bool                `xorm:"autoAcceptFollowed"`
	AlwaysMarkNsfw      bool                `xorm:"alwaysMarkNsfw"`
	CarefulBot          bool                `xorm:"carefulBot"`
	UserHost            string              `xorm:"userHost"`
	EnableWordMute      bool                `xorm:"enableWordMute"`
	//MutedWords               pq.StringArray      `xorm:"mutedWords jsonb"`
	NoCrawle                 bool           `xorm:"noCrawle"`
	ReceiveAnnouncementEmail bool           `xorm:"receiveAnnouncementEmail"`
	EmailNotificationTypes   pq.StringArray `xorm:"emailNotificationTypes jsonb"`
	MutedInstances           pq.StringArray `xorm:"mutedInstances jsonb"`
	PublicReactions          bool           `xorm:"publicReactions"`
	FfVisibility             string         `xorm:"ffVisibility"`
	ModerationNote           string         `xorm:"moderationNote"`
	//TwoFactorBackupSecret    string              `xorm:"twoFactorBackupSecret"`
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
