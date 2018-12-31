package main

import (
	"fmt"
	"log"
	"net/http"
)

func welcome(respose http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(respose, "Welcome to Golang API in Docker Container.")
}

func main() {
	http.HandleFunc("/", welcome)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
