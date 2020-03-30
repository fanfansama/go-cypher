package controllers

import (
	"encoding/json"
	"models"
	"net/http"
)

func FilesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AllFiles())
}

func FilesShow(w http.ResponseWriter, r *http.Request, filename string) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.NewFile(filename))
}