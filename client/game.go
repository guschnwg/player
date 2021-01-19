package main

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type game struct {
	app.Compo
}

func (h *game) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Game!!"),
	)
}
