package alphasoc

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	APIBase        = "https://api.alphasoc.net"
	APIKeyEnvVar   = "ALPHASOC_API_KEY"
	alertsEndpoint = "v1/alerts"
)

type Client struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

type Option func(*Client)

// New creates new API client. By default it loads APIKey from ALPHASOC_API_KEY
// env variable and uses Default http Client.
// To change default values use WithAPIKey and WithHTTPClient Options.
func New(opts ...Option) (*Client, error) {
	apiKey := os.Getenv(APIKeyEnvVar)

	c := &Client{
		baseURL: APIBase,
		apiKey:  apiKey,
		client:  http.DefaultClient,
	}

	for _, option := range opts {
		option(c)
	}

	if c.apiKey == "" {
		return nil, fmt.Errorf("apiKey cannot be empty, set it using %s env variable or WithAPIKey Option", APIKeyEnvVar)
	}

	return c, nil
}

// WithAPIKey allows creating Client with provided apiKey
// instead of taking it from environment variable.
func WithAPIKey(apiKey string) Option {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// WithHTTPClient allows creating Client with provided http client
// instead of go default client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
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
