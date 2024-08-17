package tictactoe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	boardSize       = 3
	backgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	frameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)

type Board struct {
}

func NewBoard() (*Board, error) {
	board := &Board{}
	return board, nil
}

// Size returns the board size.
func (b *Board) Size() (int, int) {
	x := boardSize*tileSize + (boardSize+1)*tileMargin
	y := x
	return x, y
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(frameColor)
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			op := &ebiten.DrawImageOptions{}
			x := j*tileSize + (j+1)*tileMargin
			y := i*tileSize + (i+1)*tileMargin
			op.GeoM.Translate(float64(x), float64(y))
			op.ColorScale.ScaleWithColor(tileColor)
			boardImage.DrawImage(tileImage, op)
		}
	}
}
