package models

type VideoInfo struct {
	URL         string `json:"webpage_url"`
	ID          string `json:"id"`
	Title       string `json:"title"`
	FPS         int    `json:"fps"`
	Uploader    string `json:"uploader"`
	UploaderURL string `json:"uploader_url"`
	Duration    int    `json:"duration"` // Duration in seconds
	ViewCount   int    `json:"view_count"`
	Thumbnail   string `json:"thumbnail"`
}
