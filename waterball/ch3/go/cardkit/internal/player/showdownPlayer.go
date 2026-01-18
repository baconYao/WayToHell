package player

import (
	"cardkit/internal/card"
	"cardkit/internal/hand"
)

type ShowdownPlayer struct {
	BasePlayer
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

func NewShowdownPlayer(playerBehaviorStrategy PlayerBehaviorStrategy) *ShowdownPlayer {
	return &ShowdownPlayer{
		BasePlayer: BasePlayer{
			hand:             &hand.Hand{Cards: make([]card.Card, 0)},
			behaviorStrategy: playerBehaviorStrategy,
		},
	}
}

// func (p *AIShowdownPlayer) NameHimself() {
// 	name := AISelectName()
// 	p.BasePlayer.SetName(name)
// }

// func (p *AIShowdownPlayer) PlayCard() card.Card {
// 	card, index := AISelectCard(p.GetHandCards())
// 	p.BasePlayer.RemoveHandCard(index)
// 	return card
// }

// type HumanShowdownPlayer struct {
// 	ShowdownPlayer
// }

// func (p *HumanShowdownPlayer) NameHimself() {
// 	name := HumanSelectName()
// 	p.BasePlayer.SetName(name)
// }

// func (p *HumanShowdownPlayer) PlayCard() card.Card {
// 	card, index := HumanSelectCard(p.GetHandCards())
// 	if index == -1 {
// 		return nil
// 	}
// 	p.BasePlayer.RemoveHandCard(index - 1)
// 	return card
// }

// // NewAIShowdownPlayer creates a new AI showdown player with an initialized hand
// func NewAIShowdownPlayer() *AIShowdownPlayer {
// 	return &AIShowdownPlayer{
// 		ShowdownPlayer: ShowdownPlayer{
// 			BasePlayer: BasePlayer{
// 				hand: &hand.Hand{
// 					Cards: make([]card.Card, 0),
// 				},
// 			},
// 		},
// 	}
// }

// // NewHumanShowdownPlayer creates a new human showdown player with an initialized hand
// func NewHumanShowdownPlayer() *HumanShowdownPlayer {
// 	return &HumanShowdownPlayer{
// 		ShowdownPlayer: ShowdownPlayer{
// 			BasePlayer: BasePlayer{
// 				hand: &hand.Hand{
// 					Cards: make([]card.Card, 0),
// 				},
// 			},
// 		},
// 	}
// }
