package client

import (
	"github.com/guschnwg/player/pkg/shared"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Song ...
type Song struct {
	app.Compo

	Song shared.SongData
}

// Render ...
func (c *Song) Render() app.UI {
	if c.Song.ID == "" {
		return app.Div().Text("Loading")
	}

	format := c.Song.Formats[len(c.Song.Formats)-1]

	return app.Div().Body(
		app.H4().Text(c.Song.Title),

		app.Video().Controls(true).Poster(c.Song.Thumbnail).Src(format.URL),
	)
}
