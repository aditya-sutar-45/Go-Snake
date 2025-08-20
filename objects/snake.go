package objects

import (
	"go-snake/constants"
	inp "go-snake/inputs"
)

type Snake struct {
	Body []*Object
	Dir  inp.Direction
}

var (
	Up    = inp.Up
	Down  = inp.Down
	Left  = inp.Left
	Right = inp.Right
)

func MovePlayer(Player *Snake) {
	head := Player.Body[0]
	newHead := Object{
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
	Player.Body = append([]*Object{&newHead}, Player.Body...)
	Player.Body = Player.Body[:len(Player.Body)-1]
}

func getTailDir(body []*Object) inp.Direction {
	n := len(body)

	if n < 2 {
		return Up
	}

	last := body[n-1]
	secondLast := body[n-2]

	if last.X < secondLast.X {
		return Left
	} else if last.X > secondLast.X {
		return Right
	} else if last.Y < secondLast.Y {
		return Up
	} else {
		return Down
	}
}

func AddNewTail(Player *Snake) {
	last := Player.Body[len(Player.Body)-1]

	dirToAdd := getTailDir(Player.Body)

	newTail := Object{
		X: last.X,
		Y: last.Y,
		W: constants.GridSize,
		H: constants.GridSize,
	}

	if dirToAdd == Down {
		newTail.Y -= constants.GridSize
	}
	if dirToAdd == Up {
		newTail.Y += constants.GridSize
	}
	if dirToAdd == Right {
		newTail.X += constants.GridSize
	}
	if dirToAdd == Left {
		newTail.X -= constants.GridSize
	}

	Player.Body = append(Player.Body, &newTail)
	// fmt.Println("added tail")
}

func HandleCollisions(Player *Snake, Apples *[]*Apple, gameOver *bool, IncreaseScore func(int)) {
	head := Player.Body[0]
	outOfBounds := head.X <= constants.Boundary || head.X >= constants.GameX+constants.Boundary ||
		head.Y <= constants.Boundary || head.Y >= constants.GameY+constants.Boundary

	for ind := 0; ind < len(*Apples); ind++ {
		apple := (*Apples)[ind]
		if head.X == apple.Obj.X && head.Y == apple.Obj.Y {
			IncreaseScore(1)
			AddNewTail(Player)
			// remove the apple
			*Apples = append((*Apples)[:ind], (*Apples)[ind+1:]...)
			ind--
		}
	}

	if outOfBounds {
		*gameOver = true
	}
}
