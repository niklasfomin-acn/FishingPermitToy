package main

import (
	"log"
	"net/http"
	"rest-backend/api"
	"rest-backend/handlers"
)

func main() {

	// Define the services
	handlers := map[string]func(w http.ResponseWriter, r *http.Request){
		"/permit":  handlers.HandlePermitRequest,
		"/citizen": handlers.HandleCitizenRequest,
	}
	// Start the server
	server := api.NewServer(":3000", handlers)
	log.Fatal(server.Start())
}
