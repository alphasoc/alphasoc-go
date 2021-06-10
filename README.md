# AlphaSOC Go



Go SDK for interacting with [AlphaSOC API](https://docs.alphasoc.com/api/api/).



## Usage



### Creating client

There are two ways to create api client:
1. By storing api key in environment variable `ALPHASOC_API_KEY`.
```go
client := alphasoc.New()
```

2. By passing api key as an Option.
```go
client, err := alphasoc.New(alphasoc.WithAPIKey("your-api-key"))
```

Client uses default http client from go net/http package, but if needed it can be changed by passing `alphasoc.WithHTTPClient(*http.Client)` Option while creating API Client.
```go
client, err := alphasoc.New(alphasoc.WithHTTPClient(httpClientWithCustomOptions))
```



### Retrieving alerts

[Alerts](https://docs.alphasoc.com/api/api/#retrieving-alerts) can be retrieved with `client.GetAlerts()` method.

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
	// Creating api client with default options
	client, err := alphasoc.New()
	if err != nil {
		log.Fatal("creating api client", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	alerts, err := client.GetAlerts(ctx, "last-follow-bookmark")
	if err != nil {
		// Any error returned by API is alphasoc.APIError.
		// It contains StatusCode and Message from API.
		if asocErr, ok := err.(alphasoc.APIError); ok {
			if asocErr.StatusCode == http.StatusTooManyRequests {
				// If returned StatusCode is 429, try again later.
			} else {
				// Other error returned by API.
			}
		}

		// Error creating request or sending request
		log.Fatal("retrieving alerts", err)
	}

	// We have retrieved alerts
}
```



## Errors

Any [response](https://docs.alphasoc.com/api/api/#responses) from API other than 200 is wrapped as `alphasoc.APIError`. It contains StatusCode and error Message returned by API.