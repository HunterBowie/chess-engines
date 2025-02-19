package draw

import (
	"github.com/veandco/go-sdl2/sdl"
)

var White sdl.Color = sdl.Color{R: 255, G: 255, B: 255, A: 255}
var Black sdl.Color = sdl.Color{R: 0, G: 0, B: 0, A: 255}
var Red sdl.Color = sdl.Color{R: 255, G: 0, B: 0, A: 255}
var Green sdl.Color = sdl.Color{R: 0, G: 255, B: 0, A: 255}
var Blue sdl.Color = sdl.Color{R: 0, G: 0, B: 255, A: 255}
var Purple sdl.Color = sdl.Color{R: 255, G: 0, B: 255, A: 255}

type DrawWindow struct {
	sdlWindow *sdl.Window
	surface   *sdl.Surface
}

// Init creates a new DrawWindow and inits SDL2
func Init(title string, width int32, height int32) *DrawWindow {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	return &DrawWindow{window, surface}
}

// QUIT destroyes DrawWindow and quits SDL
func Quit(window *DrawWindow) {
	sdl.Quit()
	window.sdlWindow.Destroy()
}

// Update updates the SDL window with new drawings
func Update(window *DrawWindow) {
	window.sdlWindow.UpdateSurface()
}

// Clear fills the screen with a single color
func Clear(color sdl.Color, window *DrawWindow) {
	pixelColor := sdl.MapRGBA(window.surface.Format, color.R, color.G, color.B, color.A)
	window.surface.FillRect(nil, pixelColor)
}

// Rect draws a filled rect on the given DrawWindow
func Rect(rect sdl.Rect, color sdl.Color, window *DrawWindow) {
	pixelColor := sdl.MapRGBA(window.surface.Format, color.R, color.G, color.B, color.A)
	window.surface.FillRect(&rect, pixelColor)
}
