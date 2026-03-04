package cardpatternhandler

import "big2/internal/card"

// FullHouse 葫蘆（三張同數字 + 兩張同數字）。
type FullHouse struct{ Base }

func NewFullHouse() *FullHouse { return &FullHouse{} }

func (f *FullHouse) Validate(cards []card.Card) Handler {
	if len(cards) != 5 {
		if f.next != nil {
			return f.next.Validate(cards)
		}
		return nil
	}
	card.SortCards(cards)
	rankCount := make(map[card.Rank]int)
	for _, c := range cards {
		rankCount[c.Rank]++
	}
	hasThree, hasTwo := false, false
	for _, n := range rankCount {
		if n == 3 {
			hasThree = true
		}
		if n == 2 {
			hasTwo = true
		}
	}
	if !hasThree || !hasTwo {
		if f.next != nil {
			return f.next.Validate(cards)
		}
		return nil
	}
	f.cards = append([]card.Card(nil), cards...)
	return f
}

func (f *FullHouse) Name() string { return "葫蘆" }

func (f *FullHouse) GetComparisonCard() card.Card {
	if len(f.cards) == 0 {
		return card.Card{}
	}
	rankCount := make(map[card.Rank][]card.Card)
	for _, c := range f.cards {
		rankCount[c.Rank] = append(rankCount[c.Rank], c)
	}
	for _, list := range rankCount {
		if len(list) == 3 {
			max := list[0]
			for i := 1; i < 3; i++ {
				if list[i].Compare(max) > 0 {
					max = list[i]
				}
			}
			return max
		}
	}
	return f.cards[0]
}

func (f *FullHouse) IsSameType(other Handler) bool {
	if _, ok := other.(*FullHouse); ok {
		return true
	}
	return false
}
