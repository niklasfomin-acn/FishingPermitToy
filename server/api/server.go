package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"rest-backend/config"
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
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	var config config.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Backend Server Running on Port %v\n", config.ServerAddress)

	for route, handler := range s.handlers {
		http.HandleFunc(route, handler)
	}

	return http.ListenAndServe(s.listenAddr, nil)
}

func ServerDB(choice int, config *config.Config) storage.Storage {
	var store storage.Storage

	switch choice {
	case 1:
		store = storage.NewMongoStorage(config.Databases[0])
	case 2:
		store = storage.NewMongoStorage(config.Databases[0])
	default:
		log.Fatal("Invalid Database Choice")
	}
	return store
}
