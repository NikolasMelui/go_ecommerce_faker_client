package efclient

import (
	"net/http"
	"time"
)

// Client ...
type Client struct {
	config     *Config
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient ...
func NewClient(config *Config) *Client {
	return &Client{
		BaseURL: config.BaseURL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
