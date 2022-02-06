package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Hello World!")
	})
	http.ListenAndServe(":8080", nil)
}
