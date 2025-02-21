package draw

import (
	"chess-engines/internal/chess"
	"github.com/veandco/go-sdl2/sdl"
)

var WhiteTileColor = sdl.Color{R: 136, G: 174, B: 163}
var BlackTileColor = sdl.Color{R: 108, G: 148, B: 137}


func DrawBoard(board *chess.Board, assets *Assets, window *DrawWindow) {
	drawBackground(window)
	var row int32
	var col int32
	for row = 0; row < chess.Rows; row++ {
		for col = 0; col < chess.Cols; col++ {
			piece := board.BitBoard[row*chess.Cols+col]
			if piece != 0 {
				image := assets.Images[piece]
				DrawImage(image, col*window.RealSquareWidth,
					row*window.RealSquareWidth, window)
			}
		}
	}

}

// drawBackground draws the light and dark tiles to window
func drawBackground(window *DrawWindow) {
	var row int32
	var col int32
	for row = 0; row < chess.Rows; row++ {
		for col = 0; col < chess.Cols; col++ {
			color := BlackTileColor
			if (col+row)%2 == 0 {
				color = WhiteTileColor
			}
			x := col * window.RealSquareWidth
			y := row * window.RealSquareWidth
			DrawRect(sdl.Rect{X: x, Y: y, W: window.RealSquareWidth, H:
				window.RealSquareWidth}, color, window)
		}
	}
}
