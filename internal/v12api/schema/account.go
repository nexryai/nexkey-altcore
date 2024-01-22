package schema

import "time"

type MyAccount struct {
	Id                              string    `json:"id"`
	Name                            string    `json:"name"`
	Username                        string    `json:"username"`
	Host                            *string   `json:"host"`
	AvatarURL                       string    `json:"avatarUrl,omitempty"`
	AvatarBlurhash                  string    `json:"avatarBlurhash,omitempty"`
	AvatarColor                     string    `json:"avatarColor,omitempty"`
	IsAdmin                         bool      `json:"isAdmin"`
	IsModerator                     bool      `json:"isModerator"`
	IsBot                           bool      `json:"isBot"`
	IsCat                           bool      `json:"isCat"`
	Emojis                          []string  `json:"emojis"`
	OnlineStatus                    string    `json:"onlineStatus"`
	DriveCapacityOverrideMb         int       `json:"driveCapacityOverrideMb"`
	URL                             string    `json:"url"`
	URI                             string    `json:"uri"`
	CreatedAt                       time.Time `json:"createdAt"`
	UpdatedAt                       time.Time `json:"updatedAt"`
	LastFetchedAt                   time.Time `json:"lastFetchedAt"`
	BannerUrl                       string    `json:"bannerUrl,omitempty"`
	BannerBlurhash                  string    `json:"bannerBlurhash,omitempty"`
	BannerColor                     string    `json:"bannerColor,omitempty"`
	IsLocked                        bool      `json:"isLocked"`
	IsSilenced                      bool      `json:"isSilenced"`
	IsSuspended                     bool      `json:"isSuspended"`
	Description                     string    `json:"description"`
	Location                        string    `json:"location,omitempty"`
	Birthday                        string    `json:"birthday,omitempty"`
	Lang                            string    `json:"lang,omitempty"`
	Fields                          []string  `json:"fields"`
	FollowersCount                  int       `json:"followersCount"`
	FollowingCount                  int       `json:"followingCount"`
	NotesCount                      int       `json:"notesCount"`
	PinnedNoteIds                   []string  `json:"pinnedNoteIds"`
	PinnedNotes                     []string  `json:"pinnedNotes"`
	PinnedPageId                    string    `json:"pinnedPageId"`
	PinnedPage                      string    `json:"pinnedPage"`
	PublicReactions                 bool      `json:"publicReactions"`
	FFVisibility                    string    `json:"ffVisibility"`
	TwoFactorEnabled                bool      `json:"twoFactorEnabled"`
	UsePasswordLessLogin            bool      `json:"usePasswordLessLogin"`
	SecurityKeys                    bool      `json:"securityKeys"`
	AvatarId                        string    `json:"avatarId"`
	BannerId                        string    `json:"bannerId"`
	InjectFeaturedNote              bool      `json:"injectFeaturedNote"`
	ReceiveAnnouncementEmail        bool      `json:"receiveAnnouncementEmail"`
	AlwaysMarkNsfw                  bool      `json:"alwaysMarkNsfw"`
	CarefulBot                      bool      `json:"carefulBot"`
	AutoAcceptFollowed              bool      `json:"autoAcceptFollowed"`
	NoCrawle                        bool      `json:"noCrawle"`
	IsExplorable                    bool      `json:"isExplorable"`
	IsDeleted                       bool      `json:"isDeleted"`
	TwoFactorBackupCodesStock       string    `json:"twoFactorBackupCodesStock"`
	HideOnlineStatus                bool      `json:"hideOnlineStatus"`
	HasUnreadSpecifiedNotes         bool      `json:"hasUnreadSpecifiedNotes"`
	HasUnreadMentions               bool      `json:"hasUnreadMentions"`
	HasUnreadAnnouncement           bool      `json:"hasUnreadAnnouncement"`
	HasUnreadAntenna                bool      `json:"hasUnreadAntenna"`
	HasUnreadChannel                bool      `json:"hasUnreadChannel"`
	HasUnreadMessagingMessage       bool      `json:"hasUnreadMessagingMessage"`
	HasUnreadNotification           bool      `json:"hasUnreadNotification"`
	HasPendingReceivedFollowRequest bool      `json:"hasPendingReceivedFollowRequest"`
	MutedWords                      []string  `json:"mutedWords"`
	MutedInstances                  []string  `json:"mutedInstances"`
	MutingNotificationTypes         []string  `json:"mutingNotificationTypes"`
	EmailNotificationTypes          []string  `json:"emailNotificationTypes"`
	ShowTimelineReplies             bool      `json:"showTimelineReplies"`
	Email                           string    `json:"email"`
	EmailVerified                   bool      `json:"emailVerified"`
	SecurityKeysList                []string  `json:"securityKeysList"`
}