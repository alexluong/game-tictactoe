package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	prestige "github.com/mbrc12/prestige"
)

const (
	Title        = "Tic Tac Toe"
	ScreenWidth  = 600
	ScreenHeight = 600
)

type State struct {
	Manager *prestige.SceneManager[State]
}

func NewGame() ebiten.Game {
	state := &State{}
	manager := prestige.NewSceneManager(&TitleScene{}, state)
	state.Manager = manager
	return manager
}
