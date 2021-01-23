package client

import (
	"net/url"
	"strings"
	"syscall/js"

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

	URL := format.URL
	parsedURL, _ := url.Parse(URL)

	parsedURL.RawQuery += "&host=https://" + parsedURL.Host + "/"

	// parsedURL.Scheme = "http"
	parsedURL.Scheme = strings.Replace(js.Global().Get("location").Get("protocol").String(), ":", "", 1)

	// parsedURL.Host = "localhost:8000"
	parsedURL.Host = js.Global().Get("location").Get("host").String()

	return app.Div().Body(
		app.H4().Text(c.Song.Title),

		app.Video().Controls(true).Poster(c.Song.Thumbnail).Src(parsedURL.String()),
	)
}
