package slack

import (
	"net"
	"net/url"
	"os"
)

var portMapping = map[string]string{"ws": "80", "wss": "443"}

func init() {
	if port := os.Getenv("PORT"); port != "" {
		portMapping["ws"] = port
	}

	if port := os.Getenv("PORT"); port != "" {
		portMapping["wss"] = port
	}
}

func websocketizeURLPort(orig string) (string, error) {
	urlObj, err := url.ParseRequestURI(orig)
	if err != nil {
		return "", err
	}
	_, _, err = net.SplitHostPort(urlObj.Host)
	if err != nil {
		return urlObj.Scheme + "://" + urlObj.Host + ":" + portMapping[urlObj.Scheme] + urlObj.Path, nil
	}
	return orig, nil
}
