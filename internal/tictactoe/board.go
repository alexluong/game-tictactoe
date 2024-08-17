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
	Tiles [3][3]*Tile
}

func NewBoard() (*Board, error) {
	board := &Board{}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board.Tiles[i][j] = &Tile{}
		}
	}
	return board, nil
}

// Size returns the board size.
func (b *Board) Size() (int, int) {
	x := boardSize*tileSize + (boardSize+1)*tileMargin
	y := x
	return x, y
}

type CursorInput struct {
	Position Position
	Pressed  bool
}

func (b *Board) Update(cursorInput *CursorInput) error {
	for i, row := range b.Tiles {
		for j, tile := range row {
			hovered := isTileHovered(cursorInput.Position, i, j)
			input := TileInput{
				Hovered: hovered,
				Pressed: false,
			}
			if hovered && cursorInput.Pressed {
				input.Pressed = true
			}
			err := tile.Update(input)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(frameColor)

	// Draw tiles
	for i, row := range b.Tiles {
		for j, tile := range row {
			x := j*tileSize + (j+1)*tileMargin
			y := i*tileSize + (i+1)*tileMargin
			tile.Draw(boardImage, Position{x: x, y: y})
		}
	}
}

func isTileHovered(cursor Position, i, j int) bool {
	x := j*tileSize + (j+1)*tileMargin
	y := i*tileSize + (i+1)*tileMargin
	return cursor.x >= x && cursor.x <= x+tileSize && cursor.y >= y && cursor.y <= y+tileSize
}
