package main

import (
	"io"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	// ???
	go main()

	resp, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	if string(body) != `{"message":"pong"}` {
		t.Fatalf("expected pong, got %s", string(body))
	}
}
