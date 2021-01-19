package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

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

func main() {
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Title:       "Hello World!",
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		songURL := r.URL.Query().Get("song_url")
		fmt.Println(songURL)

		if songURL == "" {
			songURL = "https://www.youtube.com/watch?v=lgjY-lVtJZA"
		}

		fmt.Println(songURL)

		w.Header().Set("Content-Type", "application/json")

		cmd := exec.Command("youtube-dl", songURL, "--skip-download", "--dump-json", "-4")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
		}

		var songData map[string]json.RawMessage
		err = json.Unmarshal(out.Bytes(), &songData)
		if err != nil {
			fmt.Println(err.Error())
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"results": songData,
		})
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")

		w.Header().Set("Content-Type", "application/json")

		cmd := exec.Command("youtube-dl", "--default-search", "ytsearch2:", "--skip-download", "--dump-json", "-4", query)

		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		stdout := string(out.Bytes())
		results := strings.Split(stdout, "\n")
		if len(results) == 0 {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"results": []string{},
			})
		}

		results = results[0 : len(results)-1]

		var songs []SongData
		for _, item := range results {
			var song SongData

			if err = json.Unmarshal([]byte(item), &song); err == nil {
				songs = append(songs, song)
			} else {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error": err.Error(),
				})
			}
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"results": songs,
		})
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
