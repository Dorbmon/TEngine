package TEngine

import (
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
	ret.root.SetContain(Layout.LayColumn)
	ret.root.SetSizeXY(Data.W, Data.H)
	return ret, nil
}
func (z *Window) Render() error {
	if z.body == nil {
		return nil
	}
	z.layout.Reset()
	//z.layout = Layout.NewLayout()
	z.layout.ReserveItemsCapacity(1000)
	z.root = Layout.NewLayItem(&z.layout)
	z.root.SetContain(Layout.LayColumn)
	z.root.Insert(z.body.Layout(&z.layout))
	z.root.SetSizeXY(z.win.GetSize())
	z.layout.Calculate()
	z.body.PassRenderer(z.rRender)
	sdl.Do(func() {
		z.renderer.SetDrawColor(255, 0, 0, 255)
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

func (z *Window) onMouseButton(event *sdl.MouseButtonEvent) {
	var button int
	switch event.Button {
	case sdl.BUTTON_LEFT:
		button = ButtonLeft
	case sdl.BUTTON_MIDDLE:
		button = ButtonMid
	case sdl.BUTTON_RIGHT:
		button = ButtonRight
	case sdl.BUTTON_X1:
		button = ButtonX1
	case sdl.BUTTON_X2:
		button = ButtonX2
	}
	var state int
	switch event.State {
	case sdl.RELEASED:
		state = MouseRelease
	case sdl.PRESSED:
		state = MousePress
	}
	z.body.OnHit(HitPositon{X: int(event.X), Y: int(event.Y)}, button, state)
}
