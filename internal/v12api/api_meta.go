package v12api

import (
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/instance"
	"lab.sda1.net/nexryai/altcore/internal/v12api/schema"
)

func GetMeta() *schema.Meta {
	m := instance.GetInstanceMeta()

	return &schema.Meta{
		MaintainerName:  m.MaintainerName,
		MaintainerEmail: m.MaintainerEmail,

		Version: "12.24Q1.1",

		Name:                         m.Name,
		Uri:                          config.URL,
		Description:                  m.Description,
		TosUrl:                       m.ToSURL,
		RepositoryUrl:                "https://github.com/nexryai/nexkey",
		FeedbackUrl:                  "https://github.com/nexryai/nexkey/issues",
		DisableRegistration:          m.DisableRegistration,
		DisableLocalTimeline:         m.DisableLocalTimeline,
		DisableGlobalTimeline:        m.DisableGlobalTimeline,
		DriveCapacityPerLocalUserMb:  float64(m.LocalDriveCapacityMB),
		DriveCapacityPerRemoteUserMb: float64(m.RemoteDriveCapacityMB),
		EmailRequiredForSignup:       m.EmailRequiredForSignup,
		EnableHcaptcha:               m.EnableHcaptcha,
		EnableRecaptcha:              m.EnableRecaptcha,
		EnableTurnstile:              m.EnableTurnstile,
		SwPublickey:                  m.SWPublicKey,
		ThemeColor:                   m.ThemeColor,
		MascotImageUrl:               m.MascotImageURL,
		BannerUrl:                    m.BannerURL,
		ErrorImageUrl:                m.ErrorImageUrl,
		IconUrl:                      m.IconUrl,
		BackgroundImageUrl:           m.BackgroundImageURL,
		LogoImageUrl:                 m.LogoImageUrl,
		MaxNoteTextLength:            3000,
		Emojis:                       []schema.Emoji{
			// ToDo
		},
		DefaultDarkTheme:    m.DefaultDarkTheme,
		DefaultLightTheme:   m.DefaultLightTheme,
		EnableEmail:         m.EnableEmail,
		EnableServiceWorker: m.EnableServiceWorker,
		TranslatorAvailable: m.DeeplAuthKey != "",
		// ToDo
		RequireSetup: false,
	}
}
