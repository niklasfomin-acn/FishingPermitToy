package main

import (
	"encoding/json"
	"log"
	"os"
	"rest-backend/api"
	"rest-backend/auth"
	"rest-backend/config"
	"rest-backend/handlers"
	"rest-backend/storage"
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

	// fmt.Println("Init Database Server:")
	// var choice int
	// _, err = fmt.Scan(&choice)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Get the connection string for the storage
	pathFromVault, _, err := auth.GetSecretFromVault(true, false, false, &config)
	if err != nil {
		log.Fatalf("Error getting secret from vault: %v", err)
	}
	//log.Printf("Secret from vault: %v", pathFromVault)
	store := storage.NewMongoStorage(pathFromVault)

	// Parse in the API handlers
	h := handlers.New(store)
	handlerFuncs := handlers.GetHandlerFuncs(h)

	// Start the server
	server := api.NewServer(config.ServerAddress, handlerFuncs, store)
	log.Fatal(server.Start())
}
