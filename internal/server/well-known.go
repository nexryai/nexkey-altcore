package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"lab.sda1.net/nexryai/altcore/internal/activitypub"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/instance"
	"lab.sda1.net/nexryai/altcore/internal/services/baselib"
)

type nodeSoftware struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Repository string `json:"repository"`
}

type nodeMaintainer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type nodeMetadata struct {
	NodeName                 string         `json:"nodeName"`
	NodeDescription          string         `json:"nodeDescription"`
	Maintainer               nodeMaintainer `json:"maintainer"`
	Langs                    []string       `json:"langs"`
	TosUrl                   string         `json:"tosUrl"`
	RepositoryUrl            string         `json:"repositoryUrl"`
	FeedbackUrl              string         `json:"feedbackUrl"`
	DisableRegistration      bool           `json:"disableRegistration"`
	DisableLocalTimeline     bool           `json:"disableLocalTimeline"`
	DisableGlobalTimeline    bool           `json:"disableGlobalTimeline"`
	EmailRequiredForSignup   bool           `json:"emailRequiredForSignup"`
	EnableHcaptcha           bool           `json:"enableHcaptcha"`
	EnableRecaptcha          bool           `json:"enableRecaptcha"`
	EnableTwitterIntegration bool           `json:"enableTwitterIntegration"`
	EnableGithubIntegration  bool           `json:"enableGithubIntegration"`
	EnableDiscordIntegration bool           `json:"enableDiscordIntegration"`
	EnableEmail              bool           `json:"enableEmail"`
	EnableServiceWorker      bool           `json:"enableServiceWorker"`
	ProxyAccountName         string         `json:"proxyAccountName"`
	ThemeColor               string         `json:"themeColor"`
}

type nodeUsage struct {
	Users struct {
		Total          int `json:"total"`
		ActiveHalfyear int `json:"activeHalfyear"`
		ActiveMonth    int `json:"activeMonth"`
	} `json:"users"`
	LocalPosts    int `json:"localPosts"`
	LocalComments int `json:"localComments"`
}

type nodeServices struct {
	Inbound  []string `json:"inbound"`
	Outbound []string `json:"outbound"`
}

type nodeInfo struct {
	Version           string       `json:"version"`
	Software          nodeSoftware `json:"software"`
	Protocols         []string     `json:"protocols"`
	Services          nodeServices `json:"services"`
	OpenRegistrations bool         `json:"openRegistrations"`
	Usage             nodeUsage    `json:"usage"`
	Metadata          nodeMetadata `json:"metadata"`
}

func escapeAttribute(value string) string {
	// ToDo?
	return value
}

func escapeValue(value string) string {
	// ToDo?
	return value
}

func XRD(x []map[string]interface{}) string {
	result := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><XRD xmlns=\"http://docs.oasis-open.org/ns/xri/xrd-1.0\">"
	for _, element := range x {
		result += "<" + element["element"].(string)

		if attributes, ok := element["attributes"].(map[string]string); ok {
			for k, v := range attributes {
				result += " " + k + "=\"" + escapeAttribute(v) + "\""
			}
		}

		if value, ok := element["value"].(string); ok {
			result += ">" + escapeValue(value) + "</" + element["element"].(string) + ">"
		} else {
			result += "/>"
		}
	}
	result += "</XRD>"
	return result
}

