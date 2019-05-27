package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {

	finish := make(chan bool)

	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world 1\"}"))
	})
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world 2\"}"))
	})

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	handler1 := cors.Default().Handler(mux1)

	//c below takes the place of cors.Default() but doesn't allow any origins
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{""},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler2 := c.Handler(mux2)

	go func() {
		http.ListenAndServe(":8001", handler1)
	}()

	go func() {
		http.ListenAndServe(":8002", handler2)
	}()

	<-finish

}
