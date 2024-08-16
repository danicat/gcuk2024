package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService_handlePing(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	PingHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, but got %d", http.StatusOK, res.StatusCode)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	if string(body) != `{"message":"pong"}` {
		t.Fatalf("expected pong, got %s", string(body))
	}
}
