package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"chess-engines/internal/draw"
	"chess-engines/internal/engines/minimax"
)

func main() {

	minimax.MyFunc()

	window := draw.Init("Chess", 600, 600)

	defer draw.Quit(window)

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				key := event.Keysym.Sym
				if key == sdl.K_UP {
					rect.Y -= 10
				}
				if key == sdl.K_DOWN {
					rect.Y += 10
				}
				
			}
		}

		draw.Rect(rect, draw.Purple, window)

		draw.Update(window)
		draw.Clear(draw.White, window)

	}
}
