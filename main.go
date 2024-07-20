package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

type apiConfig struct {
	fileserverHits int
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Intercepting ", cfg.fileserverHits)
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	}) 
}

func (cfg *apiConfig) metrics() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type:", "text/plain; charset=utf-8")
		response := "Hits: " + strconv.Itoa(cfg.fileserverHits)
		io.WriteString(w, response)
	}) 
}

func (cfg *apiConfig) reset() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits = 0
	}) 
}

func main() {
    // Use the http.NewServeMux() function to create an empty servemux.
	mux := http.NewServeMux()

	log.Print("Listening...")

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type:", "text/plain; charset=utf-8")
		io.WriteString(w, "OK")
	})
	
	fileHandle := http.StripPrefix("/app/", http.FileServer(http.Dir("./public")))
	apiCfg := apiConfig{0};
	mux.Handle("/app/*", apiCfg.middlewareMetricsInc(fileHandle))
	mux.Handle("POST /metrics", apiCfg.metrics())
	mux.Handle("DELETE /reset", apiCfg.reset())


	http.ListenAndServe(":8080", mux)
}
