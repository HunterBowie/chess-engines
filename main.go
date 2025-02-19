package main

import (
	"app/internal/draw"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	window := draw.Init("Chess", 600, 600)

	defer draw.Quit(window)

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				keyEvent, _ := event.(*sdl.KeyboardEvent)
				key := keyEvent.Keysym.Sym
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
