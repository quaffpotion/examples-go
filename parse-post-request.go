// test with:
// curl -X POST -d '{"test": "that"}' http://localhost:8080

package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type test_struct struct {
    Test string
}

func parseGhPost(rw http.ResponseWriter, request *http.Request) {
    decoder := json.NewDecoder(request.Body)

    var t test_struct
    err := decoder.Decode(&t)

    if err != nil {
        panic(err)
    }

    fmt.Println(t.Test)
}

func main() {
    http.HandleFunc("/", parseGhPost)
    panic(http.ListenAndServe(":8081", nil))
}
