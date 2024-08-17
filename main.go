package main

import (
	"log"

	"game-tictactoe/internal/tictactoe"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := tictactoe.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(tictactoe.ScreenWidth, tictactoe.ScreenHeight)
	ebiten.SetWindowTitle(tictactoe.Title)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
