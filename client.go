package alphasoc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	APIBase        = "https://api.alphasoc.net"
	alertsEndpoint = "v1/alerts"
)

type Client struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

// NewClientFromEnv creates new API client. It takes api key from ASOC_API_KEY env variable.
// Client uses default http client.
func NewClientFromEnv() (*Client, error) {
	apiKey := os.Getenv("ASOC_API_KEY")
	if apiKey == "" {
		return nil, errors.New("required environment variable ASOC_API_KEY not defined")
	}

	return NewClient(apiKey), nil
}

// NewClient creates new API client. Client uses default http client.
func NewClient(apiKey string) *Client {
	return &Client{
		baseURL: APIBase,
		apiKey:  apiKey,
		client:  http.DefaultClient,
	}
}

// SetTimeout allows to set http client timeout.
func (c *Client) SetTimeout(t time.Duration) {
	c.client.Timeout = t
}

func (c *Client) prepareRequest(ctx context.Context, method, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.apiUrl(endpoint), body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.apiKey, "")
	return req, nil
}

func (c *Client) apiUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", c.baseURL, endpoint)
}
