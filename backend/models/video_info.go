package models

type VideoInfo struct {
	Valid     bool
	Title     string `json:"title"`
	Uploader  string `json:"uploader"`
	Duration  int    `json:"duration"` // Duration in seconds
	ViewCount int    `json:"view_count"`
}
