package player

import (
	"cardkit/internal/card"
)

type ShowdownPlayer struct {
	BasePlayer[card.PokerCard]
	points int
}

func (p *ShowdownPlayer) GainPoint() {
	p.points++
}

func (p *ShowdownPlayer) GetPoints() int {
	return p.points
}

type AIShowdownPlayer struct {
	ShowdownPlayer
}

func NewShowdownPlayer(playerBehaviorStrategy PlayerBehaviorStrategy[card.PokerCard]) *ShowdownPlayer {
	return &ShowdownPlayer{
		BasePlayer: NewBasePlayer[card.PokerCard](playerBehaviorStrategy),
		points:     0,
	}
}
