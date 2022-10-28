package main

import (
	"errors"
	"net/http"
	"strconv"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {

	rawID := r.URL.Query().Get(":id")
	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		return 0, errors.New("invalid Id Parameter")
	}
	return id, nil
}
