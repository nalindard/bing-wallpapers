package main

import (
	"context"

	"github.com/reujab/wallpaper"
)

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

// Set wallpaper,
func (a *App) SetWallpaper(src string) {
	if len(src) < 20 {
		panic("err")
	}
	err := wallpaper.SetFromURL(src)

	if err != nil {
		panic(err)
	}

	// err = wallpaper.SetMode(wallpaper.Crop)
	// if err != nil {
	// 	panic(err)
	// }
}
