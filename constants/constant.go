package constants

import "image/color"

const (
	SHeight         = 900
	SWidth          = 900
	GameX           = 800
	GameY           = 800
	Boundary        = (SHeight - GameX) / 2
	GridSize        = 20
	MoveDelay       = 13
	MinMoveDelay    = 7
	AppleSpawnDelay = 10
)

var (
	Bg        = color.RGBA{17, 20, 22, 1}
	GridBg    = color.RGBA{17, 22, 22, 1}
	Secondary = color.RGBA{72, 78, 75, 1}
	Primary   = color.RGBA{84, 164, 84, 1}
	Accent    = color.RGBA{227, 246, 227, 1}
	Red       = color.RGBA{227, 84, 84, 1}
)
