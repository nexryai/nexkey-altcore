package activitypub

import (
	"fmt"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/core/system"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
	"lab.sda1.net/nexryai/altcore/internal/services/baselib"
	"lab.sda1.net/nexryai/altcore/internal/services/xaccount"
	"lab.sda1.net/nexryai/altcore/internal/services/xdrive"
	"time"
)

type PublicKey struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}

type PropertyValue struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Attachment struct {
	Type  string `json:"type"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Image struct {
	Type      string `json:"type"`
	URL       string `json:"url"`
	Sensitive bool   `json:"sensitive"`
	Name      string `json:"name,omitempty"`
}

type Endpoints struct {
	SharedInbox string `json:"sharedInbox"`
}

type Hashtag struct {
	Type string `json:"type"`
	Href string `json:"href"`
	Name string `json:"name"`
}

type Person struct {
	Context                   []interface{}   `json:"@context"`
	Type                      string          `json:"type"`
	Id                        string          `json:"id"`
	Inbox                     string          `json:"inbox"`
	Outbox                    string          `json:"outbox"`
	Followers                 string          `json:"followers"`
	Following                 string          `json:"following"`
	Featured                  string          `json:"featured"`
	SharedInbox               string          `json:"sharedInbox"`
	Endpoints                 Endpoints       `json:"endpoints"`
	URL                       string          `json:"url"`
	PreferredUsername         string          `json:"preferredUsername"`
	Name                      string          `json:"name"`
	Summary                   string          `json:"summary"`
	Icon                      Image           `json:"icon"`
	Image                     Image           `json:"image"`
	Tags                      []Hashtag       `json:"tag"`
	ManuallyApprovesFollowers bool            `json:"manuallyApprovesFollowers"`
	Discoverable              bool            `json:"discoverable"`
	Published                 time.Time       `json:"published"`
	PublicKey                 PublicKey       `json:"publicKey"`
	IsCat                     bool            `json:"isCat"`
	Attachment                []PropertyValue `json:"attachment"`
	Birthday                  string          `json:"vcard:bday"`
}

func RenderPerson(userId string) (*Person, error) {
	userService := baselib.UserService{
		LocalOnly: true,
	}

	userInfo, err := userService.FindOne(userId)
	if err != nil {
		return &Person{}, err
	}

	if userInfo.IsSuspended {
		return &Person{}, system.UserSuspended
	}

	userProfile, err := userService.GetProfile(userId)
	if err != nil {
		return &Person{}, err
	}

	userUrl := fmt.Sprintf("%s/users/%s", config.URL, userInfo.Id)

	keyringService := xaccount.KeyringService{
		UserId: userId,
	}

	userPublicKey, err := keyringService.GetLocalPublicKeyPem()
	if err != nil {
		logger.ErrorWithDetail("Failed to get userPublicKey", err)
		return &Person{}, err
	}

	var userAvatar entities.DriveFile
	if userInfo.AvatarId != "" {
		driveService := xdrive.DriveService{
			FileId:    userInfo.AvatarId,
			LocalOnly: true,
		}
		userAvatar, err = driveService.FindOne()
		if err != nil {
			return &Person{}, err
		}
	}

	var userBanner entities.DriveFile
	if userInfo.BannerId != "" {
		driveService := xdrive.DriveService{
			FileId:    userInfo.BannerId,
			LocalOnly: true,
		}
		userBanner, err = driveService.FindOne()
		if err != nil {
			return &Person{}, err
		}
	}

	return &Person{
		Context: []interface{}{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
			map[string]string{
				"manuallyApprovesFollowers": "as:manuallyApprovesFollowers",
				"sensitive":                 "as:sensitive",
				"Hashtag":                   "as:Hashtag",
				"quoteUrl":                  "as:quoteUrl",
				"toot":                      "http://joinmastodon.org/ns#",
				"Emoji":                     "toot:Emoji",
				"featured":                  "toot:featured",
				"discoverable":              "toot:discoverable",
				"schema":                    "http://schema.org#",
				"PropertyValue":             "schema:PropertyValue",
				"value":                     "schema:value",
				"misskey":                   "https://misskey-hub.net/ns#",
				"_misskey_content":          "misskey:_misskey_content",
				"_misskey_quote":            "misskey:_misskey_quote",
				"_misskey_reaction":         "misskey:_misskey_reaction",
				"_misskey_votes":            "misskey:_misskey_votes",
				"_misskey_talk":             "misskey:_misskey_talk",
				"isCat":                     "misskey:isCat",
				"vcard":                     "http://www.w3.org/2006/vcard/ns#",
			},
		},
		Type:        "Person",
		Id:          userUrl,
		Inbox:       fmt.Sprintf("%s/inbox", userUrl),
		Outbox:      fmt.Sprintf("%s/outbox", userUrl),
		Followers:   fmt.Sprintf("%s/followers", userUrl),
		Following:   fmt.Sprintf("%s/following", userUrl),
		Featured:    fmt.Sprintf("%s/collections/featured", userUrl),
		SharedInbox: fmt.Sprintf("%s/inbox", config.URL),
		Endpoints: Endpoints{
			SharedInbox: fmt.Sprintf("%s/inbox", config.URL),
		},
		URL:               fmt.Sprintf("%s/@%s", config.URL, userInfo.Username),
		PreferredUsername: userInfo.UsernameLower,
		Name:              userInfo.Name,
		Summary:           userProfile.Description,
		Icon: Image{
			URL:       userAvatar.URL,
			Type:      userAvatar.Type,
			Sensitive: userAvatar.IsSensitive,
			Name:      "avatar.bin",
		},
		Image: Image{
			URL:       userBanner.URL,
			Type:      userBanner.Type,
			Sensitive: userBanner.IsSensitive,
			Name:      "banner.bin",
		},
		ManuallyApprovesFollowers: userInfo.IsLocked,
		Published:                 userInfo.CreatedAt,
		IsCat:                     userInfo.IsCat,
		PublicKey: PublicKey{
			Type:         "Key",
			ID:           fmt.Sprintf("%s#main-key", userUrl),
			Owner:        userUrl,
			PublicKeyPem: userPublicKey,
		},
	}, nil
}
