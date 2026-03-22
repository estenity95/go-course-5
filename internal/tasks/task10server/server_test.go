package task10server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerSumEndpoint(t *testing.T) {
	server := httptest.NewServer(NewHandler())
	defer server.Close()

	resp, err := http.Get(server.URL + "/sum?a=2&b=3")
	if err != nil {
		t.Fatalf("GET /sum failed: %v", err)
	}
	defer resp.Body.Close()

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Fatalf("status code = %d, want %d", got, want)
	}

	var got SumResponse
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if got.Result != 5 {
		t.Fatalf("sum result = %d, want %d", got.Result, 5)
	}
}

func TestServerSumEndpointBadRequest(t *testing.T) {
	server := httptest.NewServer(NewHandler())
	defer server.Close()

	resp, err := http.Get(server.URL + "/sum?a=bad&b=3")
	if err != nil {
		t.Fatalf("GET /sum failed: %v", err)
	}

	defer resp.Body.Close()

	if got, want := resp.StatusCode, http.StatusBadRequest; got != want {
		t.Fatalf("status code = %d, want %d", got, want)
	}
}
