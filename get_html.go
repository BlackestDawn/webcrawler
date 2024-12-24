package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	httpRes, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("error fetching from URL: %v", err)
	}
	defer httpRes.Body.Close()

	if httpRes.StatusCode >= 400 {
		return "", fmt.Errorf("error returned from webserver, statuscode: %v", httpRes.StatusCode)
	}

	contentType := httpRes.Header.Get("content-type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("unsupported content type(s): %v", contentType)
	}

	retVal, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return "", fmt.Errorf("error reading HTML: %v", err)
	}

	return string(retVal), nil
}
