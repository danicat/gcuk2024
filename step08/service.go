package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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
	type response struct {
		Message string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request received: %s %s", r.Method, r.RequestURI)

		contentType := r.Header.Get("content-type")
		if contentType != "" && contentType != "application/json" {
			err := fmt.Errorf("content type not implemented: %s", contentType)
			http.Error(w, err.Error(), http.StatusNotImplemented)
			log.Print(err)
			return
		}

		switch r.Method {
		case http.MethodGet:
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(response{Message: "pong"})
		default:
			err := fmt.Errorf("method %s not allowed", r.Method)
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		}
	}
}
