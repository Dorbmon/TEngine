package TEngine

import (
	Layout "github.com/Dorbmon/GoLayout"
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	win      *sdl.Window
	renderer *sdl.Renderer
	rRender  *Renderer
	body     Widget
	layout   Layout.Context
	root     Layout.LayItem
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
	ret := &Window{win: win, renderer: renderer, rRender: NewRenderer(renderer, 0, 0, 0, 0), layout: Layout.NewLayout()}
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
	z.body.Render()
	z.renderer.Present()
	return nil
}
func (z *Window) SetBody(body Widget) error {
	w, h := z.win.GetSize()
	z.rRender.Resize(0, 0, w, h)
	z.body = body
	return z.Render()
}
func (z *Window) Run() error {
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
		if err := z.Render(); err != nil {
			return err
		}
		sdl.Delay(16)
	}
	return nil
}
