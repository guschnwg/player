package client

import (
	"fmt"

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
	return app.Div().
		Class("p-10 h-full").
		Body(
			app.H1().Class("text-2xl text-white mb-5").Text("Add here a Spotify playlist!"),

			app.Div().
				Class("flex").
				Body(
					app.Input().
						Class("bg-black text-white border-solid border border-green-500 py-2 px-4 rounded-full rounded-r-none flex-1").
						Placeholder("https://open.spotify.com/playlist/30mIdIfINRKeT4QbJOk0Qf").
						Value(c.query).
						OnKeyup(func(ctx app.Context, e app.Event) {
							c.query = ctx.JSSrc.Get("value").String()
							c.Update()
						}),
					app.Button().
						Class("mr-2 bg-green-500 text-green-100 block py-2 px-8 rounded-full rounded-l-none").
						Text("Search playlist").
						OnClick(func(ctx app.Context, e app.Event) {
							go c.crawlSpotify()
						}),
				),

			c.RenderSongs(),
		)
}

// RenderSongs ...
func (c *Spotify) RenderSongs() app.UI {
	if len(c.songs) == 0 {
		return nil
	}

	return app.Div().Body(
		app.Range(c.songs).Slice(func(i int) app.UI {
			return &SpotifySong{
				VideoElementID: fmt.Sprintf("my-video-%d", i),
				Song:           c.songs[i],
				OnEnded: func(ctx app.Context, e app.Event) {
					if i+1 < len(c.songs) {
						elem := app.Window().
							Get("document").
							Call("getElementById", fmt.Sprintf("my-video-%d", i+1))
						elem.Call("play")
					}
				},
			}
		}),
	)
}

func (c *Spotify) crawlSpotify() {
	songs, err := shared.FetchSpotifyPlaylist(c.query)

	app.Dispatch(func() {
		c.err = err
		c.songs = songs

		c.Update()
	})
}

// // // // //

// SpotifySong ...
type SpotifySong struct {
	VideoElementID string
	Song           shared.SpotifyPlaylistSong
	OnEnded        app.EventHandler

	err    error
	song   shared.SongData
	lyrics []string

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

// FetchLyrics ...
func (c *SpotifySong) FetchLyrics() {
	lyrics, err := shared.FetchLyrics(c.Song.Title + " - " + c.Song.Artist)

	app.Dispatch(func() {
		c.err = err
		c.lyrics = lyrics
		c.Update()
	})
}

// OnMount ...
func (c *SpotifySong) OnMount(ctx app.Context) {
	// This allows me to fire all at the same time!
	go c.FetchSong()
	go c.FetchLyrics()
}

// Render ...
func (c *SpotifySong) Render() app.UI {
	if c.song.ID == "" {
		return app.Div().Text("Loading " + c.Song.Title + " - " + c.Song.Artist)
	}

	return &Song{
		Song:           c.song,
		Lyrics:         c.lyrics,
		VideoElementID: c.VideoElementID,
		OnEnded:        c.OnEnded,
	}
}
