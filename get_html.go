package main

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("Response failed with status code: %d and body: %s\n", res.StatusCode, body)
	}

	contentHeader := res.Header.Get("Content-Type")

	if !strings.Contains(contentHeader, "text/html") {
		return "", fmt.Errorf("Content header is not text/html")
	}

	htmlString := html.UnescapeString(string(body))

	return htmlString, nil
}
