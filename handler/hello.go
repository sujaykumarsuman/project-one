package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World!")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("[ERROR]: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	log.Printf("Data: %s", d)
	fmt.Fprintf(w, "Hello %s!\n", d)
}
