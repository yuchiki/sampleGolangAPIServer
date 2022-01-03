package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var VERSION = ":VERSION:"

type helloResponse struct {
	Message string
	Version string
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(helloResponse{
		Message: "hello",
		Version: VERSION,
	})
}
