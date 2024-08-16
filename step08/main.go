package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
)

var Version = "develop"

func main() {
	addr := flag.String("addr", ":8080", "address to listen to")
	flag.Parse()

	ctx := context.Background()

	err := run(ctx, *addr)
	if err != nil {
		log.Fatalf("failed to run app: %s", err)
	}
}

func run(ctx context.Context, addr string) error {
	log.Printf("GOMAXPROCS: %d", runtime.GOMAXPROCS(0))
	log.Printf("go runtime version: %s", runtime.Version())
	log.Printf("server version: %s", Version)

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	go func() {
		s := NewService()

		log.Printf("server listening on: %s", addr)
		err := http.ListenAndServe(addr, s)
		if err != nil {
			log.Fatalf("failed to listen: %s", err)
		}
	}()

	<-ctx.Done()
	return ctx.Err()
}
