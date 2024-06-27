package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
    // Use the http.NewServeMux() function to create an empty servemux.
	mux := http.NewServeMux()

	log.Print("Listening...")

	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("./public"))))
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type:", "text/plain; charset=utf-8")
		io.WriteString(w, "OK")
	})

	// Then we create a new server and start listening for incoming requests
	// with the http.ListenAndServe() function, passing in our servemux for it to
	// match requests against as the second argument.
	http.ListenAndServe(":8080", mux)
}