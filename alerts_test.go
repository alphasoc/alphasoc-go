package alphasoc

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetAlerts_CheckRequestPath(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	expectedPath := "/v1/alerts"

	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expectedPath {
			t.Fatalf("Expected: %v in path, got: %v", expectedPath, r.URL.Path)
		}

		rw.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c, err := New(WithAPIKey("apikey"))
	if err != nil {
		t.Fatal("Error creating api client", err)
	}
	c.baseURL = ts.URL

	_, err = c.GetAlerts(ctx, "")
	if err != nil {
		t.Log(err)
	}
}

func TestGetAlerts_CheckIfAPIKeyIsAddedToRequest(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	expectedAPIKey := "apikey123"

	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user, pwd, ok := r.BasicAuth()
		if !ok {
			t.Fatalf("No Basic Authorization in request")
		}

		if user != expectedAPIKey || pwd != "" {
			t.Fatalf("Expected API key: %v, got: %v", expectedAPIKey, user)
		}

		rw.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c, err := New(WithAPIKey("apikey123"))
	if err != nil {
		t.Fatal("Error creating api client", err)
	}
	c.baseURL = ts.URL

	_, err = c.GetAlerts(ctx, "")
	if err != nil {
		t.Log(err)
	}
}

func TestGetAlerts_CheckIfParametersAreAddedToRequest(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	expectedFollow := "follow123"

	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		t.Log(r.URL.RawQuery)
		query := r.URL.Query()

		if query.Get("follow") != expectedFollow {
			t.Fatalf("Expected follow parameter: %v, got %v", expectedFollow, query.Get("follow"))
		}

		rw.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c, err := New(WithAPIKey("apikey"))
	if err != nil {
		t.Fatal("Error creating api client", err)
	}
	c.baseURL = ts.URL

	_, err = c.GetAlerts(ctx, "follow123")
	apiErr, ok := err.(APIError)
	if err != nil && ok {
		if apiErr.StatusCode == http.StatusOK {
			return
		}
	}

	t.Fatal(err)
}

func TestGetAlerts_AlertsResponseParsedProperly(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		resp := `{
			"follow": "id",
			"more": true,

			"alerts": [
			  {
				"eventType": "dns",
				"event": {
				  "ts": "2018-03-01T10:31:59Z",
				  "srcIP": "192.168.20.5",
				  "srcHost": "john-pc",
				  "query": "www.example.com",
				  "qtype": "A"
				},
				"threats": ["c2_communication"],
				"wisdom": {
				  "flags": ["c2", "young_domain"],
				  "labels": ["c2:TrickBot"]
				}
			  }
			],

			"threats": {
			  "c2_communication": {
				"title": "C2 communication attempt indicating infection",
				"policy": false,
				"severity": 5
			  }
			}
		  }`

		rw.Write([]byte(resp))
	}))
	defer ts.Close()

	c, err := New(WithAPIKey("apikey"))
	if err != nil {
		t.Fatal("Error creating api client", err)
	}
	c.baseURL = ts.URL

	alerts, err := c.GetAlerts(ctx, "")
	if err != nil {
		t.Fatalf("Error during request execution: %+v", err)
	}

	if *alerts.Follow != "id" {
		t.Fatalf("Expected follow: id, got %v", alerts.Follow)
	}

	if *alerts.More != true {
		t.Fatalf("Expected more: true, got %v", alerts.More)
	}
}

func TestGetAlerts_CheckResponseErrors(t *testing.T) {
	type testCase struct {
		respCode         int
		respBody         string
		expectedErrorMsg string
		exactErrorMsg    bool
	}

	cases := []testCase{
		{http.StatusBadRequest, `{"message": "bad request in message"}`, "bad request in message", true},
		{http.StatusBadRequest, `{"notMessage": "bad request in message"}`, "status bad request", true},
		{http.StatusBadRequest, `{`, "status bad request, json decode error", false},
		{http.StatusUnauthorized, `{"message": "unauthorized in message"}`, "unauthorized in message", true},
		{http.StatusUnauthorized, `{"notMessage": "unauthorized in message"}`, "status unauthorized", true},
		{http.StatusUnauthorized, `{`, "status unauthorized, json decode error", false},
		{http.StatusForbidden, `{"message": "forbidden in message"}`, "forbidden in message", true},
		{http.StatusForbidden, `{"notMessage": "forbidden in message"}`, "status forbidden", true},
		{http.StatusForbidden, `{`, "status forbidden, json decode error", false},
		{http.StatusTooManyRequests, `{"message": "too many requests in message"}`, "too many requests in message", true},
		{http.StatusTooManyRequests, `{"notMessage": "too many requests in message"}`, "status too many requests", true},
		{http.StatusTooManyRequests, `{`, "status too many requests, json decode error", false},
		{http.StatusInternalServerError, ``, "unexpected response", true},
	}

	for _, test := range cases {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Add("Content-Type", "application/json")

			rw.WriteHeader(test.respCode)
			rw.Write([]byte(test.respBody))
		}))

		c, err := New(WithAPIKey("apikey"))
		if err != nil {
			t.Fatal("Error creating api client", err)
		}
		c.baseURL = ts.URL

		_, err = c.GetAlerts(ctx, "")
		if err == nil {
			t.Fatal("Expected errror, got nil, testCase:", test)
		}

		if apiErr, ok := err.(APIError); ok {
			if apiErr.StatusCode != test.respCode {
				t.Fatalf("Expected %v statusCode, got %v, testCase: %+v", test.respCode, apiErr.StatusCode, test)
			}

			if test.exactErrorMsg {
				if apiErr.Message != test.expectedErrorMsg {
					t.Fatalf("Expected message: %v, got: %v, testCase: %+v", test.expectedErrorMsg, apiErr.Message, test)
				}
			} else {
				if !strings.Contains(apiErr.Message, test.expectedErrorMsg) {
					t.Fatalf("Expected message to contain: %v, got: %v, testCase: %+v", test.expectedErrorMsg, apiErr.Message, test)
				}
			}
		} else {
			t.Fatalf("Expected APIError, got %v, testCase: %+v", err, test)
		}

		ts.Close()
		cancel()
	}
}
