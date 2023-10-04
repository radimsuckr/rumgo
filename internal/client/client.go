// Package client is a HTTP client for crawling
package client

import (
	"io"
	"net/http"

	"github.com/EDDYCJY/fake-useragent"
)

// PageResponse holds a crawler response
type PageResponse struct {
	Content string
}

func newPageResponse(content string) PageResponse {
	return PageResponse{Content: content}
}

// SendRequest sends a request to a single URL
func SendRequest(url string) (PageResponse, error) {
	client := http.Client(*http.DefaultClient)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PageResponse{}, err
	}
	req.Header.Add("User-Agent", browser.Computer())

	resp, err := client.Do(req)
	if err != nil {
		return PageResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PageResponse{}, err
	}

	return newPageResponse(string(body)), nil
}
