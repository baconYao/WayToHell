package cardpatternhandler

import "big2/internal/card"

// Pair 對子。
type Pair struct{ Base }

func NewPair() *Pair { return &Pair{} }

func (p *Pair) Validate(cards []card.Card) Handler {
	if len(cards) != 2 {
		if p.next != nil {
			return p.next.Validate(cards)
		}
		return nil
	}
	card.SortCards(cards)
	if cards[0].Rank != cards[1].Rank {
		if p.next != nil {
			return p.next.Validate(cards)
		}
		return nil
	}
	p.cards = append([]card.Card(nil), cards...)
	return p
}

func (p *Pair) Name() string { return "對子" }

func (p *Pair) GetComparisonCard() card.Card {
	if len(p.cards) < 2 {
		return card.Card{}
	}
	if p.cards[1].Compare(p.cards[0]) > 0 {
		return p.cards[1]
	}
	return p.cards[0]
}

func (p *Pair) IsSameType(other Handler) bool {
	if _, ok := other.(*Pair); ok {
		return true
	}
	return false
}
