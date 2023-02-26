package client

import (
	"io"
	"net/http"
)

type PageResponse struct {
	Content string
}

func NewPageResponse(content string) *PageResponse {
	return &PageResponse{Content: content}
}

func SendRequest(url string) (*PageResponse, error) {
	client := http.Client(*http.DefaultClient)
	resp, req_err := client.Get(url)
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
