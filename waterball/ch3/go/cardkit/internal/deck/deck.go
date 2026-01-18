package deck

import (
	"math/rand"
	"time"

	"cardkit/internal/card"
)

type Deck struct {
	Cards []card.Card
}

func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) IsEmpty() bool {
	return len(d.Cards) == 0
}

func (d *Deck) Draw() card.Card {
	if d.IsEmpty() {
		return nil
	}

	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

func (d *Deck) Refill(cards []card.Card) {
	if len(cards) == 0 {
		return
	}

	d.Cards = append(d.Cards, cards...)
}
