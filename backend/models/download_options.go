package models

type DownloadOptions struct {
	URL          string // e.g., "https://www.youtube.com/watch?v=1234"
	DownloadType string // "video" or "audio"
	FileType     string // e.g., "mp4", "mp3"
	Resolution   string // e.g., "720p", "1080p"
	Quality      string // e.g., "best", "worst"
}
