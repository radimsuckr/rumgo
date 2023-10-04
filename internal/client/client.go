package client

import (
	"io"
	"net/http"

	"github.com/EDDYCJY/fake-useragent"
)

type PageResponse struct {
	Content string
}

func NewPageResponse(content string) PageResponse {
	return PageResponse{Content: content}
}

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

	return NewPageResponse(string(body)), nil
}
