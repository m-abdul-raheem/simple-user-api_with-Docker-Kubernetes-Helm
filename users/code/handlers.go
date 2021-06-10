package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	//get all users
	users, err := app.users.All()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Convert user list into json encoding
	b, err := json.Marshal(users)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("Users have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	// Find user by id
	m, err := app.users.FindByID(idInt)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("User not found")
			return
		}
		// Any other error will send an internal server error
		app.serverError(w, err)
		return
	}

	// Convert user to json encoding
	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("Have been found a user")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insert(w http.ResponseWriter, r *http.Request) {
	// Define user model
	var u User
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		app.serverError(w, err)
		return
	}

	//validation
	if u.Id <= 0 || u.Name == "" {
		app.serverError(w, errors.New("invalid parameters"))
		return
	}

	// Find user by id
	m, err := app.users.FindByID(u.Id)
	if m != nil {
		app.serverError(w, errors.New("user already exists with id: "+strconv.Itoa(u.Id)))
		return
	}

	// Insert new user
	insertResult, err := app.users.Insert(u)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Printf("New user have been created, id=%s", insertResult.InsertedID)
}
