package showdown

import (
	"errors"
	"math/rand"
	"time"
)

var ErrDeckEmpty = errors.New("deck is empty, cannot draw a card")

type Deck struct {
	cards []Card
	top   int
}

// NewDeck constructor, creates a deck with 52 unique cards
func NewDeck() *Deck {
	deck := &Deck{
		cards: make([]Card, 0, 52), // Pre-allocate capacity for 52 cards
	}

	// Iterate through all Rank and Suit combinations
	for rank := Rank2; rank <= RankA; rank++ {
		for suit := Club; suit <= Spade; suit++ {
			deck.cards = append(deck.cards, Card{
				rank: rank,
				suit: suit,
			})
		}
	}

	return deck
}

// GetCards returns a copy of the cards in the deck
func (d *Deck) GetCards() []Card {
	// Return a copy to prevent external modification
	cardsCopy := make([]Card, len(d.cards))
	copy(cardsCopy, d.cards)
	return cardsCopy
}

// Shuffle randomizes the order of cards in the deck
func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// DrawCard removes and returns the first card from the deck
func (d *Deck) DrawCard() (Card, error) {
	if d.top >= len(d.cards) {
		return Card{}, ErrDeckEmpty
	}
	card := d.cards[d.top]
	d.top++
	return card, nil
}
