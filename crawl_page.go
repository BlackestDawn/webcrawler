package main

import (
	"fmt"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	if !strings.Contains(rawCurrentURL, rawBaseURL) {
		fmt.Println("going outside domain:", rawBaseURL)
		return
	}

	normURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("error normalizing URL:", err)
		return
	}

	if _, ok := pages[normURL]; ok {
		fmt.Println("duplicate entry:", normURL)
		pages[normURL]++
		return
	} else {
		fmt.Println("new entry:", normURL)
		pages[normURL] = 1
	}

	fmt.Println("fetching HTML from:", rawCurrentURL)
	htmlDoc, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("error fetching HTML:", err)
		return
	}

	urls, err := getURLsFromHTML(htmlDoc, rawBaseURL)
	if err != nil {
		fmt.Println("error getting URLs from HTML:", err)
	}

	for _, u := range urls {
		crawlPage(rawBaseURL, u, pages)
	}
}
