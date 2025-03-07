package main

import (
	"net/http"
	"fmt"
)


func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

func (app *application) errorResponse(rw http.ResponseWriter, r *http.Request, status int, message string) {
	env := envlope{"error": message}
	err := app.writeJSON(rw, status, env, nil)
	if err != nil {
		app.logError(r, err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(rw http.ResponseWriter, r *http.Request, err error) {
	message := "The server encountered a problem and could not process your request"
	app.errorResponse(rw, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(rw http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	app.errorResponse(rw, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(rw http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported by this resource", r.Method)
	app.errorResponse(rw, r, http.StatusMethodNotAllowed, message)
}

func (app *application) badRequestResponse(rw http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(rw, r, http.StatusBadRequest, err.Error())
}


func (app *application) failedValidationResponse(rw http.ResponseWriter, r *http.Request, errors map[string]string) {
	env := envlope{"errors": errors} // Wrap the errors in an envelope structure
	err := app.writeJSON(rw, http.StatusUnprocessableEntity, env, nil)
	if err != nil {
		app.serverErrorResponse(rw, r, err)
	}
}