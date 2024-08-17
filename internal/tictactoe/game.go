package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Title        = "Tic Tac Toe"
	ScreenWidth  = 600
	ScreenHeight = 600
)

type Game struct {
	board      *Board
	boardImage *ebiten.Image
}

func NewGame() (*Game, error) {
	game := &Game{}

	board, err := NewBoard()
	if err != nil {
		return nil, err
	}

	game.board = board

	return game, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(g.board.Size())
	}

	screen.Fill(backgroundColor)

	// Board draws to g.boardImage
	g.board.Draw(g.boardImage)

	// Draw g.boardImage to screen
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}
