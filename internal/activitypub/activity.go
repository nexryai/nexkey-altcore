package activitypub

import "time"

var activityContext = &[]interface{}{
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
}

type ActivityObjectSource struct {
	Content   string `json:"content"`
	MediaType string `json:"mediaType"`
}

type CreateActivityObject struct {
	Type           string               `json:"type"`
	Id             string               `json:"id"`
	Actor          string               `json:"actor"`
	AttributedTo   string               `json:"attributedTo"`
	Summary        *string              `json:"summary"`
	Content        string               `json:"content"`
	MisskeyContent string               `json:"_misskey_content"`
	Source         ActivityObjectSource `json:"source"`
	Published      time.Time            `json:"published"`
	To             []string             `json:"to"`
	Cc             []string             `json:"cc"`
	InReplyTo      *string              `json:"inReplyTo"`
	Attachment     []string             `json:"attachment"`
	Sensitive      bool                 `json:"sensitive"`
	Tag            []string             `json:"tag"`
}

type CreateActivity struct {
	Context      *[]interface{}       `json:"@context"`
	ActivityType string               `json:"type"`
	Id           string               `json:"id"`
	Actor        string               `json:"actor"`
	Published    time.Time            `json:"published"`
	Object       CreateActivityObject `json:"object"`
	To           []string             `json:"to"`
	Cc           []string             `json:"cc"`
}

type FollowActivity struct {
	Type      string    `json:"type"`
	Id        string    `json:"id"`
	Actor     string    `json:"actor"`
	Published time.Time `json:"published"`
	Object    string    `json:"object"`
}
