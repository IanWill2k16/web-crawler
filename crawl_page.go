package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) error {
	fmt.Println("Now crawling page: ", rawCurrentURL)
	currDomain, err := url.Parse(rawCurrentURL)
	if err != nil {
		return err
	}
	baseDomain, err := url.Parse(rawBaseURL)
	if err != nil {
		return err
	}

	if currDomain.Host != baseDomain.Host {
		return nil
	}

	normalCurrentUrl, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return err
	}

	if pages[normalCurrentUrl] != 0 {
		pages[normalCurrentUrl]++
	} else {
		pages[normalCurrentUrl] = 1
	}

	currHtml, err := getHTML(rawCurrentURL)
	if err != nil {
		return err
	}

	futureURLs, err := getURLsFromHTML(rawBaseURL, currHtml)
	if err != nil {
		return err
	}

	for _, url := range futureURLs {
		crawlPage(rawBaseURL, url, pages)
	}

	return nil
}
