package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "5001"

type Config struct{}

func main() {
	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	log.Printf("Starting DATA service on port %s\n", webPort)

	err := srv.ListenAndServe()

	if err != nil {
		log.Panicln(err)
		return
	}
}
