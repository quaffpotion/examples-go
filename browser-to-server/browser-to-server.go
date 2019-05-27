package main

import (
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("hello friend")
}

func main() {

	http.HandleFunc("/", myHandler)

	panic(http.ListenAndServe(":8000", nil))
}
