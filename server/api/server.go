package api

import (
	"log"
	"net/http"
	"rest-backend/storage"
)

type Server struct {
	listenAddr string
	handlers   map[string]func(http.ResponseWriter, *http.Request)
	store      storage.Storage
}

// This is the constructor function for the new server
func NewServer(listenAddr string, handlers map[string]func(w http.ResponseWriter, r *http.Request), store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		handlers:   handlers,
		store:      store,
	}
}

func (s *Server) Start() error {
	log.Printf("Backend Server listening on Port 3000\n")
	for route, handler := range s.handlers {
		http.HandleFunc(route, handler)
	}
	return http.ListenAndServe(s.listenAddr, nil)
}
