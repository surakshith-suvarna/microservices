package main

import (
	"log"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Connected to broker",
	}

	/*out, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, err = w.Write(out)
	if err != nil {
		log.Println(err)
	}*/

	err := app.writeJson(w, http.StatusOK, payload)
	if err != nil {
		log.Println(err)
	}

}
