package game

import (
	"fmt"
	"go-snake/constants"
	"go-snake/objects"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Player            objects.Snake
	Apples            []*objects.Apple
	moveCounter       int
	spawnAppleCounter int
	GameOver          bool
	Score             int
}

func (g *Game) Update() error {
	if !g.GameOver {
		HandleSnakeInput(g)
		g.Player.HandleCollisions(&g.Apples, &g.GameOver, g.IncreaseScore)

		// FOR TEST DELETE LATER
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			objects.SpawnNewApple(&g.Apples)
		}
		//

		g.moveCounter++
		if g.moveCounter < constants.MoveDelay {
			return nil
		}
		g.moveCounter = 0

		g.Player.MovePlayer()

		g.spawnAppleCounter++
		if g.spawnAppleCounter < constants.AppleSpawnDelay {
			return nil
		}
		g.spawnAppleCounter = 0

		objects.SpawnNewApple(&g.Apples)
	} else {
		return nil
	}

	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(constants.Bg)

	// ebitenutil.DebugPrint(screen, "hello world")
	ebitenutil.DebugPrintAt(
		screen,
		"Score: "+strconv.Itoa(g.Score),
		constants.Boundary, constants.Boundary-constants.GridSize,
	)

	DrawGrid(screen)
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
