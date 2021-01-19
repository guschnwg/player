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
	err   string
}

// Render ...
func (c *Youtube) Render() app.UI {
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

		app.If(c.err != "", app.Span().Text(c.err)),
	)
}

func (c *Youtube) searchNormal() {
	var response map[string][]shared.SongData
	err := shared.Fetch("/search?query="+c.query, &response)

	if err != nil {
		c.err = err.Error()
		c.Update()
		return
	}

	c.songs = response["results"]
	c.Update()
}

func (c *Youtube) searchGoRoutine() {
	var response map[string][]shared.SongData
	err := shared.Fetch("/search?query="+c.query, &response)

	if err != nil {
		app.Dispatch(func() {
			c.err = err.Error()
			c.Update()
		})
		return
	}

	app.Dispatch(func() {
		c.songs = response["results"]
		c.Update()
	})
}
