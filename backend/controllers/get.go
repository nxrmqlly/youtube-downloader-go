package controllers

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"youtube-downloader-go/backend/models"
	"youtube-downloader-go/backend/utils"
)

func (a *App) GetVideoInfo(videoURL string) (models.VideoInfo, error) {
	_, ytDlpPath := utils.YtDlpSetup()

	cmd := exec.Command(ytDlpPath, "--dump-json", videoURL)

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return models.VideoInfo{Valid: false}, fmt.Errorf("1 error executing yt-dlp: %w", err)
	}

	// Parse the JSON output
	var videoInfo models.VideoInfo
	err = json.Unmarshal(output, &videoInfo)
	if err != nil {
		return models.VideoInfo{Valid: false}, fmt.Errorf("error parsing JSON: %w", err)
	}

	videoInfo.Valid = true

	return videoInfo, nil
}
