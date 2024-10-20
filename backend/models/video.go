package models

type VideoInfo struct {
	URL         string  `json:"webpage_url"`
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	FPS         float32 `json:"fps"`
	Uploader    string  `json:"uploader"`
	UploaderURL string  `json:"uploader_url"`
	Duration    float32 `json:"duration"` // Duration in seconds
	ViewCount   int     `json:"view_count"`
	Thumbnail   string  `json:"thumbnail"`
	Height      int     `json:"height"`
	Width       int     `json:"width"`
}
