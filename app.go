package main

import (
	"context"
	// "fmt"

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
	// background, err := wallpaper.Get()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Current wallpaper:", background)

	// err := wallpaper.SetFromURL("https://cn.bing.com/th?id=OHR.JerseyIsland_EN-US0109101063_UHD.jpg&rf=LaDigue_UHD.jpg&pid=hp&w=3840&h=2160&rs=1&c=4")
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
