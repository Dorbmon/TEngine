package TEngine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type App struct {
	Title     string
	render    *sdl.Renderer
	wins      map[uint32]*Window
	FrameRate uint32
}

func NewApp() (app *App, err error) {
	app = &App{wins: make(map[uint32]*Window), FrameRate: 60}
	return
}
func (z *App) AddWindow(win *Window) error {
	id, err := win.win.GetID()
	if err != nil {
		return err
	}
	z.wins[id] = win
	return err
}
func (z *App) Run() {
	sdl.Main(func() {
		running := true
		for running {
			sdl.Do(func() {
				for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
					switch event.(type) {
					case *sdl.QuitEvent:
						running = false
					case *sdl.MouseButtonEvent:
						e := event.(*sdl.MouseButtonEvent)
						z.wins[e.WindowID].onMouseButton(e)
					}
				}
			})
			for _, win := range z.wins {
				win.Render()
			}
			sdl.Delay(1000 / z.FrameRate)
		}
	})
}
