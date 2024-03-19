package main

import (
	"flag"
	"fmt"
	"log"

	"rest-backend/api"
	"rest-backend/storage"
)

func main() {
	listenAddr := flag.String("listenaddr", ":8080", "server listen address")
	flag.Parse()
	store := storage.NewMongoStorage()
	server := api.NewServer(*listenAddr, store)
	fmt.Println("FishingPermit Backend API is listening on port:", *listenAddr)
	log.Fatal(server.Start())

}
