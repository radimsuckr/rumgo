package client

import (
	"io"
	"net/http"

	"github.com/EDDYCJY/fake-useragent"
)

type PageResponse struct {
	Content string
}

func NewPageResponse(content string) *PageResponse {
	return &PageResponse{Content: content}
}

func SendRequest(url string) (*PageResponse, error) {
	client := http.Client(*http.DefaultClient)
	req, req_err := http.NewRequest("GET", url, nil)
	if req_err != nil {
		return nil, req_err
	}
	req.Header.Add("User-Agent", browser.Computer())

	resp, req_err := client.Do(req)
	if req_err != nil {
		return nil, req_err
	}
	defer resp.Body.Close()

	body, body_err := io.ReadAll(resp.Body)
	if body_err != nil {
		return nil, body_err
	}

	return NewPageResponse(string(body)), nil
}
