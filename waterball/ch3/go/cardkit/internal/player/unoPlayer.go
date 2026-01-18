package player

import (
	"cardkit/internal/card"
	"cardkit/internal/hand"
)

type UnoPlayer struct {
	Player
}

type AIUnoPlayer struct {
	UnoPlayer
}

func (p *AIUnoPlayer) NameHimself() {
	name := AISelectName()
	p.SetName(name)
}

func (p *AIUnoPlayer) PlayCard() card.Card {
	card, index := AISelectCard(p.GetHandCards())
	p.RemoveHandCard(index)
	return card
}

type HumanUnoPlayer struct {
	UnoPlayer
}

func (p *HumanUnoPlayer) NameHimself() {
	name := HumanSelectName()
	p.SetName(name)
}

func (p *HumanUnoPlayer) PlayCard() card.Card {
	card, index := HumanSelectCard(p.GetHandCards())
	if index == -1 {
		return nil
	}
	p.RemoveHandCard(index - 1)
	return card
}

// NewAIUnoPlayer creates a new AI UNO player with an initialized hand
func NewAIUnoPlayer() *AIUnoPlayer {
	return &AIUnoPlayer{
		UnoPlayer: UnoPlayer{
			Player: Player{
				hand: &hand.Hand{
					Cards: make([]card.Card, 0),
				},
			},
		},
	}
}

// NewHumanUnoPlayer creates a new human UNO player with an initialized hand
func NewHumanUnoPlayer() *HumanUnoPlayer {
	return &HumanUnoPlayer{
		UnoPlayer: UnoPlayer{
			Player: Player{
				hand: &hand.Hand{
					Cards: make([]card.Card, 0),
				},
			},
		},
	}
}
