package main

import (
	"log"
	"strconv"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

type Component struct {
	vecty.Core
}

func (c *Component) Copy() vecty.Component {
	cpy := *c
	return &cpy
}

func (c *Component) Render() vecty.ComponentOrHTML {
	return nil
}

//

func main() {
	vecty.SetTitle("Hello Vecty!")
	vecty.RenderBody(&PageView{})
}

// PageView is our main page component.
type PageView struct {
	count int

	Component
}

func (p *PageView) onClick(event *vecty.Event) {
	p.count++
	vecty.Rerender(p)
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	log.Println("OI")
	return elem.Body(
		vecty.Text("Hello Vecty!"),

		elem.Button(
			vecty.Markup(
				prop.Href("#"),
				event.Click(p.onClick).PreventDefault(),
			),

			vecty.Text(strconv.Itoa(p.count)),
		),
	)
}
