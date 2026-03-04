// Package handcards 定義手牌 HandCards：持牌列表、加入/移除牌、是否為空、取得牌列表（供顯示與出牌使用）。
// 手牌依規格保持由小到大排序。
package handcards

import (
	"sort"

	"big2/internal/card"
)

// HandCards 手牌，最多 13 張，由小到大排序。
type HandCards struct {
	cards []card.Card
}

// New 建立空手牌。
func New() *HandCards {
	return &HandCards{cards: make([]card.Card, 0, 13)}
}

// AddCard 加入一張牌並維持排序。
func (h *HandCards) AddCard(c card.Card) {
	h.cards = append(h.cards, c)
	sort.Slice(h.cards, func(i, j int) bool { return h.cards[i].Compare(h.cards[j]) < 0 })
}

// RemoveCards 移除指定的牌（依牌面，不依索引）。若某張牌不在手牌中則忽略。
func (h *HandCards) RemoveCards(target []card.Card) {
	for _, t := range target {
		for i := range h.cards {
			if h.cards[i].Compare(t) == 0 {
				h.cards = append(h.cards[:i], h.cards[i+1:]...)
				break
			}
		}
	}
}

// IsEmpty 是否沒有手牌。
func (h *HandCards) IsEmpty() bool {
	return len(h.cards) == 0
}

// GetCards 回傳手牌副本（已排序）。不可修改以影響內部狀態。
func (h *HandCards) GetCards() []card.Card {
	out := make([]card.Card, len(h.cards))
	copy(out, h.cards)
	return out
}

// CardAt 回傳索引 i 的牌。若索引越界則回傳零值與 false。
func (h *HandCards) CardAt(i int) (card.Card, bool) {
	if i < 0 || i >= len(h.cards) {
		return card.Card{}, false
	}
	return h.cards[i], true
}

// Len 手牌張數。
func (h *HandCards) Len() int {
	return len(h.cards)
}
