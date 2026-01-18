package deck

import "cardkit/internal/card"

func NewUnoDeck() *Deck {
	cards := make([]card.Card, 0, 40)

	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			cards = append(cards, card.UnoCard{
				Color:  card.Color(i),
				Number: j,
			})
		}
	}

	return &Deck{
		Cards: cards,
	}
}
