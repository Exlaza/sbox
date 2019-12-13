package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))

	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)

}
