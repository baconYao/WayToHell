package deck

import (
	"math/rand"
	"time"

	"cardkit/internal/card"
)

type Deck[T card.Card] struct {
	Cards []T
}

func (d *Deck[T]) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck[T]) IsEmpty() bool {
	return len(d.Cards) == 0
}

func (d *Deck[T]) Draw() T {
	if d.IsEmpty() {
		var zero T
		return zero
	}

	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

func (d *Deck[T]) Refill(cards []T) {
	if len(cards) == 0 {
		return
	}

	d.Cards = append(d.Cards, cards...)
}
