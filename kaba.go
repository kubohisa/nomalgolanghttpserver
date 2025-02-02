package main

import (
	"fmt"
	"log"
	"net/http"

	"context"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	fs := http.FileServer(http.Dir("./web/"))
	http.Handle("/", fs)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: addHeader(fs),
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	fmt.Print("Kaba Server.\r\n")
	fmt.Print("kubohisa. Poppyright 2025\r\n")
	fmt.Print("----\r\n")
	fmt.Print("Listening.\r\n\r\n")

	log.Fatal(server.ListenAndServe())
}

func addHeader(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate; post-check=0, pre-check=0")
		w.Header().Set("pragma", "no-cache")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		fs.ServeHTTP(w, r)
	}
}
