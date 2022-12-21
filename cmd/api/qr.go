package main

import (
	"fmt"
	"net/http"
	"pqrgen/internal/data"
	"time"
)

func (app *application) generateImageHandler(w http.ResponseWriter, r *http.Request) {
	input := struct {
		dataKind     data.Qtype
		createdAt    time.Time
		qr_code_text string
	}{
		dataKind:     1,
		createdAt:    time.Now(),
		qr_code_text: "Hello There",
	}

	// err := app.readJSON(w, r, &input)
	// if err != nil {
	// 	app.badRequestResponse(w, r, err)
	// 	return
	// }

	qr := &data.QRData{
		DataKind:     input.dataKind,
		CreatedAt:    input.createdAt,
		Qr_code_text: input.qr_code_text,
	}
	println(qr)
	// err = json.NewDecoder(r.Body).Decode(&input)
	// if err != nil {
	// 	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	// 	return

	// }
	fmt.Fprintf(w, "%+v\n", input)
}
