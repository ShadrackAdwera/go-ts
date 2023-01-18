package main

import "net/http"

type JsonRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (app *Config) GetData(w http.ResponseWriter, r *http.Request) {
	response := jsonResponse{
		Error:   false,
		Message: "Ping GET data route",
	}

	app.writeJSON(w, http.StatusCreated, response)
}

func (app *Config) PostData(w http.ResponseWriter, r *http.Request) {
	var jsonRequest JsonRequest

	err := app.readJSON(w, r, &jsonRequest)

	if err != nil {
		app.errJSON(w, err)
		return
	}
	response := jsonResponse{
		Error:   false,
		Message: "Ping POST data route",
		Data:    jsonRequest,
	}

	app.writeJSON(w, http.StatusCreated, response)
}

func (app *Config) PatchData(w http.ResponseWriter, r *http.Request) {
	response := jsonResponse{
		Error:   false,
		Message: "Ping PATCH data route",
	}

	app.writeJSON(w, http.StatusCreated, response)
}

func (app *Config) DeleteData(w http.ResponseWriter, r *http.Request) {
	response := jsonResponse{
		Error:   false,
		Message: "Ping DELETE data route",
	}

	app.writeJSON(w, http.StatusCreated, response)
}
