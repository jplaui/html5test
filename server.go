package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("src/webtest"))
	mux.Handle("/", http.StripPrefix("/", fs))

	server := http.Server{
		Addr:    "172.17.0.2:8080",
		Handler: mux,
	}
	log.Println("Server Started at 172.17.0.2:8080")
	log.Fatal(server.ListenAndServe())
}

