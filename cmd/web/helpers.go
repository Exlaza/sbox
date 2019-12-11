package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)



func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

// Just a wrapper around client error specifically for the ont found error

func (app *application) notFound(w http.ResponseWriter){
	app.clientError(w, http.StatusNotFound)
}