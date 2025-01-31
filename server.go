package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: nil,
	}

	fmt.Print("Listening.")
	
	log.Fatal(server.ListenAndServe())
}