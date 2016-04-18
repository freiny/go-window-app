package main

import (
	"fmt"
	"image"
	_ "image/png"
	"math/rand"
	"time"

	gw "github.com/freiny/go-window"
)

var cfg = gw.Config{}
var cb = gw.Callbacks{}

var img1, img2 *image.RGBA

// var installPath = "/Applications/GoWindow.app/Contents/Resources/"

var k1, _ = gw.GetImage("k1.png")
var k2, _ = gw.GetImage("k2.png")

// var k2, _ = gw.GetImagePart("k2.png", image.Point{64, 64})

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// cfg.Width = 1280
	// cfg.Height = 720
	cfg.Width = 8 * 64
	cfg.Height = 8 * 64
	cfg.X = 0
	cfg.Y = 0

	cb.Render = onRender
	cb.Key = onKey
	cb.CursorMove = onCursorMove
	cb.FPS = onFPS

	img1 = gw.RandomImage(image.Point{cfg.Width, cfg.Height})
	img2 = gw.RandomImage(image.Point{cfg.Width, cfg.Height})

	gw.SetDebug()
	gw.Init(cfg, cb)
}

func onFPS(fps int) {
	fmt.Println("FPS:", fps)

}

// func onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
// 	if action == 1 {
// 		fmt.Print(string(key), " DOWN")
//
// 	}
//
// 	if action == 0 {
// 		fmt.Print(string(key), " UP")
// 	}
// }
func onKey(w *gw.Window) {
	fmt.Println("Key Pressed...")
	fmt.Println(w.GetSize())
}

var xCursor, yCursor float64

func onCursorMove(xPos float64, yPos float64) {
	xCursor, yCursor = xPos, yPos
	// fmt.Println("CURSOR: ", xCursor, yCursor)
}

// var trgba =
var r = 0

func onRender() *image.RGBA {

	// xCur, yCur := int(xCursor), int(yCursor)
	var rgba = image.NewRGBA(image.Rect(0, 0, cfg.Width, cfg.Height))
	if gw.Toggle() {
		rgba = k1
		// rgba = gw.RandomImage(image.Point{cfg.Width, cfg.Height})
		// rgba = gw.ClearImage(rgba, color.RGBA{255, 0, 0, 255})
		// rgba, _ = gw.GetImage("black.png")

	} else {
		rgba = k2
		// Hw()
		// var a image.RGBA
		// rgba = gw.RandomImage(image.Point{cfg.Width, cfg.Height})

		// rgba = gw.ClearImage(rgba, color.RGBA{0, 255, 0, 255})
		// rgba, _ = gw.GetImage("black.png")
	}

	return rgba

}
