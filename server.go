package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("webtest"))
	mux.Handle("/", http.StripPrefix("/", fs))

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	log.Println("Server Started on Port 8080, Address 127.0.0.1")
	log.Fatal(server.ListenAndServe())
}
