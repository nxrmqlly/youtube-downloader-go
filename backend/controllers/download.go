package controllers

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"unicode"

	"youtube-downloader-go/backend/models"
	"youtube-downloader-go/backend/utils"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func slugify(input string) string {
	// Remove non-ASCII characters
	cleaned := strings.Map(func(r rune) rune {
		if r > unicode.MaxASCII || !unicode.IsPrint(r) {
			return -1 // remove character
		}
		return r
	}, input)

	// Replace spaces with underscores
	cleaned = strings.ReplaceAll(cleaned, " ", "_")

	// Remove characters that are invalid in file names
	reg, _ := regexp.Compile(`[^a-zA-Z0-9._-]`)
	cleaned = reg.ReplaceAllString(cleaned, "")

	// Ensure the filename is not empty
	if len(cleaned) == 0 {
		cleaned = "default_filename"
	}

	return cleaned
}

// Download the video by the url specified
func (a *App) DownloadVideo(videoInfo models.VideoInfo, opts models.DownloadOptions) (string, error) {
	//* check if the url is valid

	_, ytdlpPath := utils.YtDlpSetup()

	cmdArgs := []string{videoInfo.URL}
	if opts.Type == "audio" {
		fmt.Println("Downloading Audio")

		cmdArgs = append(
			cmdArgs,
			"--extract-audio",
			"--audio-format",
			opts.Extension,
		)
	} else if opts.Type == "video" {
		fmt.Println("Downloading Video")

		cmdArgs = append(
			cmdArgs,
			"-f",
			fmt.Sprintf("bestvideo[height<=%d]+bestaudio", opts.Resolution),
			"--remux-video",
			opts.Extension,
		)

	}

	pathToSave, err := runtime.SaveFileDialog(
		a.ctx, runtime.SaveDialogOptions{
			Title:           "Save File",
			DefaultFilename: slugify(videoInfo.Title) + "." + opts.Extension,
		},
	)
	if err != nil {
		return "", fmt.Errorf("error opening save file dialog: %v", err)
	}

	cmdArgs = append(cmdArgs, "-o", pathToSave)
	cmd := exec.Command(ytdlpPath, cmdArgs...)
	output, err := utils.RunCmd(cmd)

	if err != nil {
		fmt.Println(string(output))
		return "", fmt.Errorf("error executing yt-dlp: %v\nOutput: %s", err, string(output))
	}

	return "Done downloading", nil
}
