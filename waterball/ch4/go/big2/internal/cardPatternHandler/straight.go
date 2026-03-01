package cardpatternhandler

import (
	"sort"

	"big2/internal/card"
)

// Straight 順子（五張連續數字，大老二規則含 K-A-2-3-4 等）。
type Straight struct{ Base }

func NewStraight() *Straight { return &Straight{} }

func rankIndex(r card.Rank) int { return int(r) }

func (st *Straight) Validate(cards []card.Card) Handler {
	if len(cards) != 5 {
		if st.next != nil {
			return st.next.Validate(cards)
		}
		return nil
	}
	sortCards(cards)
	idx := make([]int, 5)
	for i := range cards {
		idx[i] = rankIndex(cards[i].Rank)
	}
	if !isStraightRanks(idx) {
		if st.next != nil {
			return st.next.Validate(cards)
		}
		return nil
	}
	st.cards = append([]card.Card(nil), cards...)
	return st
}

func isStraightRanks(idx []int) bool {
	s := make([]int, 5)
	copy(s, idx)
	sort.Ints(s)
	if s[4]-s[0] == 4 {
		for i := 0; i < 4; i++ {
			if s[i+1]-s[i] != 1 {
				return false
			}
		}
		return true
	}
	if s[0] == 0 && s[4] == 12 && s[3]-s[0] == 3 {
		return s[1] == 1 && s[2] == 2
	}
	return false
}

func (st *Straight) Name() string { return "順子" }

func (st *Straight) GetComparisonCard() card.Card {
	if len(st.cards) == 0 {
		return card.Card{}
	}
	max := st.cards[0]
	for i := 1; i < len(st.cards); i++ {
		if st.cards[i].Compare(max) > 0 {
			max = st.cards[i]
		}
	}
	return max
}

func (st *Straight) IsSameType(other Handler) bool {
	if _, ok := other.(*Straight); ok {
		return true
	}
	return false
}
