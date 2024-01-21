package config

type ServerConfig struct {
	URL    string `yaml:"url"`
	Port   int    `yaml:"port"`
	Client string `yaml:"client"`
	DB     struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		DB           string `yaml:"db"`
		User         string `yaml:"user"`
		Pass         string `yaml:"pass"`
		DisableCache bool   `yaml:"disableCache"`
		Extra        struct {
			SSL bool `yaml:"ssl"`
		} `yaml:"extra"`
	} `yaml:"db"`
	Redis struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		Family int    `yaml:"family"`
		Pass   string `yaml:"pass"`
		Prefix string `yaml:"prefix"`
		DB     int    `yaml:"db"`
	} `yaml:"redis"`
	Sonic struct {
		Host       string `yaml:"host"`
		Port       int    `yaml:"port"`
		Auth       string `yaml:"auth"`
		Collection string `yaml:"collection"`
		Bucket     string `yaml:"bucket"`
	} `yaml:"sonic"`
	ID                     string   `yaml:"id"`
	Secret                 string   `yaml:"secret"`
	DisableAntenna         bool     `yaml:"disableAntenna"`
	DisableSearch          bool     `yaml:"disableSearch"`
	Proxy                  string   `yaml:"proxy"`
	ProxyBypassHosts       []string `yaml:"proxyBypassHosts"`
	ProxySmtp              string   `yaml:"proxySmtp"`
	MediaProxy             string   `yaml:"mediaProxy"`
	AllowedPrivateNetworks []string `yaml:"allowedPrivateNetworks"`
	MaxFileSize            int64    `yaml:"maxFileSize"`
	QueueConcurrency       int      `yaml:"queueConcurrency"`
}
