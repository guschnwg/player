package shared

// SongData ...
type SongData struct {
	ID         string `json:"id"`
	Uploader   string `json:"uploader"`
	UploaderID string `json:"uploader_id"`

	Title string `json:"title"`

	Thumbnail string `json:"thumbnail"`

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
		ID     string `json:"format_id"`
		URL    string `json:"url"`
		Ext    string `json:"ext"`
		ACodec string `json:"acodec"`
	} `json:"formats"`
}

// SpotifyPlaylistSong ...
type SpotifyPlaylistSong struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// BeatportData ...
type BeatportData struct {
	Genre string `json:"genre"`
	BPM   string `json:"bpm"`
}
