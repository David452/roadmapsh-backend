package githubapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient 	*http.Client
	baseURL		string
}

func NewClient(baseURL string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL: baseURL,
	}
}