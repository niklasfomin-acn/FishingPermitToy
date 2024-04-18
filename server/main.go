package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"rest-backend/api"
	"rest-backend/config"
	"rest-backend/handlers"
)

func main() {

	// Parse in the Database Server configuration
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	var config config.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Init Database Server:")
	var choice int
	_, err = fmt.Scan(&choice)
	if err != nil {
		log.Fatal(err)
	}

	store := api.ServerDB(choice, &config)

	// Parse in the API handlers
	h := handlers.New(store)
	handlerFuncs := handlers.GetHandlerFuncs(h)

	// Start the server
	server := api.NewServer(config.ServerAddress, handlerFuncs, store)
	log.Fatal(server.Start())
}
