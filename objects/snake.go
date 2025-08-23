package objects

import (
	"go-snake/constants"
	inp "go-snake/inputs"
)

type Snake struct {
	Body                    []*Object
	Dir                     inp.Direction
	Growing                 bool
	CurrentMoveDelay        float32
	DelayDecreaseMultiplier float32
}

var (
	Up    = inp.Up
	Down  = inp.Down
	Left  = inp.Left
	Right = inp.Right
)

func (Player *Snake) MovePlayer() {
	head := Player.Body[0]
	newHead := &Object{
		X: head.X,
		Y: head.Y,
		W: constants.GridSize,
		H: constants.GridSize,
	}
	if Player.Dir == Up {
		newHead.Y -= constants.GridSize
	}
	if Player.Dir == Down {
		newHead.Y += constants.GridSize
	}
	if Player.Dir == Left {
		newHead.X -= constants.GridSize
	}
	if Player.Dir == Right {
		newHead.X += constants.GridSize
	}

	// new head to the start
	Player.Body = append([]*Object{newHead}, Player.Body...)

	if Player.Growing {
		// if player is growning no need to trim the end of the array
		// and set growing to false
		Player.CurrentMoveDelay *= Player.DelayDecreaseMultiplier
		if Player.CurrentMoveDelay <= constants.MinMoveDelay {
			Player.CurrentMoveDelay = constants.MinMoveDelay
		}
		Player.Growing = false
	} else {
		// if player is not growing remove the last element
		Player.Body = Player.Body[:len(Player.Body)-1]
	}
}

func (Player *Snake) HandleCollisions(Apples *[]*Apple, gameOver *bool, IncreaseScore func(int)) {
	// collisions that end the game
	head := Player.Body[0]
	outOfBounds := head.X <= constants.Boundary || head.X >= constants.GameX+constants.Boundary ||
		head.Y <= constants.Boundary || head.Y >= constants.GameY+constants.Boundary

	collidedWithBody := false

	for i := 1; i < len(Player.Body); i++ {
		body := Player.Body[i]

		if head.X == body.X && head.Y == body.Y {
			collidedWithBody = true
			break
		}
	}

	if outOfBounds || collidedWithBody {
		*gameOver = true
	}

	// apple collisions
	for ind := 0; ind < len(*Apples); ind++ {
		apple := (*Apples)[ind]
		if head.X == apple.Obj.X && head.Y == apple.Obj.Y {
			Player.Growing = true
			IncreaseScore(1)
			// remove the apple
			*Apples = append((*Apples)[:ind], (*Apples)[ind+1:]...)
			ind--
		}
	}
}
