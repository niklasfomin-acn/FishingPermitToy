package api

import (
	"log"
	"net/http"
)

type Server struct {
	listenAddr string
	handlers   map[string]func(http.ResponseWriter, *http.Request)
	//store      storage.Storage
	// TODO: Add more fields for API Gateway, Middleware, etc.
}

// This is the constructor function for the new server
func NewServer(listenAddr string, handlers map[string]func(w http.ResponseWriter, r *http.Request)) *Server {
	return &Server{
		listenAddr: listenAddr,
		handlers:   handlers,
	}
}

func (s *Server) Start() error {
	log.Printf("Backend Server listening on Port 3000\n")
	for route, handler := range s.handlers {
		http.HandleFunc(route, handler)
	}
	return http.ListenAndServe(s.listenAddr, nil)
}
