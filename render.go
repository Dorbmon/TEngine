package TEngine

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

type Renderer struct {
	x      int32
	y      int32
	w      int32
	h      int32
	render *sdl.Renderer
}

var renderers = &sync.Pool{New: func() interface{} {
	return new(Renderer)
}}

func NewRenderer(source *sdl.Renderer, x, y, w, h int32) (renderer *Renderer) {
	renderer = renderers.Get().(*Renderer)
	renderer.Resize(x, y, w, h)
	renderer.render = source
	return
}
func (z *Renderer) DrawLine(x1, y1, x2, y2 int32) error {
	return z.render.DrawLine(x1+z.x, y1+z.y, x2+z.x, y2+z.y)
}
func (z *Renderer) UpdateOffset(x, y int32) {
	z.x, z.y = x, y
}
func (z *Renderer) Resize(x, y int32, w, h int32) {
	z.x = x
	z.y = y
	z.w = w
	z.h = h
}
func (z *Renderer) ToLocal(dx, dy, w, h int32) *Renderer {
	renderer := renderers.Get().(*Renderer)
	renderer.x = z.x + dx
	renderer.y = z.y + dy
	renderer.w = w
	renderer.h = h
	renderer.render = z.render
	return renderer
}
func (z *Renderer) DrawRect(x, y, w, h int32) error {
	return z.render.DrawRect(&sdl.Rect{X: x + z.x, Y: y + z.y, W: w, H: h})
}
func (z *Renderer) FillRect(x, y, w, h int32) {
	z.render.FillRect(&sdl.Rect{X: x + z.x, Y: y + z.y, W: w, H: h})
}
func (z *Renderer) SetDrawColor(Color RColor) {
	r, g, b, a := Color.ToRGBA()
	z.render.SetDrawColor(r, g, b, a)
}

//Release release the renderer
func (z *Renderer) Release() {
	z.render.Destroy()
	renderers.Put(z)
}
