package main

func (s *Service) setupRoutes() {
	s.mux.Handle("/ping", s.handlePing())
}
