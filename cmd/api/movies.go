package main

import (
	"strconv"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createMovieHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Create a new movie"))
}

func (app *application) showMovieHandler(rw http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(rw, r)
		return
	}

	fmt.Fprintf(rw, "Show the details of movie %d\n", id)

}