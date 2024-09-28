package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/wader/goutubedl"
)

var youtubeRegex, _ = regexp.Compile(`^((?:https?:)?\/\/)?((?:www|m)\.)?((?:youtube(?:-nocookie)?\.com|youtu.be))(\/(?:[\w\-]+\?v=|embed\/|live\/|v\/)?)([\w\-]+)(\S+)?$`)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Check if the url is valid youtube video URL
func isValidYouTubeURL(url string) bool {
	return youtubeRegex.MatchString(url)
}
func makeOutputDir() {
	_, err := os.Stat("./output")
	if os.IsNotExist(err) {
		err := os.Mkdir("./output", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Download the video by the url specified
func (a *App) DownloadVideo(url string) (string, error) {
	fmt.Println("url: ", url)
	//* check if the url is valid
	if !isValidYouTubeURL(url) {
		return "Invalid URL", fmt.Errorf("invalid URL")
	}
	// check if ./bin-deps/yt-dlp.exe exists
	if _, err := os.Stat("./bin-deps/yt-dlp.exe"); os.IsNotExist(err) {
		GetYtDlp()
	}

	var cwd, err = os.Getwd()
	goutubedl.Path = filepath.Join(cwd, "bin-deps", "yt-dlp.exe")

	result, err := goutubedl.New(context.Background(), url, goutubedl.Options{})

	if err != nil {
		return "Error getting video", err
	}

	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		return "Error downloading", err
	}

	defer downloadResult.Close()

	// todo: Type of video/other options
	// todo: audio only

	//save the video
	makeOutputDir()
	var videoTitle = result.Info.Title
	pathToCreate := fmt.Sprintf("%s/output/%s.mp4", cwd, videoTitle)
	fmt.Println(pathToCreate)
	f, err := os.Create(pathToCreate)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(f, downloadResult)

	return fmt.Sprintf("%s/output/%s.mp4 downloaded", cwd, videoTitle), nil
}
