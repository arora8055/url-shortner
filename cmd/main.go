package main

import (
	"fmt"
	"log"
	"net/http"

	"url-shortner/api"
)

func main() {
	http.HandleFunc("/shorten", api.ShortenURLHandler)
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
