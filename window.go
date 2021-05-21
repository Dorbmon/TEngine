package TEngine

import (
	"fmt"
	"time"

	Layout "github.com/Dorbmon/GoLayout"
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	win       *sdl.Window
	renderer  *sdl.Renderer
	rRender   *Renderer
	body      Widget
	layout    Layout.Context
	root      Layout.LayItem
	frameRate uint32
}

func (z *Window) SetTitle(Title string) {
	z.win.SetTitle(Title)
}
func (z *Window) SetSize(w, h int32) {
	z.win.SetSize(w, h)
	z.root.SetSizeXY(w, h)
}
func NewWindow(Title string, Data sdl.Rect) (*Window, error) {
	win, err := sdl.CreateWindow(Title, Data.X, Data.Y, Data.W, Data.H, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}
	renderer, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return nil, err
	}
	rRender := NewRenderer(renderer, 0, 0, 0, 0)
	ret := &Window{win: win, renderer: renderer, rRender: rRender, layout: Layout.NewLayout(), frameRate: 60}
	ret.layout.ReserveItemsCapacity(1000)
	ret.root = Layout.NewLayItem(&ret.layout)
	ret.root.SetContain(Layout.LayRow)
	ret.root.SetSizeXY(Data.W, Data.H)
	return ret, nil
}
func (z *Window) Render() error {
	if z.body == nil {
		return nil
	}
	z.layout.Destroy()
	z.layout = Layout.NewLayout()
	z.layout.ReserveItemsCapacity(1000)
	z.root = Layout.NewLayItem(&z.layout)
	z.root.SetContain(Layout.LayRow)
	z.root.Insert(z.body.Layout(&z.layout))
	z.root.SetSizeXY(z.win.GetSize())
	z.layout.Calculate()
	z.body.PassRenderer(z.rRender)
	sdl.Do(func() {
		z.renderer.Clear()
	})
	z.body.Render()
	sdl.Do(func() {
		z.renderer.Present()
	})
	return nil
}
func (z *Window) SetBody(body Widget) error {
	w, h := z.win.GetSize()
	z.rRender.Resize(0, 0, w, h)
	z.body = body
	return nil
}
func (z *Window) Run() error {
	sdl.Main(func() {
		running := true
		for running {
			sdl.Do(func() {
				for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
					switch event.(type) {
					case *sdl.QuitEvent:
						running = false
					case *sdl.MouseButtonEvent:

					}
				}
			})
			t := time.Now()
			if err := z.Render(); err != nil {
				return
			}
			t1 := time.Now()
			fmt.Println("render cost:", t1.Sub(t).Milliseconds())
			sdl.Delay(1000 / z.frameRate)
		}
	})

	return nil
}
