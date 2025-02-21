package draw

import (
	"chess-engines/internal/chess"

	"github.com/veandco/go-sdl2/sdl"
)

var White sdl.Color = sdl.Color{R: 255, G: 255, B: 255, A: 255}
var Black sdl.Color = sdl.Color{R: 0, G: 0, B: 0, A: 255}
var Red sdl.Color = sdl.Color{R: 255, G: 0, B: 0, A: 255}
var Green sdl.Color = sdl.Color{R: 0, G: 255, B: 0, A: 255}
var Blue sdl.Color = sdl.Color{R: 0, G: 0, B: 255, A: 255}
var Purple sdl.Color = sdl.Color{R: 255, G: 0, B: 255, A: 255}

const Width = 600
const Height = 600 
const squareWidth = Width / chess.Cols

const WindowTitle = "Chess Engines"

type DrawWindow struct {
	sdlWindow   *sdl.Window
	renderer    *sdl.Renderer
	RealWidth   int32
	RealHeight  int32
	RealSquareWidth int32
}

// Init creates a new DrawWindow and inits SDL2
func Init() *DrawWindow {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(WindowTitle,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		Width, Height, sdl.WINDOW_SHOWN|sdl.WINDOW_ALLOW_HIGHDPI)

	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	realWidth, realHeight, err := renderer.GetOutputSize()

	if err != nil {
		panic(err)
	}

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	realSquareWidth := realWidth / chess.Cols


	return &DrawWindow{window, renderer, realWidth,
		realHeight, realSquareWidth}
}

// QUIT destroyes DrawWindow and quits SDL
func Quit(window *DrawWindow) {
	sdl.Quit()
	window.sdlWindow.Destroy()
	window.renderer.Destroy()
}

// Update updates the SDL window with new drawings
func Update(window *DrawWindow) {
	window.renderer.Present()
}

// Clear fills the screen with a single color
func Clear(color sdl.Color, window *DrawWindow) {
	window.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	window.renderer.Clear()
}

// Rect draws a filled rect on the given DrawWindow
func DrawRect(rect sdl.Rect, color sdl.Color, window *DrawWindow) {
	window.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	window.renderer.FillRect(&rect)
}
