package TEngine

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

func TestALl(t *testing.T) {
	sdl.Init(sdl.INIT_VIDEO)
	window, err := NewWindow("Rx-TEngine", sdl.Rect{X: 0, Y: 0, W: 100, H: 100})
	if err != nil {
		panic(err)
	}
	vbox := NewVBox()
	window.SetBody(vbox)
	block := NewRect(100, 100)
	vbox.Append(block)
	window.Run()
}
