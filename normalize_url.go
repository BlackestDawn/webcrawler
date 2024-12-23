package main

import (
	"net/url"
	"strings"
)

func normalizeURL(input string) (string, error) {
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}
	return strings.ToLower(strings.TrimRight(u.Hostname()+u.Path, "/")), nil
}
