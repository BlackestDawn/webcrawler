package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}

	htmlReader := strings.NewReader(htmlBody)
	htmlNodes, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("error parsing html: %v", err)
	}

	var retVal []string
	for n := range htmlNodes.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					url, err := url.Parse(a.Val)
					if err != nil {
						fmt.Printf("error parsing URL '%v': %v\n", a.Val, err)
						continue
					}

					resURL := baseURL.ResolveReference(url)
					retVal = append(retVal, resURL.String())
				}
			}
		}
	}

	return retVal, nil
}
