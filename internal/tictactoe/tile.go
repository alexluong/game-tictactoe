package tictactoe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tileSize   = 100
	tileMargin = 4
)

var (
	tileColor = color.NRGBA{0xee, 0xe4, 0xda, 0x59}
	tileImage = ebiten.NewImage(tileSize, tileSize)
)

func init() {
	tileImage.Fill(color.White)
}
