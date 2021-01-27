package server

import (
	"encoding/json"
	"net/http"
	"strings"

	colly "github.com/gocolly/colly/v2"
)

// TestLyrics ...
func TestLyrics(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		query = "Teste"
	}

	lyrics := []string{}

	c := colly.NewCollector(colly.MaxDepth(1))

	c.OnHTML("#wrapper > div.wrapper-inner > div.coltwo-wide-2 > div:nth-child(5) > a", func(e *colly.HTMLElement) {
		c.Visit(e.Attr("href"))
	})
	c.OnHTML("#songLyricsDiv", func(e *colly.HTMLElement) {
		lyrics = strings.Split(e.DOM.Text(), "\n")
	})

	c.OnRequest(func(r *colly.Request) {
		print("Visiting " + r.URL.String())
	})

	URL := "http://www.songlyrics.com/index.php?section=search&searchW=" + strings.ReplaceAll(query, " ", "+") + "&submit=Search&searchIn1=artist&searchIn2=album&searchIn3=song&searchIn4=lyrics"
	err := c.Visit(URL)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"results": lyrics,
	})
}
