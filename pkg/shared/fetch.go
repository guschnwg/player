package shared

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Fetch ...
func Fetch(url string, dst interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

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

// FetchSongs ...
func FetchSongs(query string) ([]SongData, error) {
	var response struct {
		Results []SongData `json:"results"`
	}
	err := Fetch("/search?query="+query, &response)
	if err != nil {
		return []SongData{}, err
	}

	return response.Results, nil
}

// FetchSpotifyPlaylist ...
func FetchSpotifyPlaylist(query string) ([]SpotifyPlaylistSong, error) {
	var response struct {
		Results []SpotifyPlaylistSong `json:"results"`
	}
	err := Fetch("/spotify/test?query="+query, &response)
	if err != nil {
		return []SpotifyPlaylistSong{}, err
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
