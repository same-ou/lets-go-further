package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(rw http.ResponseWriter, r *http.Request) {
	data := envlope{
		"status": "available",
		"system_info": map[string]string{
		"environment": app.config.env,
		"version": version,
		},
	}
	err := app.writeJSON(rw, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(rw, r, err)
		return
	}

}	