package main

import (
	"data/repo"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
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

	app.writeJSON(w, http.StatusOK, response)
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

	if err != nil {
		app.errJSON(w, err)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: "Ping POST data route",
		Data:    res,
	}

	app.writeJSON(w, http.StatusCreated, response)
}

func (app *Config) PatchData(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if len(id) == 0 {
		app.errJSON(w, errors.New("provide the ID"), http.StatusBadRequest)
		return
	}

	var reqBody JsonRequest

	err := app.readJSON(w, r, &reqBody)

	if err != nil {
		app.errJSON(w, err)
		return
	}

	updatedData := &repo.DataEntry{
		Title:       reqBody.Title,
		Description: reqBody.Description,
	}

	data, err := updatedData.UpdateData(id)

	if err != nil {
		app.errJSON(w, err)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: "Ping PATCH data route",
		Data:    data,
	}

	app.writeJSON(w, http.StatusAccepted, response)
}

func (app *Config) DeleteData(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if len(id) == 0 {
		app.errJSON(w, errors.New("provide the ID"), http.StatusBadRequest)
		return
	}

	data, err := app.Models.DataEntry.Delete(id)

	if err != nil {
		app.errJSON(w, err)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: "Ping DELETE data route",
		Data:    data,
	}

	app.writeJSON(w, http.StatusAccepted, response)
}
