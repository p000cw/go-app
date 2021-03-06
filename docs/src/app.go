// +build wasm

package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type hello struct {
	app.Compo
	name string
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.Main().Body(
			app.H1().
				OnContextMenu(h.OnContextMenu).
				Body(
					app.Text("Hello, "),
					app.If(h.name != "",
						app.Text(h.name),
					).Else(
						app.Text("World"),
					),
				),
			app.Input().
				Value(h.name).
				Placeholder("What is your name?").
				AutoFocus(true).
				OnChange(h.OnInputChange),
		),
	)
}

func (h *hello) OnInputChange(ctx app.Context, e app.Event) {
	h.name = ctx.JSSrc.Get("value").String()
	h.Update()
}

func (h *hello) OnContextMenu(ctx app.Context, e app.Event) {
	e.PreventDefault()

	app.NewContextMenu(
		app.MenuItem().
			Label("hello").
			Keys("cmd+a").
			OnClick(func(app.Context, app.Event) {
				fmt.Println("hello!")
			}),

		app.MenuItem().Separator(),
		app.MenuItem().
			Label("world").
			Keys("cmd+b").
			OnClick(func(app.Context, app.Event) {
				fmt.Println("world!")
			}),
	)
}

func main() {
	app.Route("/", &hello{})
	app.Route("/hello", &hello{})
	app.Run()
}
