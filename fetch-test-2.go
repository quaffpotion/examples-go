// curl -H "Origin: http://example.com" -H "Access-Control-Request-Method: POST" -H "Access-Control-Request-Headers: X-Requested-With" -X OPTIONS --verbose localhost:8080


package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

func handleEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"hello\": \"world\"}"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", handleEndpoint)

    log.Fatal(http.ListenAndServe(":8080", 
        handlers.LoggingHandler(os.Stdout, handlers.CORS(
            handlers.AllowedMethods([]string{"POST"}),
            handlers.AllowedOrigins([]string{"*"}),
            handlers.AllowedHeaders([]string{"X-Requested-With"}))(router))))
}
