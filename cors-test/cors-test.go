package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler1 := cors.Default().Handler(mux1)

c := cors.New(cors.Options{
    AllowedOrigins: []string{""},
    AllowCredentials: true,
    // Enable Debugging for testing, consider disabling in production
    Debug: true,
})

	handler2 := c.Handler(mux2) 

	
	//probably not the correct way to run things? e.g. maybe goroutines

	//handler1 has default CORS options so browser will show response
	http.ListenAndServe(":8081", handler1)

	//handler2 has no CORS so browser will prevent it
	http.ListenAndServe(":8082", handler2)

}
