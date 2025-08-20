package game

import (
	inp "go-snake/inputs"

	"github.com/hajimehoshi/ebiten/v2"
)

func HandleSnakeInput(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyW) && g.Player.Dir != inp.Down {
		g.Player.Dir = inp.Up
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && g.Player.Dir != inp.Up {
		g.Player.Dir = inp.Down
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) && g.Player.Dir != inp.Right {
		g.Player.Dir = inp.Left
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && g.Player.Dir != inp.Left {
		g.Player.Dir = inp.Right
	}
}
