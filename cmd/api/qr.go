package main

import (
	"fmt"
	"net/http"
)

func genQrHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Generate qrcode")
}

func showQrHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get(":id")
	fmt.Fprintf(w, "Show  qrcode %s", id)
}
