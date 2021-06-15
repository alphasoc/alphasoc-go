[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/alphasoc/alphasoc-go)

# AlphaSOC Go

Go SDK for interacting with [AlphaSOC API](https://docs.alphasoc.com/api/api/).

## Usage

```go
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/alphasoc/alphasoc-go"
)

func main() {
	// Create the client with default options. If you want to pass your
	// API key directly instead of getting it from the ALPHASOC_API_KEY
	// environment variable, use the WithAPIKey option:
	//
	// client, err := alphasoc.New(alphasoc.WithAPIKey("your-api-key"))
	//
	client, err := alphasoc.New()
	if err != nil {
		log.Fatal("creating api client", err)
	}

	// In every response there is a `Follow` bookmark attached, which should be passed
	// to consecutive requests as a parameter, so only new alerts are being returned.
	// Once the last page is returned, `More` property in the response is set to false.
	var follow string

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

		alerts, err := client.Alerts(ctx, follow)
		cancel()

		if err != nil {
			// Errors returned by the API are of alphasoc.APIError type.
			if apiErr, ok := err.(alphasoc.APIError); ok {
				if apiErr.StatusCode == http.StatusTooManyRequests {
					// If StatusCode is 429, try again later.
					time.Sleep(1 * time.Minute)
					continue
				} else {
					// Other error returned by API.
					log.Fatal("api error", err)
				}
			}

			// An error not directly related to the API has occurred
			// (e.g. a network error).
			log.Fatal("retrieving alerts", err)
		}

		// We have retrieved the alerts
		fmt.Println(alerts)

		// No more alerts at the moment
		if !*alerts.More {
			return
		}

		// Set next follow bookmark
		follow = *alerts.Follow
	}
}

```
