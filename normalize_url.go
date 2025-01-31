package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}
	returnURL := parsedURL.Host + parsedURL.Path
	returnURL = strings.TrimSuffix(returnURL, "/")
	return returnURL, nil
}
