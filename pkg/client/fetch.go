package client

import (
	"context"
	"encoding/json"
	"time"

	"github.com/guschnwg/player/pkg/shared"
	fetch "marwan.io/wasm-fetch"
)

// Fetch ...
func Fetch(url string, dst interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	resp, err := fetch.Fetch(url, &fetch.Opts{
		Method: fetch.MethodGet,
		Signal: ctx,
	})
	if err != nil {
		return err
	}

	jsonErr := json.Unmarshal(resp.Body, dst)
	if jsonErr != nil {
		return jsonErr
	}

	return nil
}

// FetchSongs ...
func FetchSongs(query string) ([]shared.SongData, error) {
	var response struct {
		Results []shared.SongData `json:"results"`
	}
	err := Fetch("/search?query="+query, &response)
	if err != nil {
		return []shared.SongData{}, err
	}

	return response.Results, nil
}

// FetchSpotifyPlaylist ...
func FetchSpotifyPlaylist(query string) ([]shared.SpotifyPlaylistSong, error) {
	var response struct {
		Results []shared.SpotifyPlaylistSong `json:"results"`
	}
	err := Fetch("/spotify/test?query="+query, &response)
	if err != nil {
		return []shared.SpotifyPlaylistSong{}, err
	}

	return response.Results, nil
}

// FetchLyrics ...
func FetchLyrics(query string) ([]string, error) {
	var response struct {
		Results []string `json:"results"`
	}
	err := Fetch("/lyrics/test?query="+query, &response)
	if err != nil {
		return []string{}, err
	}

	return response.Results, nil
}

// FetchBeatport ...
func FetchBeatport(query string) (shared.BeatportData, error) {
	var response struct {
		Results shared.BeatportData `json:"results"`
	}
	err := Fetch("/beatport/test?query="+query, &response)
	if err != nil {
		return shared.BeatportData{}, err
	}

	return response.Results, nil
}

// FetchAdvice ...
func FetchAdvice(query string) (string, error) {
	var response struct {
		Slip struct {
			ID     int    `json:"id"`
			Advice string `json:"advice"`
		} `json:"slip"`
	}

	err := Fetch("https://api.adviceslip.com/advice", &response)
	if err != nil {
		return "", err
	}

	return response.Slip.Advice, nil
}
