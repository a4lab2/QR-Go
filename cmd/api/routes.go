package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {

	m := pat.New()

	m.NotFound = http.HandlerFunc(app.notFoundResponse)
	// m.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	m.Get("/v1/healthcheck", http.HandlerFunc(app.healthcheckHandler))

	m.Get("/v1/qr/:id", http.HandlerFunc(app.showUrlHandler))

	return m
}
