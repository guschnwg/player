package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	colly "github.com/gocolly/colly/v2"
)

type songData struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// TestSpotify ...
func TestSpotify(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		query = "http://go-colly.org/"
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
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit(query)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"results": songs,
	})
}
