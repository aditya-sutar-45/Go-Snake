package game

import (
	"go-snake/constants"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawGrid(screen *ebiten.Image) {
	for x := constants.Boundary + constants.GridSize; x <= constants.GameX+constants.Boundary-constants.GridSize; x += constants.GridSize {
		vector.StrokeLine(
			screen,
			float32(x), constants.Boundary+constants.GridSize,
			float32(x), constants.GameY+constants.Boundary,
			2, constants.GridBg,
			false,
		)
	}
	for y := constants.Boundary + constants.GridSize; y <= constants.GameY+constants.Boundary-constants.GridSize; y += constants.GridSize {
		vector.StrokeLine(
			screen,
			constants.Boundary+constants.GridSize, float32(y),
			constants.GameY+constants.Boundary, float32(y),
			2, constants.GridBg,
			false,
		)
	}

}

func DrawBoundaries(screen *ebiten.Image) {
	// top constants.Boundary
	vector.DrawFilledRect(screen,
		constants.Boundary+constants.GridSize, constants.Boundary,
		constants.GameX-constants.GridSize, constants.GridSize,
		constants.Secondary,
		false,
	)
	// bottom constants.Boundary
	vector.DrawFilledRect(screen,
		constants.Boundary, constants.GameY+constants.Boundary,
		constants.GameX, constants.GridSize,
		constants.Secondary,
		false,
	)
	// left constants.Boundary
	vector.DrawFilledRect(screen,
		constants.Boundary, constants.Boundary,
		constants.GridSize, constants.GameY,
		constants.Secondary,
		false,
	)
	// right constants.Boundary
	vector.DrawFilledRect(screen,
		constants.GameX+constants.Boundary, constants.Boundary,
		constants.GridSize, constants.GameY+constants.GridSize,
		constants.Secondary,
		false,
	)
}

func DrawPlayer(screen *ebiten.Image, g *Game) {
	for i, obj := range g.Player.Body {
		if i == 0 {
			vector.DrawFilledRect(
				screen,
				float32(obj.X), float32(obj.Y),
				float32(obj.W), float32(obj.H),
				constants.Accent,
				false,
			)
		}
		vector.DrawFilledRect(
			screen,
			float32(obj.X), float32(obj.Y),
			float32(obj.W), float32(obj.H),
			constants.Primary,
			false,
		)
	}
}

func DrawApples(screen *ebiten.Image, g *Game) {
	for _, apple := range g.Apples {
		vector.DrawFilledRect(
			screen,
			float32(apple.Obj.X), float32(apple.Obj.Y),
			float32(apple.Obj.W), float32(apple.Obj.H),
			apple.Color,
			false,
		)
	}
}
