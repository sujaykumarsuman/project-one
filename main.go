package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Hello World!")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("[ERROR]: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		log.Printf("Data: %s", d)
		fmt.Fprintf(w, "Hello %s!\n", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye!")
	})
	http.ListenAndServe(":8080", nil)
}
