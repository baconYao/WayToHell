package game

import "math/rand"

type AI struct {
	player
}

func NewAI(number int) Player {
	return &AI{player: player{number: number}}
}

func (a AI) Decide() Decision {
	randomNum := rand.Intn(3)
	switch randomNum {
	case 0:
		return Scissors
	case 1:
		return Paper
	default:
		return Stone
	}
}
