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
		Handler: addHeader(fs),
	}

	fmt.Print("Kaba Server.\r\n")
	fmt.Print("kubohisa. Poppyright 2025\r\n")
	fmt.Print("----\r\n")
	fmt.Print("Listening.\r\n")
	
	log.Fatal(server.ListenAndServe())
}

func addHeader(fs http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Cache-Control", "no-cache")
        fs.ServeHTTP(w, r)
    }
}