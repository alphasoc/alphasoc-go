package alphasoc

import "fmt"

type APIError struct {
	Message    string
	StatusCode int
}

func (err APIError) Error() string {
	return fmt.Sprintf("alphasoc api error: %s, response code: %d", err.Message, err.StatusCode)
}
