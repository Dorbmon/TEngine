package TEngine

import (
	Layout "github.com/Dorbmon/GoLayout"
)

type Box struct {
	v        bool //VBox or HBox
	children []Widget
	layouts  []*Layout.LayItem
	box      *Layout.LayItem
}

func NewVBox() *Box {
	return &Box{
		v:        true,
		children: make([]Widget, 0),
	}
}
func NewHBox() *Box {
	return &Box{
		v:        false,
		children: make([]Widget, 0),
	}
}
func (z *Box) Append(Item Widget) {
	z.children = append(z.children, Item)
}
func (z *Box) Layout(ctx *Layout.Context) *Layout.LayItem {
	tmp := Layout.NewLayItem(ctx)
	z.box = &tmp
	if z.v {
		z.box.SetContain(Layout.LayColumn)
	} else {
		z.box.SetContain(Layout.LayRow)
	}
	z.layouts = make([]*Layout.LayItem, len(z.children))
	for index, item := range z.children {
		layout := item.Layout(ctx)
		z.box.Insert(layout)
		z.layouts[index] = layout
	}
	return z.box
}
func (z *Box) PassRenderer(renderer *Renderer) {
	all := z.box.GetRect()
	for index, item := range z.children {
		layout := z.layouts[index]
		rect := layout.GetRect()
		item.PassRenderer(renderer.ToLocal(int32(rect.X1-all.X1), int32(rect.Y1-all.Y1), int32(rect.X2-rect.X1), int32(rect.Y2-rect.Y1)))
	}
}
func (z *Box) Render() error {
	for _, item := range z.children {
		item.Render()
	}
	return nil
}
