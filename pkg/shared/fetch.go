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
