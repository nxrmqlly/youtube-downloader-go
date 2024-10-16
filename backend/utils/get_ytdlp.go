package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const YtDlpGitgubRelease = "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp.exe"

// If yt-dlp isn't installed yet, download it locally.
func downloadYtDlpBinary() (string, error) {
	folderPath := "./bin-deps"

	// If the folder doesn't exist, create it
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(folderPath, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return "", err
		}
	}

	// create a new file in the folder

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pathToCreate := filepath.Join(cwd, "bin-deps", "yt-dlp.exe")

	out, err := os.Create(pathToCreate)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	resp, err := http.Get(YtDlpGitgubRelease)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	io.Copy(out, resp.Body)

	return pathToCreate, nil
}

// Get yt-dlp if not exists, returns cwd and path of yt-dlp
func YtDlpSetup() (string, string) {
	var ytDlpPath string

	var cwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	// check if ./bin-deps/yt-dlp.exe exists
	if _, err := os.Stat("./bin-deps/yt-dlp.exe"); os.IsNotExist(err) {
		ytDlpPath, _ = downloadYtDlpBinary()
	} else {
		ytDlpPath = filepath.Join(cwd, "bin-deps", "yt-dlp.exe")

	}
	return cwd, ytDlpPath
}
