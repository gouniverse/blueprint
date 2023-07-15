package server

import (
	"net/http"
	"time"
)

type Server struct {
	*http.Server
}

func (s *Server) Start() error {
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}

func NewServer(address string, handler func(w http.ResponseWriter, r *http.Request)) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	return &Server{
		&http.Server{
			Addr:    address,
			Handler: mux,
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout:      15 * time.Second,
			ReadTimeout:       15 * time.Second,
			IdleTimeout:       30 * time.Second,
			ReadHeaderTimeout: 2 * time.Second,
		},
	}
}
