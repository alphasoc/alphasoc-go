package alphasoc

import (
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestNew_WithoutOptions(t *testing.T) {
	expectedAPIKey := "testAPIKey"

	os.Setenv(APIKeyEnvVar, expectedAPIKey)
	defer os.Unsetenv(APIKeyEnvVar)

	c, err := New()
	if err != nil {
		t.Fatal("Error creating client, expected none")
	}

	if c.apiKey != expectedAPIKey {
		t.Fatalf("Expected apiKey: %v, got: %v", expectedAPIKey, c.apiKey)
	}
}

func TestNew_WithAPIKeyOption(t *testing.T) {
	expectedAPIKey := "testAPIKey"

	os.Setenv(APIKeyEnvVar, "differentAPIKey")
	defer os.Unsetenv(APIKeyEnvVar)

	c, err := New(WithAPIKey(expectedAPIKey))
	if err != nil {
		t.Fatal("Error creating client, expected none")
	}

	if c.apiKey != expectedAPIKey {
		t.Fatalf("Expected apiKey: %v, got: %v", expectedAPIKey, c.apiKey)
	}
}

func TestNew_WithAPIKeyOption_EmptyEnvVar(t *testing.T) {
	expectedAPIKey := "testAPIKey"

	c, err := New(WithAPIKey(expectedAPIKey))
	if err != nil {
		t.Fatal("Error creating client, expected none")
	}

	if c.apiKey != expectedAPIKey {
		t.Fatalf("Expected apiKey: %v, got: %v", expectedAPIKey, c.apiKey)
	}
}

func TestNew_WithHTTPClientOption(t *testing.T) {
	apiKey := "testAPIKey"
	expectedTimeout := 15 * time.Second

	os.Setenv(APIKeyEnvVar, apiKey)
	defer os.Unsetenv(APIKeyEnvVar)

	c, err := New(WithHTTPClient(&http.Client{Timeout: expectedTimeout}))
	if err != nil {
		t.Fatal("Error creating client, expected none")
	}

	if c.client.Timeout != expectedTimeout {
		t.Fatalf("Expected Timeout: %v, got: %v", expectedTimeout, c.client.Timeout)
	}
}

func TestNew_WithMultipleOptions(t *testing.T) {
	expectedAPIKey := "testAPIKey"
	expectedTimeout := 15 * time.Second

	os.Setenv(APIKeyEnvVar, "differentAPIKey")
	defer os.Unsetenv(APIKeyEnvVar)

	c, err := New(WithAPIKey(expectedAPIKey), WithHTTPClient(&http.Client{Timeout: expectedTimeout}))
	if err != nil {
		t.Fatal("Error creating client, expected none")
	}

	if c.apiKey != expectedAPIKey {
		t.Fatalf("Expected apiKey: %v, got: %v", expectedAPIKey, c.apiKey)
	}

	if c.client.Timeout != expectedTimeout {
		t.Fatalf("Expected Timeout: %v, got: %v", expectedTimeout, c.client.Timeout)
	}
}

func TestNew_ReturnsErrorWhenNoApiKeyProvided(t *testing.T) {
	_, err := New()
	if err == nil {
		t.Fatal("Expected error, got none")
	}

	if !strings.Contains(err.Error(), "apiKey cannot be empty") {
		t.Fatalf("Expected error message should contain 'apiKey cannot be empty', got %v", err.Error())
	}
}
