package objects

import (
	"go-snake/constants"
	"image/color"
	"math/rand"
)

type Apple struct {
	Obj   Object
	Color color.RGBA
}

func getRandomPos() (int, int) {
	min := constants.Boundary + constants.GridSize
	maxX := constants.GameX + constants.Boundary - constants.GridSize
	maxY := constants.GameY + constants.Boundary - constants.GridSize

	gridX := rand.Intn((maxX-min)/constants.GridSize + 1)
	gridY := rand.Intn((maxY-min)/constants.GridSize + 1)

	x := min + gridX*constants.GridSize
	y := min + gridY*constants.GridSize

	return x, y
}

func SpawnNewApple(apples *[]*Apple) {
	randX, randY := getRandomPos()

	newApple := Apple{
		Obj: Object{
			X: randX,
			Y: randY,
			H: constants.GridSize,
			W: constants.GridSize,
		},
		Color: constants.Red,
	}

	*apples = append(*apples, &newApple)
}
