package main

import (
	"fmt"
	"log"
	"net/http"
)

const httpPort = "80"

type Config struct{}

func main() {

	app := Config{}

	log.Printf("Starting server at port %s", httpPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
