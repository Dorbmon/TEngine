package TEngine

import (
	Layout "github.com/Dorbmon/GoLayout"
)

type Box struct {
	RWidget
	v        bool //VBox or HBox
	children []Widget
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
	z.UpdateChildren(z.children)
}
func (z *Box) Layout(ctx *Layout.Context) *Layout.LayItem {
	tmp := Layout.NewLayItem(ctx)
	z.box = &tmp
	if z.v {
		z.box.SetContain(Layout.LayColumn)
	} else {
		z.box.SetContain(Layout.LayRow)
	}
	for _, item := range z.children {
		layout := item.Layout(ctx)
		z.box.Insert(layout)
	}
	z.RWidget.SetLay(z.box)
	return z.box
}
func (z *Box) PassRenderer(renderer *Renderer) {
	for _, item := range z.children {
		item.PassRenderer(renderer)
	}
}
func (z *Box) Render() error {
	for _, item := range z.children {
		item.Render()
	}
	return nil
}
