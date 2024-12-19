package main

import (
	"net/http"
	"time"

	"github.com/same-ou/lets-go-further/internal/data"
)

func (app *application) createMovieHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Create a new movie"))
}

func (app *application) showMovieHandler(rw http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(rw, r)
		return
	}

	movie := data.Movie{
		ID: id,
		Title: "Casablanca",
		Year: 1942,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Runtime: 102,
		Genres: []string{"drama", "romance", "war"},
		Version: 1,
	}

	err = app.writeJSON(rw, http.StatusOK, envlope{"movie":movie}, nil)
	if err != nil {
		app.serverErrorResponse(rw, r, err)
		return
	}

}