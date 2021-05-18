package TEngine

import Layout "github.com/Dorbmon/GoLayout"

type Widget interface {
	Layout(ctx *Layout.Context) *Layout.LayItem
	PassRenderer(renderer *Renderer)
	Render() error
}
