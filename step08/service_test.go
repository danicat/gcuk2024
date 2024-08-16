package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService_handlePing(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)

	s := NewService()
	pingHandler := s.handlePing()
	pingHandler(w, req)

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

func TestService_handlePing2(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)

	s := NewService()
	pingHandler := s.handlePing()
	pingHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, but got %d", http.StatusOK, res.StatusCode)
	}

	defer res.Body.Close()

	var msg struct {
		Message string
	}

	err := json.NewDecoder(res.Body).Decode(&msg)
	if err != nil {
		t.Fatalf("error decoding message: %s", err)
	}

	if msg.Message != "pong" {
		t.Fatalf("expected pong, got %s", msg.Message)
	}
}

func TestService_handlePing3(t *testing.T) {

	tbl := []struct {
		name        string
		method      string
		contentType string
		// expected
		status  int
		message string
	}{
		{
			name:    "OK",
			method:  http.MethodGet,
			status:  http.StatusOK,
			message: "pong",
		},
		{
			name:   "POST",
			method: http.MethodPost,
			status: http.StatusMethodNotAllowed,
		},
		{
			name:        "content type xml",
			method:      http.MethodGet,
			contentType: "application/xml",
			status:      http.StatusNotImplemented,
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(tc.method, "/ping", nil)
			if tc.contentType != "" {
				req.Header.Set("Content-Type", tc.contentType)
			}

			s := NewService()
			pingHandler := s.handlePing()
			pingHandler(w, req)

			res := w.Result()
			if res.StatusCode != tc.status {
				t.Fatalf("expected status %d, but got %d", tc.status, res.StatusCode)
			}

			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				return
			}

			var msg struct {
				Message string
			}

			err := json.NewDecoder(res.Body).Decode(&msg)
			if err != nil {
				t.Fatalf("error decoding message: %s", err)
			}

			if msg.Message != tc.message {
				t.Fatalf("expected %s, got %s", tc.message, msg.Message)
			}
		})
	}
}
