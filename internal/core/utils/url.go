package utils

import "net/url"

func GetHostFromUrl(urlStr string) string {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	return parsedURL.Host
}
