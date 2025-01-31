package main

import (
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(rawBaseURL, htmlBody string) ([]string, error) {
	normalizedUrls := []string{}
	htmlReader := strings.NewReader(htmlBody)
	htmlNodes, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}
	returnURLs := traverseHTML(htmlNodes)

	for _, url := range returnURLs {
		if !strings.Contains(url, "http") {
			url = rawBaseURL + url
		}
		normalURL, err := normalizeURL(url)
		if err != nil {
			return nil, err
		}
		normalizedUrls = append(normalizedUrls, "https://"+normalURL)
	}
	return normalizedUrls, nil
}

func traverseHTML(htmlNode *html.Node) []string {
	returnURLs := []string{}
	if htmlNode.Data == "a" {
		for _, htmlAttribute := range htmlNode.Attr {
			if htmlAttribute.Key == "href" {
				returnURLs = append(returnURLs, htmlAttribute.Val)
			}
		}
	}

	for htmlChild := htmlNode.FirstChild; htmlChild != nil; htmlChild = htmlChild.NextSibling {
		returnURLs = append(returnURLs, traverseHTML(htmlChild)...)
	}
	return returnURLs
}
