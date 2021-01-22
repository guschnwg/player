package client

import (
	"github.com/guschnwg/player/pkg/shared"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Spotify ...
type Spotify struct {
	query string

	err   error
	songs []shared.SpotifyPlaylistSong

	app.Compo
}

// Render ...
func (c *Spotify) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Hello World! "+c.query),
		app.Input().
			Value(c.query).
			OnKeyup(func(ctx app.Context, e app.Event) {
				c.query = ctx.JSSrc.Get("value").String()
				c.Update()
			}),
		app.Button().OnClick(func(ctx app.Context, e app.Event) {
			c.crawlSpotify()
		}),

		app.If(
			len(c.songs) > 0,
			app.Ul().Body(
				app.Range(c.songs).Slice(func(i int) app.UI {
					return app.Li().Text(c.songs[i].Title + " - " + c.songs[i].Artist)
				}),
			),
		),
	)
}

func (c *Spotify) crawlSpotify() {
	songs, err := shared.FetchSpotifyPlaylist(c.query)

	c.err = err
	c.songs = songs

	c.Update()
}
