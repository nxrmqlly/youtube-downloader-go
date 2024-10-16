package controllers

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) OpenBrowser(url string) {
	runtime.BrowserOpenURL(a.ctx, url)
}
