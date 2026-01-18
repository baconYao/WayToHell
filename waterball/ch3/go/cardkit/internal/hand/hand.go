package hand

import (
	"errors"

	"cardkit/internal/card"
)

type Hand struct {
	Cards []card.Card
}

func (h *Hand) Add(c card.Card) {
	h.Cards = append(h.Cards, c)
}

func (h *Hand) Size() int {
	return len(h.Cards)
}

func (h *Hand) GetCards() []card.Card {
	replica := make([]card.Card, len(h.Cards))
	copy(replica, h.Cards)
	return replica
}

func (h *Hand) Remove(index int) error {
	if index < 0 || index >= len(h.Cards) {
		return errors.New("index out of range of hand cards")
	}

	h.Cards = append(h.Cards[:index], h.Cards[index+1:]...)
	return nil
}
