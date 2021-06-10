package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/", app.all).Methods("GET")
	r.HandleFunc("/users/{id}", app.findByID).Methods("GET")
	r.HandleFunc("/users/", app.insert).Methods("POST")
	return r
}
