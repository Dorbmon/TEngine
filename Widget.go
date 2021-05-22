package TEngine

import (
	Layout "github.com/Dorbmon/GoLayout"
)

type Widget interface {
	Layout(ctx *Layout.Context) *Layout.LayItem
	PassRenderer(renderer *Renderer)
	Render() error
	OnHit(Position HitPositon, HitButton int, State int)
}
type RWidget struct {
	layout   *Layout.LayItem
	onClick  func()
	children []Widget
}

func (z *RWidget) SetLay(layout *Layout.LayItem) {
	z.layout = layout
}
func (z *RWidget) OnHit(Position HitPositon, HitButton int, State int) {
	rect := z.layout.GetRect()
	if HitButton == ButtonLeft && State == MousePress && rect.X1 <= Position.X && Position.X <= rect.X2 && Position.Y >= rect.Y1 && Position.Y <= rect.Y2 {
		if z.onClick != nil {
			z.onClick()
		}
	}
	if z.children != nil {
		for _, kid := range z.children {
			kid.OnHit(Position, HitButton, State)
		}
	}
}

// UpdateChildren the type of kid should be Widget or []Widget
func (z *RWidget) UpdateChildren(kid interface{}) {
	switch t := kid.(type) {
	case Widget:
		z.children = []Widget{t}
	case []Widget:
		z.children = t
	default:
		panic("error type of kid offered to UpdateChildren")
	}
}
func (z *RWidget) OnClick(f func()) {
	z.onClick = f
}
