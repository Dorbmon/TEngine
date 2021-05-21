package main

import (
	"TEngine"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	app, _ := TEngine.NewApp()
	window, err := TEngine.NewWindow("Rx-TEngine", sdl.Rect{X: 0, Y: 0, W: 500, H: 200})
	if err != nil {
		panic(err)
	}
	vbox := TEngine.NewVBox()
	window.SetBody(vbox)
	block := TEngine.NewRect(500, 100, TEngine.NewColorFromHex(0xCF143F, 255))
	vbox.Append(block)
	block2 := TEngine.NewRect(500, 100, TEngine.NewColorFromHex(0xFFFFF, 255))
	vbox.Append(block2)
	app.AddWindow(window)
	app.Run()
}
