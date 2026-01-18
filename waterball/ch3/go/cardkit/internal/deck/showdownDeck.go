package deck

import (
	"cardkit/internal/card"
)

func NewShowdownDeck() *Deck {
	cards := make([]card.Card, 0, 52)

	suits := []card.Suit{card.Club, card.Diamond, card.Heart, card.Spade}
	ranks := []card.Rank{
		card.Two, card.Three, card.Four, card.Five, card.Six,
		card.Seven, card.Eight, card.Nine, card.Ten,
		card.Jack, card.Queen, card.King, card.Ace,
	}

	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, card.PokerCard{
				Suit: suit,
				Rank: rank,
			})
		}
	}

	return &Deck{
		Cards: cards,
	}
}
