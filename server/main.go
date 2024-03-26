package main

import (
	"log"
	"rest-backend/api"
	"rest-backend/handlers"
)

func main() {
	server := api.NewServer(":8080", handlers.HandlePermitRequest)
	log.Fatal(server.Start())
}
