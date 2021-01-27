package client

import (
	"github.com/guschnwg/player/pkg/shared"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Menu ...
func Menu(home string) app.UI {
	return app.Div().Body(
		app.A().Href("/").Text(home),
		app.A().Href("/foo").Text("Foo!"),
		app.A().Href("/youtube").Text("Youtube!"),
		app.A().Href("/spotify").Text("Spotify!"),
	)
}

// MenuAsCompo ...
type MenuAsCompo struct {
	Home string

	advice string
	value  string

	app.Compo
}

// OnMount ...
func (h *MenuAsCompo) OnMount(ctx app.Context) {
	advice, _ := shared.FetchAdvice("https://api.adviceslip.com/advice")
	h.advice = advice
	h.Update()
}

// Render ...
func (h *MenuAsCompo) Render() app.UI {
	return app.Div().Body(
		app.A().Href("/").Text(h.Home),
		app.A().Href("/foo").Text("Foo!"),

		app.Input().
			Value(h.value).
			OnKeyup(func(ctx app.Context, e app.Event) {
				h.value = ctx.JSSrc.Get("value").String()
				h.Update()
			}),

		app.P().Text(h.advice+" - "+h.value),
	)
}
