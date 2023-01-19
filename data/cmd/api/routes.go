package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{}))

	mux.Use(middleware.Heartbeat("/ping"))
	// TODO: use auth middleware here . . .
	mux.Get("/data", app.GetData)
	mux.Post("/data", app.PostData)
	mux.Patch("/data", app.PatchData)   // mux.Patch("/data/:id", app.PatchData)
	mux.Delete("/data", app.DeleteData) // mux.Delete("/data/:id", app.DeleteData)

	return mux
}
