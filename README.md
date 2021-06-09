# AlphaSOC Go



Go SDK for interacting with [AlphaSOC API](https://docs.alphasoc.com/api/api/).



## Usage



### Creating client

There are two ways to create api client:
1. By passing api key as an argument to function.
```go
client := alphasoc.NewClient("your-api-key")
```

2. By storing api key in environment variable `ASOC_API_KEY`.
```go
client, err := alphasoc.NewClientFromEnv()
```

Client uses default http client from go net/http package. Requests can be invoked with context or without it. It is possible to set http client timeout with `client.SetTimeout(time.Duration)` method.



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
	client := alphasoc.NewClient("your-api-key")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	alerts, err := client.GetAlertsCtx(ctx, "last-follow-bookmark")
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