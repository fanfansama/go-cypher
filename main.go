package main

import (
	"log"
	"net/http"
)


func main() {
	router := InitializeRouter()

	log.Fatal(http.ListenAndServe(":4450", router))
}