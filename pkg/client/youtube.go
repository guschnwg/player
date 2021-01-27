package client

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"

	"github.com/guschnwg/player/pkg/shared"
)

// Youtube ...
type Youtube struct {
	app.Compo

	query string

	songs []shared.SongData
	err   error
}

// Render ...
func (c *Youtube) Render() app.UI {
	err := ""
	if c.err != nil {
		err = c.err.Error()
	}

	return app.Div().Body(
		app.H1().Text("Youtube!!!"),

		app.Div().
			Style("display", "flex").
			Style("flex-direction", "column").
			Body(
				app.Label().For("query").Text("Query"),
				app.Input().
					ID("query").
					Value(c.query).
					OnKeyup(func(ctx app.Context, e app.Event) {
						c.query = ctx.JSSrc.Get("value").String()
						c.Update()
					}),
			),

		app.Button().OnClick(func(ctx app.Context, e app.Event) {
			go c.searchGoRoutine()
		}).Text("Search With GoRoutine"),

		app.Button().OnClick(func(ctx app.Context, e app.Event) {
			c.searchNormal()
		}).Text("Search Normal"),

		app.If(
			len(c.songs) > 0,
			app.Div().Body(
				app.Range(c.songs).Slice(func(i int) app.UI {
					return &Song{Song: c.songs[i]}
				}),
			),
		),

		app.If(err != "", app.Span().Text(err)),
	)
}

func (c *Youtube) searchNormal() {
	songs, err := FetchSongs(c.query)

	c.err = err
	c.songs = songs

	c.Update()
}

func (c *Youtube) searchGoRoutine() {
	songs, err := FetchSongs(c.query)

	app.Dispatch(func() {
		c.err = err
		c.songs = songs

		c.Update()
	})
}
