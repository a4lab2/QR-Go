package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {

	m := pat.New()
	m.Post("/v1/qr", http.HandlerFunc(genQrHandler))
	m.Get("/v1/qr/:id", http.HandlerFunc(showQrHandler))

	return m
}
