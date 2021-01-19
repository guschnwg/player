package client

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"

	"github.com/guschnwg/player/pkg/shared"
)

// Youtube ...
type Youtube struct {
	app.Compo

	song shared.SongData

	songURL string
}

// Render ...
func (h *Youtube) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Youtube!!!"),
	)
}
