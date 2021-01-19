package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v7/pkg/app"

	"github.com/guschnwg/player/pkg/server"
)

func main() {
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Title:       "Hello World!",
	})

	http.HandleFunc("/test", server.Test)

	http.HandleFunc("/query", server.Query)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
