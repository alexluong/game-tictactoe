package game

import (
	"game-tictactoe/internal/utils"
	"image/color"
	rand "math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	prestige "github.com/mbrc12/prestige"
)

type TitleScene struct {
	color color.Color
}

func (scene *TitleScene) Update(state *State) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return state.Manager.Transition(&TicTacToeScene{}, prestige.NewFadeTransition(1))
	}
	return nil
}

func (scene *TitleScene) Draw(state *State, screen *ebiten.Image) {
	screen.Fill(scene.color)
	utils.DrawTextWithShadow(screen, "Tic Tac Toe", 300, 200, 4, color.White, text.AlignCenter, text.AlignCenter)
	utils.DrawTextWithShadow(screen, "Press SPACE to start.", 300, 400, 2, color.White, text.AlignCenter, text.AlignCenter)
}

func randomColor() (col color.Color) {
	f := func() uint8 { return uint8(rand.Int() % 255) }
	col = color.RGBA{f(), f(), f(), 255}
	return
}

func (scene *TitleScene) EnterStart(state *State) {
	scene.color = randomColor()
}

func (scene *TitleScene) EnterEnd(state *State) {
}

func (scene *TitleScene) ExitStart(state *State) {
}

func (scene *TitleScene) ExitEnd(state *State) {
}

func (scene *TitleScene) Layout(state *State, w, h int) (int, int) {
	return w, h
}
