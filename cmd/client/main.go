package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"

	"github.com/guschnwg/player/pkg/client"
)

func main() {
	app.Route("/", &client.Home{})
	app.Route("/youtube", &client.Youtube{})
	app.Run()
}
