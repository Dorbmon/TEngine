package TEngine

import Layout "github.com/Dorbmon/GoLayout"

type Text struct {
	RWidget
	w, h     int32
	word     string
	renderer *Renderer
	item     Layout.LayItem
}

func NewText(word string) *Text {
	ret := &Text{
		word: word,
	}
	return ret
}
func (z *Text) Layout(ctx *Layout.Context) *Layout.LayItem {
	z.item = Layout.NewLayItem(ctx)
	z.item.SetSizeXY(z.w, z.h)
	z.RWidget.SetLay(&z.item)
	return &z.item
}
func (z *Text) PassRenderer(renderer *Renderer) {
	z.renderer = renderer
	rect := z.item.GetRect()
	z.renderer.UpdateOffset(int32(rect.X1), int32(rect.Y1))
}
func (z *Text) Render() error {
	z.renderer.SetDrawColor(z.color)
	z.renderer.FillRect(0, 0, z.w, z.h)
	return nil
}
