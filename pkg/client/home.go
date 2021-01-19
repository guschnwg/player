package client

import "github.com/maxence-charriere/go-app/v7/pkg/app"

// Home ...
type Home struct {
	Value string

	app.Compo
}

// Render ...
func (h *Home) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Hello World! "+h.Value),
		app.Input().
			Value(h.Value).
			OnKeyup(func(ctx app.Context, e app.Event) {
				h.Value = ctx.JSSrc.Get("value").String()
				h.Update()
			}),
		Menu(h.Value),
		&MenuAsCompo{Home: h.Value},
	)
}
