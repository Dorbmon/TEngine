package main

import (
	"TEngine"
	"fmt"

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
	block.OnClick(func() {
		fmt.Println("Got Click On block")
	})
	vbox.Append(block)
	block2 := TEngine.NewRect(500, 100, TEngine.NewColorFromHex(0xFFFFF, 255))
	block2.OnClick(func() {
		fmt.Println("Got Click On block2")
	})
	vbox.Append(block2)
	app.AddWindow(window)
	app.Run()
}
