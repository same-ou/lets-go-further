package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/same-ou/lets-go-further/internal/data"
	"github.com/same-ou/lets-go-further/internal/validator"
)

func (app *application) createMovieHandler(rw http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(rw, r, &input)
	if err != nil {
		app.badRequestResponse(rw, r, err)
		return
	}
	// Copy the values from the input struct to a new Movie struct.
	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}
	// Initialize a new Validator.
	v := validator.New()

	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(rw, r, v.Errors)
		return
	}

	fmt.Fprintf(rw, "%+v", input)
}

func (app *application) showMovieHandler(rw http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(rw, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		Title:     "Casablanca",
		Year:      1942,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(rw, http.StatusOK, envlope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(rw, r, err)
		return
	}

}
