package enum

type FollowType string

const (
	Followers FollowType = "followers"
	Followees FollowType = "followees"
)

type TimelineType string

const (
	HomeTimeline   TimelineType = "public"
	LocalTimeline  TimelineType = "local"
	SocialTimeline TimelineType = "social"
	GlobalTimeline TimelineType = "global"
)

type NoteVisibility string

const (
	NoteVisibilityPublic        NoteVisibility = "public"
	NoteVisibilityHome          NoteVisibility = "home"
	NoteVisibilityFollowersOnly NoteVisibility = "followers"
	NoteVisibilityDirect        NoteVisibility = "direct"
)
