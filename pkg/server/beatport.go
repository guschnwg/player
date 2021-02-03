package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	colly "github.com/gocolly/colly/v2"
	"github.com/guschnwg/player/pkg/shared"
)

const baseURL = "https://www.beatport.com"

// TestBeatport ...
func TestBeatport(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		query = "Eats Everything - Space Raiders (Charlotte de Witte Remix)"
	}

	c := colly.NewCollector(colly.MaxDepth(1))

	data := shared.BeatportData{}

	visited := false
	c.OnHTML(".bucket-items > .bucket-item:first-child > .buk-track-meta-parent > .buk-track-title > a", func(e *colly.HTMLElement) {
		if !visited {
			visited = true
			c.Visit(baseURL + e.Attr("href"))
		}
	})

	c.OnHTML(".interior-track-content-list > .interior-track-genre > .value > a", func(e *colly.HTMLElement) {
		fmt.Println("FOUND " + e.Text)
		data.Genre = e.Text
	})

	c.OnHTML(".interior-track-content-list > .interior-track-bpm > .value", func(e *colly.HTMLElement) {
		fmt.Println("FOUND " + e.Text)
		data.BPM = e.Text
	})

	err := c.Visit(baseURL + "/search?q=" + url.QueryEscape(query))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"results": data,
	})
}
