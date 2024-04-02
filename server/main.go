package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"rest-backend/api"
	"rest-backend/config"
	"rest-backend/handlers"
	"rest-backend/storage"
)

func main() {

	// Read the configuration for the server
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	var config config.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Define the handlers and database for the server
	store := storage.NewMongoStorage(config.Databases[0])
	h := handlers.New(store)
	handlerFuncs := map[string]func(w http.ResponseWriter, r *http.Request){
		"/permit":  h.HandlePermitRequest,
		"/citizen": h.HandleCitizenRequest,
	}
	// Start the server
	server := api.NewServer(config.ServerAddress, handlerFuncs, store)
	log.Fatal(server.Start())
}
