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
	card.SortCards(cards)
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
	// 非環狀：五張連續，如 0,1,2,3,4
	if s[4]-s[0] == 4 {
		for i := 0; i < 4; i++ {
			if s[i+1]-s[i] != 1 {
				return false
			}
		}
		return true
	}
	// 環狀：含 0 與 12，四段間隔中恰有一段為 9（跨 2→3），其餘為 1
	// 例如 0,1,2,3,12（2-3-4-5-6）、0,9,10,11,12（Q-K-A-2-3）等
	if s[0] != 0 || s[4] != 12 {
		return false
	}
	g1 := s[1] - s[0]
	g2 := s[2] - s[1]
	g3 := s[3] - s[2]
	g4 := s[4] - s[3]
	// 四段中恰有一段為 9，其餘為 1
	nine, one := 0, 0
	for _, g := range []int{g1, g2, g3, g4} {
		if g == 9 {
			nine++
		} else if g == 1 {
			one++
		}
	}
	return nine == 1 && one == 3
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
