package main

import (
	"embed"
	"fmt"
	"os"

	"youtube-downloader-go/backend/controllers"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	fmt.Println(os.Getwd())

	// Create an instance of the app structure
	app := controllers.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "youtube-downloader-go",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
