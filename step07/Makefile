VERSION ?= `git rev-parse --short HEAD`

.PHONY: build
build:
	go build -ldflags="-X 'main.Version=${VERSION}'"