package config

import (
	"gopkg.in/yaml.v3"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"net/url"
	"os"
	"strings"
)

var (
	serverConfig     = loadConfigFromFile()
	Id               = serverConfig.ID
	Secret           = serverConfig.Secret
	URL              = serverConfig.URL
	Port             = serverConfig.Port
	Client           = serverConfig.Client
	DB               = serverConfig.DB
	Redis            = serverConfig.Redis
	Sonic            = serverConfig.Sonic
	Host             = getHostFromUrl(serverConfig.URL)
	QueueConcurrency = serverConfig.QueueConcurrency
)

func getHostFromUrl(urlStr string) string {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	return parsedURL.Host
}

func loadConfigFromFile() *ServerConfig {
	// ToDO: Redisキャッシュするようにする
	file, err := os.ReadFile(".config/default.yml")
	if err != nil {
		logger.FatalWithDetail("Error reading the configuration file", err)
		panic(err)
	}

	var config ServerConfig
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		logger.FatalWithDetail("Error unmarshalling YAML", err)
		panic(err)
	}

	// スラッシュがあれば取り除く
	if strings.HasSuffix(config.URL, "/") {
		config.URL = strings.TrimSuffix(config.URL, "/")
	}

	if strings.HasSuffix(config.Client, "/") {
		config.Client = strings.TrimSuffix(config.Client, "/")
	}

	return &config
}
