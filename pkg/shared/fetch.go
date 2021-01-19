package shared

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Fetch ...
func Fetch(url string, dst interface{}) error {
	client := http.Client{
		Timeout: time.Second * 10, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		return getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

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
