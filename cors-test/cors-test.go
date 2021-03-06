package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/cors"

	"log"
)

type test_struct struct {
	Test string
}

func main() {

	finish := make(chan bool)

	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling 8001")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world 1\"}"))
		w.Header().Set("Access-Control-Allow-Origin", "*")

		myDecoder := json.NewDecoder(r.Body)

		var t test_struct
		err := myDecoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		log.Println(t.Test)

		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	panic(err.Error())
		// }
		// var myData interface{}
		// err = json.Unmarshal(body, &myData)
		// if err != nil {
		// 	panic(err.Error())
		// }
		// log.Println("Results: %v\n", myData)

	})
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling 8002")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world 2\"}"))
		w.Header().Set("Access-Control-Allow-Origin", "*")

		myDecoder := json.NewDecoder(r.Body)

		var t test_struct
		err := myDecoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		log.Println(t.Test)

		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	panic(err.Error())
		// }
		// var myData interface{}
		// err = json.Unmarshal(body, &myData)
		// if err != nil {
		// 	panic(err.Error())
		// }
		// log.Println("Results: %v\n", myData)

	})

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	handler1 := cors.Default().Handler(mux1)

	//c below takes the place of cors.Default() but doesn't allow any origins
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // try "*" to make everything go through
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

	// r, err := http.Get("https://localhost:8001")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// var myData interface{}
	// err = json.Unmarshal(body, &myData)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// log.Println("Results from hardcoded GET: %v\n", myData)

	<-finish

}
