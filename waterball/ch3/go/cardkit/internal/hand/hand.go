package hand

import (
	"errors"

	"cardkit/internal/card"
)

type Hand[T card.Card] struct {
	Cards []T
}

func (h *Hand[T]) Add(c T) {
	h.Cards = append(h.Cards, c)
}

func (h *Hand[T]) GetCards() []T {
	replica := make([]T, len(h.Cards))
	copy(replica, h.Cards)
	return replica
}

func (h *Hand[T]) Remove(index int) error {
	if index < 0 || index >= len(h.Cards) {
		return errors.New("index out of range of hand cards")
	}

	h.Cards = append(h.Cards[:index], h.Cards[index+1:]...)
	return nil
}
