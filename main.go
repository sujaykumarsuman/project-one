package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sujaykumarsuman/project-one/handler"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.Default().Flags())
	hh := handler.NewHello(logger)
	gh := handler.NewGoodbye(logger)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/bye", gh)

	http.ListenAndServe(":8080", sm)
}
