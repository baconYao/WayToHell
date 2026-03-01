// Package cardpattern 實作牌型責任鏈（CardPatternHandler）：驗證選牌是否為合法牌型，
// 並支援單張、對子、順子、葫蘆等。擴充新牌型時僅新增 Handler，符合 OCP。
package cardpatternhandler

import (
	"sort"

	"big2/internal/card"
)

// Handler 牌型責任鏈節點：驗證選牌、回傳比較用牌、判斷是否同類型。
type Handler interface {
	SetNext(h Handler)
	Validate(cards []card.Card) Handler
	Name() string
	GetComparisonCard() card.Card
	IsSameType(other Handler) bool
	Cards() []card.Card
}

// Base 共通邏輯：串接下一節點、持有選牌。
type Base struct {
	next  Handler
	cards []card.Card
}

func (b *Base) SetNext(h Handler) { b.next = h }

func (b *Base) Cards() []card.Card {
	out := make([]card.Card, len(b.cards))
	copy(out, b.cards)
	return out
}

// sortCards 依牌大小排序（同 package 內 concrete handler 使用）。
func sortCards(cards []card.Card) {
	sort.Slice(cards, func(i, j int) bool { return cards[i].Compare(cards[j]) < 0 })
}

// BuildChain 建立責任鏈：Single -> Pair -> Straight -> FullHouse，回傳鏈頭。
func BuildChain() Handler {
	single := NewSingle()
	pair := NewPair()
	straight := NewStraight()
	full := NewFullHouse()
	single.SetNext(pair)
	pair.SetNext(straight)
	straight.SetNext(full)
	return single
}
