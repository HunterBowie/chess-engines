package draw

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const ImagesPath = "assets/images/"

var pieceNames = []string{"pawn", "bishop",
	"knight", "rook", "queen", "king"}

type Assets struct {
	Images map[int8]*sdl.Texture
	// Sounds *[]
}

// LoadImage returns an image text given the image name
func LoadImage(imageName string, window *DrawWindow) *sdl.Texture {
	texture, err := img.LoadTexture(window.renderer,
		ImagesPath+imageName)

	if err != nil {
		panic(err)
	}
	return texture
}

// LoadAssets loads all images an sounds into a Assets struct
func LoadAssets(window *DrawWindow) *Assets {
	assets := Assets{make(map[int8]*sdl.Texture)}

	for nameIndex, name := range pieceNames {
		for colorIndex, color := range []string{"white", "black"} {
			image := LoadImage(color+"-"+name+".png", window)
			pieceColor := int8(8 * colorIndex)
			pieceName := int8(1 + nameIndex)
			assets.Images[pieceColor|pieceName] = image
		}
	}
	return &assets
}

// Image draws an image on to the given window
func DrawImage(image *sdl.Texture, x, y int32, window *DrawWindow) {
	dest := sdl.Rect{X: x, Y: y,
		W: window.RealSquareWidth, H: window.RealSquareWidth}
	window.renderer.Copy(image, nil, &dest)
}
