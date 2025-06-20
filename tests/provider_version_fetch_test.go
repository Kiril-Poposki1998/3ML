package tests

import (
	handleform "3ML/handleForm"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProviderVersionFetch(t *testing.T) {
	expectedPlaceholder := "eg. 3.5.0"

	// Mock the HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/provider-version-placeholder" {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		response := map[string]string{"placeholder": expectedPlaceholder}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Set the base URL for the mock server
	baseURL := server.URL

	// Call the function to test
	placeholder, err := handleform.FetchVersionPlaceholder(baseURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if placeholder != expectedPlaceholder {
		t.Errorf("Expected placeholder %s, got %s", expectedPlaceholder, placeholder)
	}
}
