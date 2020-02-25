package main

import (
	. "go-rest-mongodb/config"
	"go-rest-mongodb/routers"
	"log"
	"net/http"
	"time"
)

var config = Config{}

func init() {
	config.Read()
}

func main() {
	routers := routers.Routers()
	server := &http.Server {
		Handler: routers,
		Addr:    config.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
