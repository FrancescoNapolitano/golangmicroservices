package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"franco.it/microserv/handlers"
)

func main() {
	l := log.New(os.Stdout, "Product-api", log.LstdFlags)
	hh := handlers.NewHelloHandler(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}
