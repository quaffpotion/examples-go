package main

import (
	"net/http"
	"log"
)


func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("hello friend")
	enableCors(&w)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}	

func main() {
	http.Handle("/", http.FileServer(http.Dir("htmlfiles/")))
	http.HandleFunc("/hello", myHandler)
	log.Printf("Navigate to localhost:80 to see files and localhost:80/hello to see a message here, open or refresh test.html to use javascript to send a fetch request to the server; you'll see another message")
	log.Printf("Starting server on port 80")
	panic(http.ListenAndServe("127.0.0.1:80", nil))
}
