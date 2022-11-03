package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// func genCode() {
// 	str := "mailto:a@mail.com&subject=Mail subject&body=Random text body"
// 	qrFname := "qr.png"

// 	err := qrcode.WriteFile(str, qrcode.Medium, 256, qrFname)
// 	if err != nil {
// 		panic(err)
// 	}
// }

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "", 4000, "Api server port")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//initialize app
	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
	// genCode()
}
