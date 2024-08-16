package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	t.Cleanup(cancel)

	err := run(ctx, ":8080")
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected no errors, but got %s", err)
	}
}

func TestPing(t *testing.T) {
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
