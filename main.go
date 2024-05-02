package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Received a request on /")
    fmt.Fprint(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloWorldHandler)
    log.Println("Starting server on port 2024")
    err := http.ListenAndServe(":2024", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
