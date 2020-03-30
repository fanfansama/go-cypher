package main

import (
	"github.com/gorilla/mux"
	"controllers"
	"net/http"

)

const HomeFolder = "/app/"
func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /files/ to /files
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/").Name("Home").HandlerFunc(controllers.Home)
	router.Methods("GET").Path("/files").Name("Index").HandlerFunc(controllers.FilesIndex)
	router.Methods("POST").Path("/encode").Name("Encode").HandlerFunc(controllers.Encode)
	router.Methods("POST").Path("/decode").Name("Decode").HandlerFunc(controllers.Decode)

	router.Handle("/files/encoded/{rest}", http.StripPrefix("/files/encoded/", http.FileServer(http.Dir(HomeFolder + "out-enc/"))))
	router.Handle("/files/decoded/{rest}", http.StripPrefix("/files/decoded/", http.FileServer(http.Dir(HomeFolder + "out-dec/"))))

	 
//	router.Methods("GET").Path("/cars/{id}").Name("Show").HandlerFunc(controllers.CarsShow)
//	router.Methods("PUT").Path("/cars/{id}").Name("Update").HandlerFunc(controllers.CarsUpdate)
//	router.Methods("DELETE").Path("/cars/{id}").Name("DELETE").HandlerFunc(controllers.CarsDelete)
	return router
}