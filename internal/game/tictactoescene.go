package game

import (
	"game-tictactoe/internal/tictactoe"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	prestige "github.com/mbrc12/prestige"
)

type TicTacToeScene struct {
	game *tictactoe.Game
}

func (scene *TicTacToeScene) Update(state *State) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		return state.Manager.Transition(&TitleScene{}, prestige.NewFadeTransition(1))
	}
	return scene.game.Update()
}

func (scene *TicTacToeScene) Draw(state *State, screen *ebiten.Image) {
	scene.game.Draw(screen)
}

func (scene *TicTacToeScene) EnterStart(state *State) {
	game, err := tictactoe.NewGame()
	if err != nil {
		log.Panic(err)
	}
	scene.game = game
}

func (scene *TicTacToeScene) EnterEnd(state *State) {
}

func (scene *TicTacToeScene) ExitStart(state *State) {
}

func (scene *TicTacToeScene) ExitEnd(state *State) {
}

func (scene *TicTacToeScene) Layout(state *State, w, h int) (int, int) {
	return w, h
}
