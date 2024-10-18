package models

type DownloadOptions struct {
	Extension  string `json:"ext"`  // mp3, mp4, avi, opus ...
	Type       string `json:"type"` // audio or video
	Resolution int    `json:"res"`
}
