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

func main() {

	var isProduction = false

	resPath := ""
	if isProduction {
		resPath = "/Applications/gwApp.app/Contents/Resources/"
	} else {
		resPath = "Resources/"
	}

	img1, _ = gw.GetImage(resPath + "k1.png")
	img2, _ = gw.GetImage(resPath + "k2.png")

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

	// img1 = gw.RandomImage(image.Point{cfg.Width, cfg.Height})
	// img2 = gw.RandomImage(image.Point{cfg.Width, cfg.Height})

	// gw.SetDebug()
	gw.Init(cfg, cb)
}

func onFPS(fps int) {
	fmt.Println("FPS:", fps)

}

func onKey(w *gw.Window, key gw.Key, scancode int, action gw.Action, mods gw.ModifierKey) {

	if action == 1 {
		fmt.Print(w.GetSize())
		fmt.Println(" : ", string(key), scancode, action, mods)
	}
}

var xCursor, yCursor float64

func onCursorMove(xPos float64, yPos float64) {
	xCursor, yCursor = xPos, yPos
	// fmt.Println("CURSOR: ", xCursor, yCursor)
}

var r = 0

func onRender() *image.RGBA {

	// xCur, yCur := int(xCursor), int(yCursor)
	var rgba = image.NewRGBA(image.Rect(0, 0, cfg.Width, cfg.Height))
	if gw.Toggle() {
		rgba = img1
		// rgba = gw.RandomImage(image.Point{cfg.Width, cfg.Height})
		// rgba = gw.ClearImage(rgba, color.RGBA{255, 0, 0, 255})
		// rgba, _ = gw.GetImage("black.png")

	} else {
		rgba = img2
		// Hw()
		// var a image.RGBA
		// rgba = gw.RandomImage(image.Point{cfg.Width, cfg.Height})

		// rgba = gw.ClearImage(rgba, color.RGBA{0, 255, 0, 255})
		// rgba, _ = gw.GetImage("black.png")
	}

	return rgba

}
