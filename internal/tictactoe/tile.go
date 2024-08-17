package tictactoe

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	tileSize   = 100
	tileMargin = 4
)

var (
	tileColor       = color.NRGBA{0xee, 0xe4, 0xda, 0x59}
	tileHoverColor  = color.NRGBA{0xee, 0xe4, 0xda, 0x19}
	tileImage       = ebiten.NewImage(tileSize, tileSize)
	xColor          = color.NRGBA{0xff, 0x00, 0x00, 0xff}
	xHoverColor     = color.NRGBA{0xff, 0xcc, 0xcc, 0xff}
	oColor          = color.NRGBA{0x00, 0x00, 0xff, 0xff}
	oHoverColor     = color.NRGBA{0xcc, 0xcc, 0xff, 0xff}
	mplusFaceSource *text.GoTextFaceSource
)

func init() {
	tileImage.Fill(color.White)

	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

type Tile struct {
	Hovered bool
	Value   string
}

type TileInput struct {
	Hovered bool
	Pressed bool
	Turn    string // "X" or "O"
}

func (t *Tile) Update(input TileInput) error {
	t.Hovered = input.Hovered

	if t.Value != "" {
		return nil
	}

	if !t.Hovered {
		return nil
	}
	if input.Pressed {
		t.Value = input.Turn
	}

	return nil
}

func (t *Tile) Draw(boardImage *ebiten.Image, position Position, turn string) {
	t.drawTileBackground(boardImage, position)
	t.drawValue(boardImage, position, turn)
}

func (t *Tile) drawTileBackground(boardImage *ebiten.Image, position Position) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(position.x), float64(position.y))

	backgroundColor := tileColor
	if t.Value == "" && t.Hovered {
		backgroundColor = tileHoverColor
	}

	op.ColorScale.ScaleWithColor(backgroundColor)
	boardImage.DrawImage(tileImage, op)
}

func (t *Tile) drawValue(boardImage *ebiten.Image, position Position, turn string) {
	if t.Value == "" && !t.Hovered {
		return
	}

	str := t.Value
	if str == "" {
		str = turn
	}
	var color color.Color
	if str == "X" {
		if t.Hovered && t.Value == "" {
			color = xHoverColor
		} else {
			color = xColor
		}
	} else {
		if t.Hovered && t.Value == "" {
			color = oHoverColor
		} else {
			color = oColor
		}
	}

	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(float64(position.x)+float64(tileSize)/2, float64(position.y)+float64(tileSize)/2)
	textOp.ColorScale.ScaleWithColor(color)
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter
	text.Draw(boardImage, str, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   72,
	}, textOp)
}
