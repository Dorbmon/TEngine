package TEngine

import Layout "github.com/Dorbmon/GoLayout"

type Rect struct {
	w, h     int32
	renderer *Renderer
	item     Layout.LayItem
}

func NewRect(w, h int32) *Rect {
	return &Rect{w: w, h: h}
}

func (z *Rect) Layout(ctx *Layout.Context) *Layout.LayItem {
	z.item = Layout.NewLayItem(ctx)
	z.item.SetSizeXY(z.w, z.h)
	return &z.item
}
func (z *Rect) PassRenderer(renderer *Renderer) {
	z.renderer = renderer
}
func (z *Rect) Render() error {
	z.renderer.SetColor(0xAD99C0, 0)
	z.renderer.FillRect(0, 0, z.w, z.h)
	return nil
}
