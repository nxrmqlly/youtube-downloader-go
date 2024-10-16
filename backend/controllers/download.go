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
func (a *App) DownloadVideo(opts models.DownloadOptions) (string, error) {
	//* check if the url is valid

	_, ytdlpPath := utils.YtDlpSetup()

	var url = opts.URL

	pathToSave, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{Title: "Save File"})

	if err != nil {
		return err.Error(), err
	}

	cmdArgs := []string{"-o", pathToSave, url}

	// Set download type
	if opts.DownloadType == "video" {
		cmdArgs = append(cmdArgs, "--format", fmt.Sprintf("bestvideo[ext=%s]", opts.FileType))
	} else if opts.DownloadType == "audio" {
		cmdArgs = append(cmdArgs, "--format", fmt.Sprintf("bestaudio[ext=%s]", opts.FileType))
	} else {
		err := fmt.Errorf("invalid download type: %s", opts.DownloadType)
		return err.Error(), err
	}

	// Set resolution and quality options if provided
	if opts.Resolution != "" {
		cmdArgs = append(cmdArgs, "--format", fmt.Sprintf("%s+%s", opts.Resolution, opts.Quality))
	}

	// Execute yt-dlp command
	cmd := exec.Command(ytdlpPath, cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		err := fmt.Errorf("error executing yt-dlp: %v\nOutput: %s", err, string(output))
		return err.Error(), err
	}

	// Print useful information
	fmt.Printf("Downloaded %s from %s to %s\n", opts.DownloadType, url, pathToSave)
	fmt.Println(string(output))

	return "Done downloading", nil
}
