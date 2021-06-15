package alphasoc

import "fmt"

// APIError represents an AlphaSOC API Error response
// https://docs.alphasoc.com/api/api/#responses
type APIError struct {
	Message    string
	StatusCode int
}

func (err APIError) Error() string {
	return fmt.Sprintf("alphasoc api error: %s, response code: %d", err.Message, err.StatusCode)
}
