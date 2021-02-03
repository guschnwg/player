package main

import (
	"log"
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/v7/pkg/app"

	"github.com/guschnwg/player/pkg/server"
)

func main() {
	port := os.Getenv("PORT")

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Title:       "Hello World!",
		Styles: []string{
			"https://unpkg.com/tailwindcss@^2/dist/tailwind.min.css",
			"/web/main.css",
		},
	})

	http.HandleFunc("/test", server.Test)
	http.HandleFunc("/search", server.Search)
	http.HandleFunc("/spotify/test", server.TestSpotify)
	http.HandleFunc("/lyrics/test", server.TestLyrics)
	http.HandleFunc("/beatport/test", server.TestBeatport)

	server.BindProxy()

	log.Println("Server running on port: " + port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