func MkWellKnownRouter(app *fiber.App) {
	allPath := "/.well-known/(.*)"
	webFingerPath := "/.well-known/webfinger"
	jrd := "application/jrd+json"
	xrd := "application/xrd+xml"

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowHeaders:  "Accept",
		AllowMethods:  "GET, OPTIONS",
		AllowOrigins:  "*",
		ExposeHeaders: "Vary",
	}))

	// OPTIONS handler
	app.Options(allPath, func(ctx *fiber.Ctx) error {
		ctx.Status(fiber.StatusNoContent)
		return nil
	})

	// /.well-known/host-instance
	// Mastodonがこれないと認識しないらしい
	app.Get("/.well-known/host-instance", func(ctx *fiber.Ctx) error {
		xrdResponse := XRD([]map[string]interface{}{
			{"element": "Link", "attributes": map[string]string{
				"rel":      "lrdd",
				"type":     xrd,
				"template": "CHANGEME" + webFingerPath + "?resource={uri}",
			}},
		})
		ctx.Set("Content-Type", xrd)
		return ctx.SendString(xrdResponse)
	})

	// /.well-known/host-instance.json
	app.Get("/.well-known/host-instance.json", func(ctx *fiber.Ctx) error {
		jrdResponse := map[string]interface{}{
			"links": []map[string]interface{}{
				{
					"rel":      "lrdd",
					"type":     jrd,
					"template": config.URL + webFingerPath + "?resource={uri}",
				},
			},
		}
		ctx.Set("Content-Type", "application/json")
		return ctx.JSON(jrdResponse)
	})

	// /.well-known/nodeinfo
	app.Get("/.well-known/nodeinfo/:version", func(ctx *fiber.Ctx) error {
		meta := instance.GetInstanceMeta()

		version := ctx.Params("version")
		if version != "2.0" && version != "2.1" {
			return ctx.SendStatus(404)
		}

		return ctx.JSON(nodeInfo{
			Version:   version,
			Protocols: []string{"activitypub"},
			Software: nodeSoftware{
				Name:       "nexkey-codename-altcore (Go backend)",
				Version:    "0.01",
				Repository: "https://github.com/nexryai/nexkey",
			},
			// これ要る？
			Services: nodeServices{
				Inbound:  []string{},
				Outbound: []string{"atom1.0", "rss2.0"},
			},
			Metadata: nodeMetadata{
				NodeName:        meta.Name,
				NodeDescription: meta.Description,
				Maintainer: nodeMaintainer{
					Name:  meta.MaintainerName,
					Email: meta.MaintainerEmail,
				},
				TosUrl:                 meta.ToSURL,
				RepositoryUrl:          "https://github.com/nexryai/nexkey",
				FeedbackUrl:            "https://github.com/nexryai/nexkey/issues",
				DisableRegistration:    meta.DisableRegistration,
				DisableLocalTimeline:   meta.DisableLocalTimeline,
				DisableGlobalTimeline:  meta.DisableGlobalTimeline,
				EmailRequiredForSignup: meta.EmailRequiredForSignup,
				EnableHcaptcha:         meta.EnableHcaptcha,
				EnableRecaptcha:        meta.EnableRecaptcha,
				EnableEmail:            meta.EnableEmail,
				ProxyAccountName:       meta.ProxyAccountID,
				ThemeColor:             meta.ThemeColor,
			},
		})
	})

	// /.well-known/webfinger
	app.Get("/.well-known/webfinger", func(ctx *fiber.Ctx) error {
		subjectQuery := ctx.Query("resource")
		username, err := activitypub.GetUserNameFromSubject(subjectQuery)
		if err != nil {
			return ctx.SendStatus(400)
		}

		userService := baselib.UserService{
			LocalOnly: true,
		}
		u, err := userService.FindOneByName(username)
		if err != nil {
			return ctx.SendStatus(500)
		} else if u.Id == "" {
			return ctx.SendStatus(404)
		}

		jrdResponse := map[string]interface{}{
			"subject": subjectQuery,
			"links": []map[string]interface{}{
				{
					"rel":  "self",
					"type": "application/activity+json",
					"href": fmt.Sprintf("%s/users/%s", config.URL, u.Id),
				},
				{
					"rel":  "https://webfinger.net/rel/profile-page",
					"type": "text/html",
					"href": fmt.Sprintf("%s/@%s", config.URL, u.Username),
				},
				{
					"rel":      "http://ostatus.org/schema/1.0/subscribe",
					"type":     "text/html",
					"template": fmt.Sprintf("%s/authorize-follow?acct={uri}", config.URL),
				},
			},
		}
		ctx.Set("Content-Type", "application/json")
		return ctx.JSON(jrdResponse)
	})
}
