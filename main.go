package main

import (
	"go-snake/constants"
	"go-snake/game"
	"go-snake/inputs"
	"go-snake/objects"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(constants.SWidth, constants.SHeight)
	ebiten.SetWindowTitle("go go snake")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g := &game.Game{
		Player: objects.Snake{
			Body: []*objects.Object{
				{
					X: constants.SWidth / 2,
					Y: constants.SHeight / 2,
					H: constants.GridSize,
					W: constants.GridSize,
				},
			},
			Dir:                     inputs.Up,
			Growing:                 false,
			CurrentMoveDelay:        constants.MoveDelay,
			DelayDecreaseMultiplier: 0.95,
		},
		Apples: []*objects.Apple{
			{
				Obj: objects.Object{
					X: 350,
					Y: 450,
					W: constants.GridSize,
					H: constants.GridSize,
				},
				Color: constants.Red,
			},
		},
		GameOver: false,
		Score:    0,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
