package play

import (
	"big2/internal/card"
	cardpatternhandler "big2/internal/cardPatternHandler"
)

// NormalPlay 打出的牌型。儲存牌與比較用牌的快照，避免責任鏈被下一位玩家覆寫。
type NormalPlay struct {
	PlayerIndex int
	Cards       []card.Card
	CompareCard card.Card
	PatternName string
}

func (n *NormalPlay) GetPlayerIndex() int { return n.PlayerIndex }

func (n *NormalPlay) GetCards() []card.Card {
	return n.Cards
}

// IsStrongerThan 比較牌型是否更強.
// other 為 topPlay 的牌, 如果 other 為 nil 或 PassPlay, 則返回 true.
func (n *NormalPlay) IsStrongerThan(other Play) bool {
	// 當前玩家的牌比topPlay的牌更強
	if other == nil {
		return true
	}
	if _, ok := other.(*PassPlay); ok {
		return true
	}
	// 比較牌型是否相同
	otherNorm, ok := other.(*NormalPlay)
	if !ok {
		return false
	}
	if n.PatternName != otherNorm.PatternName {
		return false
	}
	return n.CompareCard.Compare(otherNorm.CompareCard) > 0
}

// NewNormalPlay 從驗證通過的 handler 建立 NormalPlay（複製牌與比較牌，不保留 handler 參考）。
func NewNormalPlay(playerIndex int, h cardpatternhandler.Handler) *NormalPlay {
	cards := h.Cards()
	cp := make([]card.Card, len(cards))
	copy(cp, cards)
	return &NormalPlay{
		PlayerIndex: playerIndex,
		Cards:       cp,
		CompareCard: h.GetComparisonCard(),
		PatternName: h.Name(),
	}
}
