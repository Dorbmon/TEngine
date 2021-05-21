package TEngine

import (
	Layout "github.com/Dorbmon/GoLayout"
)

type Rect struct {
	w, h     int32
	renderer *Renderer
	item     Layout.LayItem
	color    RColor
}

func NewRect(w, h int32, color RColor) *Rect {
	return &Rect{w: w, h: h, color: color}
}

func (z *Rect) Layout(ctx *Layout.Context) *Layout.LayItem {
	z.item = Layout.NewLayItem(ctx)
	z.item.SetSizeXY(z.w, z.h)
	return &z.item
}
func (z *Rect) PassRenderer(renderer *Renderer) {
	z.renderer = renderer
	rect := z.item.GetRect()
	z.renderer.UpdateOffset(int32(rect.X1), int32(rect.Y1))
}
func (z *Rect) Render() error {
	z.renderer.SetDrawColor(z.color)
	z.renderer.FillRect(0, 0, z.w, z.h)
	return nil
}
