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
	// createImagesFolder()
	// err := downloadImage("save.img", "https://gophercoding.com/img/logo-original.png")
	// if err != nil {
	// 	panic(err)
	// }

	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	// err := DownloadFile("https://cn.bing.com/th?id=OHR.ValGardenaItaly_EN-US8887980856_UHD.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	// err = wallpaper.SetFromFile("wailpaper.jpeg")
	// if err != nil {
	// 	panic(err)
	// }
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

// // Create images folder,
// func createImagesFolder() {
// 	// Get the user's document folder path
// 	docPath, err := os.UserHomeDir()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// docPath = filepath.Join(docPath, "Documents")
// 	// docPath = filepath.Join(docPath, "Pictures")

// 	// Create a folder named 'images' if not exist
// 	imgPath := filepath.Join(docPath, "Wailspapers")
// 	if _, err := os.Stat(imgPath); os.IsNotExist(err) {
// 		err := os.Mkdir(imgPath, 0755)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Println("Folder 'images' created successfully")
// 	} else {
// 		fmt.Println("Folder 'images' already exists")
// 	}
// }

// // Download the image,
// func downloadImage(imageName string, url string) error {

// 	docPath, err := os.UserHomeDir()
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	imgPath := filepath.Join(docPath, "Wailspapers")

// 	// Get the data
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Create the file
// 	out, err := os.Create(imgPath + imageName)
// 	if err != nil {
// 		return err
// 	}
// 	defer out.Close()

// 	// Write the body to file
// 	_, err = io.Copy(out, resp.Body)
// 	return err
// }

// func DownloadFile(url string) error {
// 	// Get the data
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Create the file
// 	// filepath := path.Base(resp.Request.URL.String())
// 	filepath := "wailpaper.jpeg"
// 	out, err := os.Create(filepath)
// 	if err != nil {
// 		return err
// 	}
// 	defer out.Close()

// 	err = os.Remove(filepath)
// 	if err != nil {
// 		return err
// 	}

// 	// Write the body to file
// 	_, err = io.Copy(out, resp.Body)
// 	return err
// }
