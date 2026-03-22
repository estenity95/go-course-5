package task05handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()

	PingHandler(rec, req)

	if got, want := rec.Code, http.StatusOK; got != want {
		t.Fatalf("status code = %d, want %d", got, want)
	}

	if got, want := strings.TrimSpace(rec.Body.String()), "pong"; got != want {
		t.Fatalf("response body = %q, want %q", got, want)
	}

	if got := rec.Header().Get("Content-Type"); got != "text/plain; charset=utf-8" {
		t.Fatalf("content type = %q, want %q", got, "text/plain; charset=utf-8")
	}
}
