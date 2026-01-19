package player

import (
	"cardkit/internal/card"
)

type UnoPlayer struct {
	BasePlayer[card.UnoCard]
}

type AIUnoPlayer struct {
	UnoPlayer
}

func NewUnoPlayer(playerBehaviorStrategy PlayerBehaviorStrategy) *UnoPlayer {
	return &UnoPlayer{
		BasePlayer: NewBasePlayer[card.UnoCard](playerBehaviorStrategy),
	}
}

// func (p *AIUnoPlayer) NameHimself() {
// 	name := AISelectName()
// 	p.BasePlayer.SetName(name)
// }

// func (p *AIUnoPlayer) PlayCard() card.Card {
// 	card, index := AISelectCard(p.GetHandCards())
// 	p.BasePlayer.RemoveHandCard(index)
// 	return card
// }

// type HumanUnoPlayer struct {
// 	UnoPlayer
// }

// func (p *HumanUnoPlayer) NameHimself() {
// 	name := HumanSelectName()
// 	p.BasePlayer.SetName(name)
// }

// func (p *HumanUnoPlayer) PlayCard() card.Card {
// 	card, index := HumanSelectCard(p.GetHandCards())
// 	if index == -1 {
// 		return nil
// 	}
// 	p.BasePlayer.RemoveHandCard(index - 1)
// 	return card
// }

// // NewAIUnoPlayer creates a new AI UNO player with an initialized hand
// func NewAIUnoPlayer() *AIUnoPlayer {
// 	return &AIUnoPlayer{
// 		UnoPlayer: UnoPlayer{
// 			BasePlayer: BasePlayer{
// 				hand: &hand.Hand{
// 					Cards: make([]card.Card, 0),
// 				},
// 			},
// 		},
// 	}
// }

// // NewHumanUnoPlayer creates a new human UNO player with an initialized hand
// func NewHumanUnoPlayer() *HumanUnoPlayer {
// 	return &HumanUnoPlayer{
// 		UnoPlayer: UnoPlayer{
// 			BasePlayer: BasePlayer{
// 				hand: &hand.Hand{
// 					Cards: make([]card.Card, 0),
// 				},
// 			},
// 		},
// 	}
// }
