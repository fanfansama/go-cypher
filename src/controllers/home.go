package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello")
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	u := []string{"Hello"}
	json.NewEncoder(w).Encode(u)
}