package main

import (
	"log"
	"net/http"

	"github.com/surakshith-suvarna/microservices/logger-service/data"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPayload
	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJson(w, err)
	}

	//insert the data
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}
	err = app.Models.LogEntry.Insert(event)
	if err != nil {
		log.Println(err)
		app.errorJson(w, err)
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	err = app.writeJson(w, http.StatusAccepted, resp)
	if err != nil {
		log.Println(err)
		app.errorJson(w, err)
	}
}
