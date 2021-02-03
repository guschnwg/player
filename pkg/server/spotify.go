package server

import (
	"encoding/json"
	"net/http"

	colly "github.com/gocolly/colly/v2"
)

type songData struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// TestSpotify ...
func TestSpotify(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)

	query := r.URL.Query().Get("query")
	if query == "" {
		query = "https://open.spotify.com/playlist/30mIdIfINRKeT4QbJOk0Qf"
	}

	c := colly.NewCollector()

	songs := []songData{}

	c.OnHTML(".tracklist-row > .tracklist-col.name > .track-name-wrapper", func(e *colly.HTMLElement) {
		children := e.DOM.Children()

		songContainer := children.First()
		artistContainer := children.Last().Children().First()

		songs = append(songs, songData{
			songContainer.Text(),
			artistContainer.Text(),
		})
	})

	c.OnRequest(func(r *colly.Request) {
		print("Visiting " + r.URL.String())
	})

	err := c.Visit(query)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"results": songs,
	})
}
