package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/nfnt/resize"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

// Fetch ...
func Fetch(url string, dst interface{}) error {
	client := http.Client{
		Timeout: time.Second * 10, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		return getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}

	jsonErr := json.Unmarshal(body, dst)
	if jsonErr != nil {
		return jsonErr
	}

	return nil
}

func menu(home string) app.UI {
	return app.Div().Body(
		app.A().Href("/").Text(home),
		app.A().Href("/foo").Text("Foo!"),
		app.A().Href("/youtube").Text("Youtube!"),
	)
}

type menuAsCompo struct {
	Home string

	advice string
	value  string

	app.Compo
}

func (h *menuAsCompo) OnMount(ctx app.Context) {
	resp := struct {
		Slip struct {
			ID     int    `json:"id"`
			Advice string `json:"advice"`
		} `json:"slip"`
	}{}

	err := Fetch("https://api.adviceslip.com/advice", &resp)
	if err != nil {
		return
	}

	h.advice = resp.Slip.Advice
	h.Update()
}

func (h *menuAsCompo) Render() app.UI {
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

type pixel struct {
	r, g, b, a uint32
}

type imageHandler struct {
	image image.Image

	pixels [][]pixel

	scale int

	pixelsScaled [][]pixel

	app.Compo
}

func (h *imageHandler) Render() app.UI {
	return app.Div().Body(
		app.Input().
			Type("file").
			OnChange(func(ctx app.Context, e app.Event) {
				file := ctx.JSSrc.Get("files").Get("0")
				reader := app.Window().Get("FileReader").New()
				reader.Call("addEventListener", "loadend", app.FuncOf(func(this app.Value, args []app.Value) interface{} {
					if args[0].Get("target").Get("readyState").Int() == 2 {
						result := args[0].Get("target").Get("result")

						uint8Array := app.Window().Get("Uint8Array").New(result)

						file := make([]byte, uint8Array.Get("length").Int())
						app.CopyBytesToGo(file, uint8Array)

						r := bytes.NewReader(file)
						imgCfg, _, err := image.DecodeConfig(r)

						if err != nil {
							fmt.Println(err)

							return nil
						}

						r.Seek(0, 0)

						img, _, err := image.Decode(r)
						h.image = img

						h.pixels = make([][]pixel, imgCfg.Height)

						for y := 0; y < imgCfg.Height; y++ {
							h.pixels[y] = make([]pixel, imgCfg.Width)

							for x := 0; x < imgCfg.Width; x++ {
								r, g, b, a := img.At(x, y).RGBA()

								h.pixels[y][x] = pixel{r, g, b, a}
							}
						}

						h.Update()
					}
					return nil
				}))
				reader.Call("readAsArrayBuffer", file)
			}),
		app.Div().Body(
			app.Range(h.pixels).Slice(func(i int) app.UI {
				return app.Div().Style("height", "1px").Style("width", fmt.Sprintf("%dpx", len(h.pixels[i]))).Body(
					app.Range(h.pixels[i]).Slice(func(j int) app.UI {
						return app.Span().
							Style("display", "inline-block").
							Style("height", "1px").
							Style("width", "1px").
							Style("background", fmt.Sprintf("rgba(%d,%d,%d,%d)", h.pixels[i][j].r*255/h.pixels[i][j].a, h.pixels[i][j].g*255/h.pixels[i][j].a, h.pixels[i][j].b*255/h.pixels[i][j].a, h.pixels[i][j].a/255))
					}),
				)
			}),
		),
		app.Div().Body(
			app.Input().Type("range").Value(h.scale).Max(20).Min(1).OnChange(func(ctx app.Context, e app.Event) {
				h.scale = ctx.JSSrc.Get("valueAsNumber").Int()

				resizedImage := resize.Resize(uint(h.image.Bounds().Dy())*uint(h.scale), uint(h.image.Bounds().Dx())*uint(h.scale), h.image, resize.Lanczos3)

				h.pixelsScaled = make([][]pixel, resizedImage.Bounds().Dy())

				for y := 0; y < resizedImage.Bounds().Dy(); y++ {
					h.pixelsScaled[y] = make([]pixel, resizedImage.Bounds().Dx())

					for x := 0; x < resizedImage.Bounds().Dx(); x++ {
						r, g, b, a := resizedImage.At(x, y).RGBA()

						h.pixelsScaled[y][x] = pixel{r, g, b, a}
					}
				}

				h.Update()
			}),
			app.Span().Text(h.scale),
		),
		app.Div().Body(
			app.Range(h.pixels).Slice(func(i int) app.UI {
				return app.Div().Style("height", fmt.Sprintf("%dpx", h.scale)).Style("width", fmt.Sprintf("%dpx", len(h.pixels[i])*h.scale)).Body(
					app.Range(h.pixels[i]).Slice(func(j int) app.UI {
						return app.Span().
							Style("display", "inline-block").
							Style("height", fmt.Sprintf("%dpx", h.scale)).
							Style("width", fmt.Sprintf("%dpx", h.scale)).
							Style("background", fmt.Sprintf("rgba(%d,%d,%d,%d)", h.pixels[i][j].r*255/h.pixels[i][j].a, h.pixels[i][j].g*255/h.pixels[i][j].a, h.pixels[i][j].b*255/h.pixels[i][j].a, h.pixels[i][j].a/255))
					}),
				)
			}),
		),
		app.Div().Body(
			app.Range(h.pixelsScaled).Slice(func(i int) app.UI {
				return app.Div().Style("height", "1px").Style("width", fmt.Sprintf("%dpx", len(h.pixelsScaled[i]))).Body(
					app.Range(h.pixelsScaled[i]).Slice(func(j int) app.UI {
						return app.Span().
							Style("display", "inline-block").
							Style("height", "1px").
							Style("width", "1px").
							Style("background", fmt.Sprintf("rgba(%d,%d,%d,%d)", h.pixelsScaled[i][j].r*255/h.pixelsScaled[i][j].a, h.pixelsScaled[i][j].g*255/h.pixelsScaled[i][j].a, h.pixelsScaled[i][j].b*255/h.pixelsScaled[i][j].a, h.pixelsScaled[i][j].a/255))
					}),
				)
			}),
		),
	)
}

type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Hello World!"),
		menu("Home!"),
		&menuAsCompo{Home: "Home!"},
	)
}

type foo struct {
	Value string

	app.Compo
}

func (h *foo) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Foo! "+h.Value),
		app.Input().
			Value(h.Value).
			OnKeyup(func(ctx app.Context, e app.Event) {
				h.Value = ctx.JSSrc.Get("value").String()
				h.Update()
			}),
		menu(h.Value),
		&menuAsCompo{Home: h.Value},

		&imageHandler{},
	)
}

func main() {
	app.Route("/", &hello{})
	app.Route("/game", &game{})
	app.Route("/youtube", &youtube{})
	app.Route("/foo", &foo{})
	app.Run()
}
