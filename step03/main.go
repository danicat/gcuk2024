package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	err := run(ctx)
	if err != nil {
		log.Fatalf("failed to run app: %s", err)
	}
}

func run(ctx context.Context) error {
	http.HandleFunc("/ping", PingHandler)

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatalf("failed to listen: %s", err)
		}
	}()

	<-ctx.Done()
	return ctx.Err()
}

func PingHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`{"message":"pong"}`))
}
