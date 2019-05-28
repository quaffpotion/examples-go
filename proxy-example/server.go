//In Postman get:
// ... Processing POST

//In chrome loading localhost:8000 get:
// ... Processing GET
// ... Processing GET
// "Hello world!" in browser

//In console running:
//fetch('http://localhost:8000', {mode: 'no-cors'}).then(r => console.log(r))
// ... Processing GET
// (in brwoser: get type "opaque")

//In running cors-test.html:
// Still blocked by CORS even before it gets to server


package main

import (
	"io"
	"net/http"
	"log"

	"github.com/heppu/simple-cors"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing " + r.Method)
	io.WriteString(w, "Hello world!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":8000", cors.CORS(mux))
}
