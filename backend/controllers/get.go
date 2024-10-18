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

	cmd := exec.Command(ytDlpPath, "-J", videoURL)

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return models.VideoInfo{}, fmt.Errorf("yt-dlp error: %w", err)
	}

	// Parse the JSON output
	var videoInfo models.VideoInfo
	err = json.Unmarshal(output, &videoInfo)
	if err != nil {
		return models.VideoInfo{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return videoInfo, nil
}
