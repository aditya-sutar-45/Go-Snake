package game

import (
	"fmt"
	"go-snake/constants"
	"go-snake/objects"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Player      objects.Snake
	Apples      []*objects.Apple
	moveCounter int
	GameOver    bool
	Score       int
}

func (g *Game) Update() error {
	if !g.GameOver {
		HandleSnakeInput(g)
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			objects.AddNewTail(&g.Player)
		}

		objects.HandleCollisions(&g.Player, &g.Apples, &g.GameOver, g.IncreaseScore)

		g.moveCounter++
		if g.moveCounter < constants.MoveDelay {
			return nil
		}
		g.moveCounter = 0

		objects.MovePlayer(&g.Player)
	} else {
		return nil
	}

	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(constants.Bg)

	DrawBoundaries(screen)

	DrawPlayer(screen, g)

	DrawApples(screen, g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.SWidth, constants.SHeight
}

func (g *Game) IncreaseScore(amountToIncrase int) {
	g.Score += amountToIncrase
	fmt.Printf("Score: %d\n", g.Score)
}
