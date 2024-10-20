package controllers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os/exec"

	"youtube-downloader-go/backend/models"
	"youtube-downloader-go/backend/utils"
)

func (a *App) GetVideoInfo(videoURL string) (models.VideoInfo, error) {

	_, err := url.ParseRequestURI(videoURL)
	if err != nil {
		return models.VideoInfo{}, fmt.Errorf("invalid URL: %w", err)
	}

	_, ytDlpPath := utils.YtDlpSetup()

	cmd := exec.Command(ytDlpPath, "-J", videoURL)

	// Capture the output
	output, err := utils.RunCmd(cmd)
	if err != nil {
		return models.VideoInfo{}, fmt.Errorf("%s; output: %s", err, output)
	}

	// Parse the JSON output
	var videoInfo models.VideoInfo
	err = json.Unmarshal(output, &videoInfo)
	if err != nil {
		return models.VideoInfo{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return videoInfo, nil
}
