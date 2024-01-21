package schema

type Meta struct {
	MaintainerName               string   `json:"maintainerName"`
	MaintainerEmail              string   `json:"maintainerEmail"`
	Version                      string   `json:"version"`
	Name                         string   `json:"name"`
	Uri                          string   `json:"uri"`
	Description                  string   `json:"description"`
	Langs                        []string `json:"langs"`
	TosUrl                       string   `json:"tosUrl"`
	RepositoryUrl                string   `json:"repositoryUrl"`
	FeedbackUrl                  string   `json:"feedbackUrl"`
	DefaultDarkTheme             string   `json:"defaultDarkTheme"`
	DefaultLightTheme            string   `json:"defaultLightTheme"`
	DisableRegistration          bool     `json:"disableRegistration"`
	DisableLocalTimeline         bool     `json:"disableLocalTimeline"`
	DisableGlobalTimeline        bool     `json:"disableGlobalTimeline"`
	DriveCapacityPerLocalUserMb  float64  `json:"driveCapacityPerLocalUserMb"`
	DriveCapacityPerRemoteUserMb float64  `json:"driveCapacityPerRemoteUserMb"`
	CacheRemoteFiles             bool     `json:"cacheRemoteFiles"`
	EmailRequiredForSignup       bool     `json:"emailRequiredForSignup"`
	EnableHcaptcha               bool     `json:"enableHcaptcha"`
	HcaptchaSiteKey              string   `json:"hcaptchaSiteKey"`
	EnableRecaptcha              bool     `json:"enableRecaptcha"`
	RecaptchaSiteKey             string   `json:"recaptchaSiteKey"`
	EnableTurnstile              bool     `json:"enableTurnstile"`
	TurnstileSiteKey             string   `json:"turnstileSiteKey"`
	SwPublickey                  string   `json:"swPublickey"`
	ThemeColor                   string   `json:"themeColor"`
	MascotImageUrl               string   `json:"mascotImageUrl"`
	BannerUrl                    string   `json:"bannerUrl"`
	ErrorImageUrl                string   `json:"errorImageUrl"`
	IconUrl                      string   `json:"iconUrl"`
	BackgroundImageUrl           string   `json:"backgroundImageUrl"`
	LogoImageUrl                 string   `json:"logoImageUrl"`
	MaxNoteTextLength            float64  `json:"maxNoteTextLength"`
	Emojis                       []Emoji  `json:"emojis"`
	Ads                          []Ad     `json:"ads"`
	RequireSetup                 bool     `json:"requireSetup"`
	EnableEmail                  bool     `json:"enableEmail"`
	EnableTwitterIntegration     bool     `json:"enableTwitterIntegration"`
	EnableGithubIntegration      bool     `json:"enableGithubIntegration"`
	EnableDiscordIntegration     bool     `json:"enableDiscordIntegration"`
	EnableServiceWorker          bool     `json:"enableServiceWorker"`
	TranslatorAvailable          bool     `json:"translatorAvailable"`
	ProxyAccountName             string   `json:"proxyAccountName"`
	Features                     Features `json:"features"`
}

type Emoji struct {
	Id       string   `json:"id"`
	Aliases  []string `json:"aliases"`
	Category string   `json:"category"`
	Host     string   `json:"host"`
	Url      string   `json:"url"`
}

type Ad struct {
	Place    string `json:"place"`
	Url      string `json:"url"`
	ImageUrl string `json:"imageUrl"`
}

type Features struct {
	Registration   bool `json:"registration"`
	LocalTimeLine  bool `json:"localTimeLine"`
	GlobalTimeLine bool `json:"globalTimeLine"`
	Elasticsearch  bool `json:"elasticsearch"`
	Hcaptcha       bool `json:"hcaptcha"`
	Recaptcha      bool `json:"recaptcha"`
	ObjectStorage  bool `json:"objectStorage"`
	Twitter        bool `json:"twitter"`
	Github         bool `json:"github"`
	Discord        bool `json:"discord"`
	ServiceWorker  bool `json:"serviceWorker"`
	Miauth         bool `json:"miauth"`
}
