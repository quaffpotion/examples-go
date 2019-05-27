// CORS provides Cross-Origin Resource Sharing middleware.
// Example:

package main

import (
    "net/http"

    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    
    "log"
)

func myHandler (w http.ResponseWriter, r *http.Request) {
	log.Println("> Handling request...")
	if (*r).Method == "OPTIONS" {
		log.Println(">> Handling options...")
		return
	}

}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", myHandler)

    // Apply the CORS middleware to our top-level router, with the defaults.
    http.ListenAndServe(":8000", handlers.CORS()(r))
}

// 5-27-19: 
// 	localhost:8000 in browser works
