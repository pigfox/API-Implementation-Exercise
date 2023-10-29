package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var db DB
var key Key
var apiKeyName = "X-API-KEY"

func main() {
	setUp()

	port := ":8888"
	router := mux.NewRouter()
	router.HandleFunc("/alerts", createHandler).Methods("POST")
	router.HandleFunc("/alerts", serviceHandler).Methods("GET")
	router.HandleFunc("/list", listHandler).Methods("GET")

	fmt.Println("Listening on port" + port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
