package cardpatternhandler

import "big2/internal/card"

// Single 單張。
type Single struct{ Base }

func NewSingle() *Single { return &Single{} }

func (s *Single) Validate(cards []card.Card) Handler {
	if len(cards) != 1 {
		if s.next != nil {
			return s.next.Validate(cards)
		}
		return nil
	}
	s.cards = append([]card.Card(nil), cards...)
	return s
}

func (s *Single) Name() string { return "單張" }

func (s *Single) GetComparisonCard() card.Card {
	if len(s.cards) == 0 {
		return card.Card{}
	}
	return s.cards[0]
}

func (s *Single) IsSameType(other Handler) bool {
	if _, ok := other.(*Single); ok {
		return true
	}
	return false
}
