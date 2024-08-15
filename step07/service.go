package main

import "net/http"

type Service struct {
	mux *http.ServeMux
}

func NewService() *Service {
	var s Service

	s.mux = http.NewServeMux()
	s.setupRoutes()

	return &s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.mux.ServeHTTP(w, req)
}

func (s *Service) handlePing() http.HandlerFunc {
	return PingHandler
}

func PingHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`{"message":"pong"}`))
}
