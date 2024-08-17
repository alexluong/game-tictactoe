package tictactoe

import (
	"fmt"
	"game-tictactoe/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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
	winner        string
	over          bool
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

func (g *Game) Update() error {
	if g.over {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			return g.restart()
		}
		return nil
	}

	mx, my := ebiten.CursorPosition()
	g.cursor = &Position{
		x: mx,
		y: my,
	}
	boardCursor := g.getBoardCursorPosition()

	pressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)

	if err := g.board.Update(&CursorInput{
		Position: boardCursor,
		Pressed:  pressed,
	}, g.turn()); err != nil {
		return err
	}

	g.winner = g.checkWinner()
	g.over = g.checkOver()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(g.board.Size())
	}

	screen.Fill(backgroundColor)

	// Board draws to g.boardImage
	g.board.Draw(g.boardImage, g.turn())

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

	// Print result
	if g.over {
		result := ""
		if g.winner != "" {
			result = fmt.Sprintf("%s won.\n\nPress Space to restart.", g.winner)
		} else if g.over {
			result = "Draw. Press Space to restart."
		}

		darkOverlay := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())
		darkOverlay.Fill(color.RGBA{0, 0, 0, 128})
		screen.DrawImage(darkOverlay, nil)
		utils.DrawTextWithShadow(screen, result, 300, 300, 3, color.White, text.AlignCenter, text.AlignCenter)
	}
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

func (g *Game) turn() string {
	filledTiles := 0
	for _, row := range g.board.Tiles {
		for _, tile := range row {
			if tile.Value != "" {
				filledTiles++
			}
		}
	}

	if filledTiles%2 == 0 {
		return "X"
	} else {
		return "O"
	}
}

func (g *Game) checkWinner() string {
	// Check rows
	for _, row := range g.board.Tiles {
		if row[0].Value != "" && row[0].Value == row[1].Value && row[1].Value == row[2].Value {
			return row[0].Value
		}
	}

	// Check columns
	for col := 0; col < 3; col++ {
		if g.board.Tiles[0][col].Value != "" && g.board.Tiles[0][col].Value == g.board.Tiles[1][col].Value && g.board.Tiles[1][col].Value == g.board.Tiles[2][col].Value {
			return g.board.Tiles[0][col].Value
		}
	}

	// Check diagonals
	if g.board.Tiles[0][0].Value != "" && g.board.Tiles[0][0].Value == g.board.Tiles[1][1].Value && g.board.Tiles[1][1].Value == g.board.Tiles[2][2].Value {
		return g.board.Tiles[0][0].Value
	}
	if g.board.Tiles[0][2].Value != "" && g.board.Tiles[0][2].Value == g.board.Tiles[1][1].Value && g.board.Tiles[1][1].Value == g.board.Tiles[2][0].Value {
		return g.board.Tiles[0][2].Value
	}

	// No winner
	return ""
}

func (g *Game) checkOver() bool {
	if g.winner != "" {
		return true
	}

	filledTiles := 0
	for _, row := range g.board.Tiles {
		for _, tile := range row {
			if tile.Value != "" {
				filledTiles++
			}
		}
	}

	return filledTiles == 9
}

func (g *Game) restart() error {
	board, err := NewBoard()
	if err != nil {
		return err
	}
	g.board = board
	g.winner = ""
	g.over = false
	return nil
}
