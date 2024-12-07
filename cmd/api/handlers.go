package main

import (
	"net/http"
)


func (app *application) home(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello, World!"))
}