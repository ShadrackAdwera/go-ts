package main

import (
	"data/repo"
	"net/http"
)

type JsonRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (app *Config) GetData(w http.ResponseWriter, r *http.Request) {

	data, err := app.Models.DataEntry.GetData()

	if err != nil {
		app.errJSON(w, err)
	}

	response := jsonResponse{
		Error:   false,
		Message: "Ping GET data route",
		Data:    data,
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

	reqBody := repo.DataEntry{
		Title:       jsonRequest.Title,
		Description: jsonRequest.Description,
	}

	res, err := app.Models.DataEntry.Insert(reqBody)

	response := jsonResponse{
		Error:   false,
		Message: "Ping POST data route",
		Data:    res,
	}

	app.writeJSON(w, http.StatusCreated, response)
}

func (app *Config) PatchData(w http.ResponseWriter, r *http.Request) {
	data, err := app.Models.DataEntry.UpdateData()

	if err != nil {
		app.errJSON(w, err)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: "Ping PATCH data route",
		Data:    data,
	}

	app.writeJSON(w, http.StatusCreated, response)
}

func (app *Config) DeleteData(w http.ResponseWriter, r *http.Request) {

	data, err := app.Models.DataEntry.Delete()

	if err != nil {
		app.errJSON(w, err)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: "Ping DELETE data route",
		Data:    data,
	}

	app.writeJSON(w, http.StatusCreated, response)
}
