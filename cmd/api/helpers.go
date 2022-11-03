package main

import (
	"encoding/json"
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

type envelope map[string]interface{}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')

	for k, v := range headers {
		w.Header()[k] = v
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

// func genCode() {

// 	file, err := os.Open("vcard.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	b, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	str := string(b)
// 	qrFname := "vcard-qr.png"

// 	err = qrcode.WriteFile(str, qrcode.Medium, 256, qrFname)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
