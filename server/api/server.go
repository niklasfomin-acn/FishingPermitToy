package api

import (
	"net/http"
)

type Server struct {
	listenAddr string
	handler    func(http.ResponseWriter, *http.Request)
	//store      storage.Storage
	// TODO: Add more fields for API Gateway, Middleware, etc.
}

// This is the constructor function for the new server
func NewServer(listenAddr string, handler func(w http.ResponseWriter, r *http.Request)) *Server {
	return &Server{
		listenAddr: listenAddr,
		handler:    handler,
		// A server is reponsible for handling requests from different clients
		// It needs to know how to handle the requests and route them to the correct handler
		// It needs to provide database interfaces to the handlers
		// It needs to provide documentAI interfaces to the handlers
		// It needs to provide simple business logic to process the users' requests and the admins processingG
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/permit", s.handler)
	return http.ListenAndServe(s.listenAddr, nil)
}
