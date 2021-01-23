package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"os/exec"
	"strings"

	"github.com/guschnwg/player/pkg/shared"
)

// Test ...
func Test(w http.ResponseWriter, r *http.Request) {
	songURL := r.URL.Query().Get("song_url")

	if songURL == "" {
		songURL = "https://www.youtube.com/watch?v=lgjY-lVtJZA"
	}

	w.Header().Set("Content-Type", "application/json")

	cmd := exec.Command("youtube-dl", songURL, "--skip-download", "--dump-json", "-4")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	var songData map[string]json.RawMessage
	err = json.Unmarshal(out.Bytes(), &songData)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"results": songData,
	})
}

// Search ...
func Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		query = "cats"
	}

	w.Header().Set("Content-Type", "application/json")

	cmdData := exec.Command("youtube-dl", "--default-search", "ytsearch1:", "--skip-download", "--dump-json", "-4", url.QueryEscape(query))
	var out bytes.Buffer
	cmdData.Stdout = &out
	err := cmdData.Run()
	if err != nil {
		w.WriteHeader(500)
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

	var songs []shared.SongData
	for _, item := range results {
		var song shared.SongData

		if err = json.Unmarshal([]byte(item), &song); err == nil {
			songs = append(songs, song)
		} else {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"results": songs,
	})
}
