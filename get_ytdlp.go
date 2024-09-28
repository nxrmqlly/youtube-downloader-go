package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const YtDlpGitgubRelease = "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp.exe"

// If yt-dlp isn't installed yet, download it locally.
func GetYtDlp() {
	folderPath := "./bin-deps"

	// Check if the folder exists
	_, err := os.Stat(folderPath)

	// If the folder doesn't exist, create it
	if os.IsNotExist(err) {
		err := os.Mkdir(folderPath, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
	}
	out, err := os.Create("./bin-deps/yt-dlp.exe")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	resp, err := http.Get(YtDlpGitgubRelease)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(out, resp.Body)

}
