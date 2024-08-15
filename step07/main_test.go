package main

import (
	"context"
	"errors"
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
