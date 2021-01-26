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

	Song    shared.SongData
	Lyrics  []string
	OnEnded app.EventHandler

	isPlaying bool

	VideoElementID string
}

// Render ...
func (c *Song) Render() app.UI {
	URL := c.getURL()

	video := app.Video().ID(c.VideoElementID).Class("h-full w-full").Controls(true).Poster(c.Song.Thumbnail).Src(URL)

	video.OnPlaying(func(ctx app.Context, e app.Event) {
		c.isPlaying = true

		c.Update()
	})
	video.OnEnded(func(ctx app.Context, e app.Event) {
		c.isPlaying = false

		if c.OnEnded != nil {
			c.OnEnded(ctx, e)
		}

		c.Update()
	})

	extraClasses := " "
	if c.isPlaying {
		extraClasses += "pulse-border "
	}

	return app.Div().Class("rounded shadow border border-gray-800 flex my-5 h-56"+extraClasses).Body(
		app.Div().Class("w-1/3").Body(
			video,
		),
		app.Div().Class("flex flex-col flex-1 p-4").Body(
			app.H5().Class("font-bold mb-4 mt-0 text-white").Text(c.Song.Title),

			app.Div().
				Class("overflow-y-auto h-full").
				Body(
					app.Range(c.Lyrics).Slice(func(i int) app.UI {
						return app.Div().Class("text-white").Style("min-height", "24px").Text(c.Lyrics[i])
					}),
				),
		),
	)
}

func (c *Song) getURL() string {
	URL := c.Song.Formats[len(c.Song.Formats)-1].URL
	parsedURL, _ := url.Parse(URL)
	parsedURL.RawQuery += "&host=https://" + parsedURL.Host + "/"
	parsedURL.Scheme = strings.Replace(js.Global().Get("location").Get("protocol").String(), ":", "", 1) // "http"
	parsedURL.Host = js.Global().Get("location").Get("host").String()                                    // "localhost:8000"
	return parsedURL.String()
}
