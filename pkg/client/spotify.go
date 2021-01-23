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
		app.Button().
			Text("Search playlist").
			OnClick(func(ctx app.Context, e app.Event) {
				c.crawlSpotify()
			}),

		app.If(
			len(c.songs) > 0,
			app.Div().Body(
				app.Range(c.songs).Slice(func(i int) app.UI {
					return &SpotifySong{Song: c.songs[i]}
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

// // // // //

// SpotifySong ...
type SpotifySong struct {
	Song shared.SpotifyPlaylistSong

	err  error
	song shared.SongData

	app.Compo
}

// FetchSong ...
func (c *SpotifySong) FetchSong() {
	songs, err := shared.FetchSongs(c.Song.Title + " - " + c.Song.Artist)

	app.Dispatch(func() {
		c.err = err
		if len(songs) > 0 {
			c.song = songs[0]
		}
		c.Update()
	})
}

// OnMount ...
func (c *SpotifySong) OnMount(ctx app.Context) {
	// This allows me to fire all at the same time!
	go c.FetchSong()
}

// Render ...
func (c *SpotifySong) Render() app.UI {
	return app.Div().Body(
		app.Div().Text(c.Song.Title+" - "+c.Song.Artist),

		&Song{Song: c.song},
	)
}
