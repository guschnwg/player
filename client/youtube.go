package main

import (
	"encoding/json"
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type song struct {
	Filename string `json:"_filename"`
	Formats  []struct {
		URL string `json:"url"`
	} `json:"formats"`
}

type SongData struct {
	ID         string `json:"id"`
	Uploader   string `json:"uploader"`
	UploaderID string `json:"uploader_id"`

	Title string `json:"title"`

	Thumbnails []struct {
		ID     string `json:"id"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"thumbnails"`

	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	Tags        []string `json:"tags"`

	Formats []struct {
		ID  string `json:"format_id"`
		URL string `json:"url"`
		Ext string `json:"ext"`
	} `json:"formats"`
}

type youtube struct {
	app.Compo

	song song

	songURL string
}

type search struct {
	query string
	songs []SongData

	app.Compo
}

func (h *search) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Search!!!"),
		app.Input().Value(h.query).OnChange(func(ctx app.Context, e app.Event) {
			h.query = ctx.JSSrc.Get("value").String()
			h.songs = searchSong(h.query)
			h.Update()
		}),
		app.Span().Text(h.query),

		app.Range(h.songs).Slice(func(i int) app.UI {
			song := h.songs[i]

			if len(song.Formats) == 0 {
				return nil
			}

			format := song.Formats[0]
			for _, f := range song.Formats {
				if f.Ext == "mp4" {
					format = f
					break
				}
			}

			return app.Video().Controls(true).Src(format.URL)
		}),
	)
}

func getSong(url string) song {
	var songData map[string]json.RawMessage

	err := Fetch("/test?song_url="+url, &songData)
	if err != nil {
		fmt.Println(err.Error())
		return song{}
	}

	results := song{}
	err = json.Unmarshal(songData["results"], &results)

	if err != nil {
		fmt.Println(err.Error())
		return song{}
	}

	return results
}

func searchSong(query string) []SongData {
	var response map[string]json.RawMessage

	err := Fetch("/search?query="+query, &response)
	if err != nil {
		fmt.Println(err.Error())
		return []SongData{}
	}

	songs := []SongData{}
	err = json.Unmarshal(response["results"], &songs)

	if err != nil {
		fmt.Println(err.Error())
		return []SongData{}
	}

	return songs
}

func (h *youtube) OnMount(ctx app.Context) {
	h.song = getSong("")
	h.Update()
}

func (h *youtube) Render() app.UI {
	fmt.Println(h.song.Filename)

	url := ""

	if len(h.song.Formats) > 0 {
		fmt.Println(h.song.Formats[0].URL)
		url = h.song.Formats[0].URL
	}

	return app.Div().Body(
		app.H1().Text("Youtube!!!"),
		app.Input().Value(h.songURL).OnChange(func(ctx app.Context, e app.Event) {
			h.songURL = ctx.JSSrc.Get("value").String()
			h.song = getSong(h.songURL)
			h.Update()
		}),
		app.Span().Text(h.songURL),
		app.Audio().Controls(true).Src(url),
		app.Video().Controls(true).Src(url),

		&search{},
	)
}
