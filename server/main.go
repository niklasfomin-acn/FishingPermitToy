/*
	GPT Experiment zum Untersuchungsaspekt Module / Datenbanken

Code Snippet zur Anbindung von AWS RDS Postgres an das Backend
Version: 2
Bemerkungen:
*/
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

	/*
		Define Databases here
		[0] MongoDB
		[1] Postgres
		[2] OracleDB
		[3] MongoDB Cluster Deployment
	*/

	// MongoDB local
	//store := storage.NewMongoStorage(config.Databases[0])

	//Postgres
	store, err := storage.NewPostgresStorage(config.Databases[1])
	if err != nil {
		log.Fatal(err)
	}

	// OracleDB
	// store, err := storage.NewOracleStorage(config.Databases[2])
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// MongoDB Cluster Deployment
	//store := storage.NewMongoStorage(config.Databases[3])

	// Define the API handlers here
	h := handlers.New(store)
	handlerFuncs := map[string]func(w http.ResponseWriter, r *http.Request){
		"/SaveCitizenPermit":           h.HandleCitizenPermitRequest,
		"/GetAllCitizenPermitRequests": h.GetCitizenPermitRequests,
	}
	// Start the server
	server := api.NewServer(config.ServerAddress, handlerFuncs, store)
	log.Fatal(server.Start())
}
