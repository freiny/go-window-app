package main

import (
	"fmt"
	"image"
	_ "image/png"
	"math/rand"
	"time"

	glgw "github.com/freiny/go-window"
)

var dbg = glgw.Debugger{}.Init()

var g = struct {
	wc glgw.WinConfig
	// cb   gw.Callbacks
	cur  glgw.Cursor
	img1 *image.RGBA
	img2 *image.RGBA
}{
	img1: gw.GetImage(gw.GetResourcePath() + "k1.png"),
	img2: gw.GetImage(gw.GetResourcePath() + "k2.png"),
}

var gw = glgw.Config(0, 0, 512, 512)

func init() {
	dbg.Log("main.go", "main()")
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	gw.RegisterCallback(onRender)
	gw.RegisterCallback(onCursorMove)
	gw.RegisterCallback(onKey)
	gw.RegisterCallback(onFPS)

	gw.Start()
}

func onFPS(fps int) {
	if dbg.Profiling() {
		dbg.Log("main.go", "onFPS()", "FPS:", fps)
	}
}

func onKey(w *glgw.Window, key glgw.Key, scancode int, action glgw.Action, mods glgw.ModifierKey) {
	dbg.Log("main.go", "onKey()")

	gw.SetPos(150, 150) // set new window position on key press

	if action == 1 {
		fmt.Print(w.GetSize())
		fmt.Println(" : ", string(key), scancode, action, mods)
	}
}

func onCursorMove(c glgw.Cursor) {
	dbg.Log("main.go", "onCursorMove()")
	g.cur = c
}

func onRender() *image.RGBA {
	if dbg.Profiling() {
		dbg.Log("main.go", "onRender()")
	}

	// xCur, yCur := int(xCursor), int(yCursor)
	var rgba = image.NewRGBA(image.Rect(0, 0, g.wc.W, g.wc.Y))
	if glgw.SetFrameToggle() {
		rgba = g.img1
	} else {
		rgba = g.img2
	}

	return rgba

}
