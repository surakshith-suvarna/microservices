package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	log.Println("User: ", requestPayload.Email)
	log.Println("Password: ", requestPayload.Password)
	//validate the user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		log.Println(err)
		//app.errorJson(w, errors.New("credentials invalid"), http.StatusBadRequest)
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		//app.errorJson(w, errors.New("credentials invalid"), http.StatusBadRequest)
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	app.writeJson(w, http.StatusAccepted, payload)

}
