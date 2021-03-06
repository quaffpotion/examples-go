package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"log"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	// Use default options
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8000", handler))

}
