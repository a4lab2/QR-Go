package main

import (
	"net/http"
	"pqrgen/internal/data"
	"time"
)

func (app *application) showUrlHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	urlData := data.UrlData{
		ID:             id,
		CreatedAt:      time.Now(),
		Url:            "The url to encode",
		Downloadswitch: true,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"urlData": urlData}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
