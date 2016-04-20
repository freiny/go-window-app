package main

import (
	"fmt"
	"image"
	_ "image/png"
	"math/rand"
	"time"

	gw "github.com/freiny/go-window"
)

// Global variables relative to user application
type Global struct {
	wc   gw.WinConfig
	cb   gw.Callbacks
	cur  gw.Cursor
	img1 *image.RGBA
	img2 *image.RGBA
}

var g = Global{}

func main() {

	var isProduction = false

	resPath := ""
	if isProduction {
		resPath = "/Applications/gwApp.app/Contents/Resources/"
	} else {
		resPath = "Resources/"
	}

	g.img1, _ = gw.GetImage(resPath + "k1.png")
	g.img2, _ = gw.GetImage(resPath + "k2.png")

	rand.Seed(time.Now().UTC().UnixNano())

	g.wc = gw.WinConfig{
		W: 8 * 64,
		H: 8 * 64,
		X: 0,
		Y: 0,
	}
	g.cb = gw.Callbacks{
		// Render: onRender,
		// CursorMove: onCursorMove,
		// Key:        onKey,
		FPS: onFPS,
	}

	gw.SetDebug()
	gw.Init(g.wc, g.cb)
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

func onCursorMove(c gw.Cursor) {
	g.cur = c
}

func onRender() *image.RGBA {

	// xCur, yCur := int(xCursor), int(yCursor)
	var rgba = image.NewRGBA(image.Rect(0, 0, g.wc.W, g.wc.Y))
	if gw.SetFrameToggle() {
		rgba = g.img1
	} else {
		rgba = g.img2
	}

	return rgba

}
