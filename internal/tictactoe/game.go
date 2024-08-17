package tictactoe

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	Title        = "Tic Tac Toe"
	ScreenWidth  = 600
	ScreenHeight = 600
)

type Position struct {
	x int
	y int
}

type Bound struct {
	w int
	h int
}

type Location struct {
	*Position
	*Bound
}

type Game struct {
	board         *Board
	boardImage    *ebiten.Image
	boardLocation *Location
	cursor        *Position
}

func NewGame() (*Game, error) {
	game := &Game{}

	board, err := NewBoard()
	if err != nil {
		return nil, err
	}

	game.board = board
	game.boardLocation = &Location{
		Position: &Position{},
		Bound:    &Bound{},
	}

	return game, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	g.cursor = &Position{
		x: mx,
		y: my,
	}
	boardCursor := g.getBoardCursorPosition()

	if err := g.board.Update(&CursorInput{
		Position: boardCursor,
		Pressed:  ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft),
	}); err != nil {
		return err
	}

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
	g.boardLocation.Position.x = x
	g.boardLocation.Position.y = y
	g.boardLocation.Bound.w = bw
	g.boardLocation.Bound.h = bh
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)

	in := false
	if g.cursor.x >= x && g.cursor.x <= x+bw && g.cursor.y >= y && g.cursor.y <= y+bh {
		in = true
	}

	msg := fmt.Sprintf("(%d, %d)\n(%d, %d)\n(%d, %d)\n%t", g.cursor.x, g.cursor.y, x, y, bw, bh, in)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) isCursorInBoard() bool {
	return g.cursor.x >= g.boardLocation.x && g.cursor.x <= g.boardLocation.x+g.boardLocation.w && g.cursor.y >= g.boardLocation.y && g.cursor.y <= g.boardLocation.y+g.boardLocation.h
}

func (g *Game) getBoardCursorPosition() Position {
	if !g.isCursorInBoard() {
		return Position{x: -1, y: -1}
	}
	return Position{x: g.cursor.x - g.boardLocation.x, y: g.cursor.y - g.boardLocation.y}
}
