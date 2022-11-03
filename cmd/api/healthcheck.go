package main

import (
	"net/http"
)

const version = "1.0.0"

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := envelope{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)

	}

}
