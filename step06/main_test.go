package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

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
}
