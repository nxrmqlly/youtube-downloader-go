package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wader/goutubedl"
)

const YtDlpGitgubRelease = "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp.exe"

// If yt-dlp isn't installed yet, download it locally.
func getYtDlp() {
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

// Get yt-dlp if not exists, returns cwd and path of yt-dlp
func YtDlSetup() (string, string) {
	// check if ./bin-deps/yt-dlp.exe exists
	if _, err := os.Stat("./bin-deps/yt-dlp.exe"); os.IsNotExist(err) {
		getYtDlp()
	}

	var cwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	goutubedl.Path = filepath.Join(cwd, "bin-deps", "yt-dlp.exe")

	return cwd, goutubedl.Path
}
