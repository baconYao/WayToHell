package player

import (
	"cardkit/internal/card"
	"cardkit/internal/hand"
)

type ShowdownPlayer struct {
	Player
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

func (p *AIShowdownPlayer) NameHimself() {
	name := AISelectName()
	p.SetName(name)
}

func (p *AIShowdownPlayer) PlayCard() card.Card {
	card, index := AISelectCard(p.GetHandCards())
	p.RemoveHandCard(index)
	return card
}

type HumanShowdownPlayer struct {
	ShowdownPlayer
}

func (p *HumanShowdownPlayer) NameHimself() {
	name := HumanSelectName()
	p.SetName(name)
}

func (p *HumanShowdownPlayer) PlayCard() card.Card {
	card, index := HumanSelectCard(p.GetHandCards())
	if index == -1 {
		return nil
	}
	p.RemoveHandCard(index - 1)
	return card
}

// NewAIShowdownPlayer creates a new AI showdown player with an initialized hand
func NewAIShowdownPlayer() *AIShowdownPlayer {
	return &AIShowdownPlayer{
		ShowdownPlayer: ShowdownPlayer{
			Player: Player{
				hand: &hand.Hand{
					Cards: make([]card.Card, 0),
				},
			},
		},
	}
}

// NewHumanShowdownPlayer creates a new human showdown player with an initialized hand
func NewHumanShowdownPlayer() *HumanShowdownPlayer {
	return &HumanShowdownPlayer{
		ShowdownPlayer: ShowdownPlayer{
			Player: Player{
				hand: &hand.Hand{
					Cards: make([]card.Card, 0),
				},
			},
		},
	}
}
