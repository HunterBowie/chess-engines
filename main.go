package main

import (
	"chess-engines/internal/chess"
	"chess-engines/internal/draw"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	window := draw.Init()
	defer draw.Quit(window)

	assets := draw.LoadAssets(window)

	board := chess.CreateBoard()

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
			case *sdl.MouseButtonEvent:
				if event.Type == sdl.MOUSEBUTTONDOWN {
					// x and y are not real
					x := event.X
					y := event.Y
					col := x / (draw.Width / chess.Cols)
					row := y / (draw.Height / chess.Rows)
					chess.SetBoardPos(row, col, 1, board)
				}

			}

		}
		draw.DrawBoard(board, assets, window)
		draw.Update(window)
		draw.Clear(draw.White, window)

	}
}
