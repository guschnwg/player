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

	format := c.Song.Formats[0]

	for _, f := range c.Song.Formats {
		if f.Ext == "mp4" && f.ACodec != "none" {
			format = f
			break
		}
	}

	return app.Div().Body(
		app.H4().Text(c.Song.Title),

		app.Video().Controls(true).Src(format.URL),
	)
}
