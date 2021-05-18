package TEngine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type App struct {
	Title  string
	render *sdl.Renderer
}

func NewApp() (app *App, err error) {
	app = &App{
		Title: "T-Engine",
	}
	return
}

func (z *App) Run() {

}
