package play

import (
	"big2/internal/card"
	cardpatternhandler "big2/internal/cardPatternHandler"
)

// NormalPlay 打出的牌型；持有驗證通過的 CardPatternHandler，牌與比較用牌於建立時複製以避免責任鏈被下一位玩家覆寫。
type NormalPlay struct {
	handler     cardpatternhandler.Handler
	PlayerIndex int
	Cards       []card.Card
	CompareCard card.Card
}

func (n *NormalPlay) GetPlayerIndex() int { return n.PlayerIndex }

func (n *NormalPlay) GetCards() []card.Card {
	return n.Cards
}

// PatternName 回傳此手牌的牌型名稱
func (n *NormalPlay) PatternName() string {
	return n.handler.Name()
}

// Handler 回傳此手牌的牌型 handler
func (n *NormalPlay) Handler() cardpatternhandler.Handler {
	return n.handler
}

// IsStrongerThan 比較牌型是否更強。
// other 為 topPlay 的牌；若 other 為 nil 或 PassPlay 則回傳 true。
func (n *NormalPlay) IsStrongerThan(other Play) bool {
	if other == nil {
		return true
	}
	if _, ok := other.(*PassPlay); ok {
		return true
	}
	otherNorm, ok := other.(*NormalPlay)
	if !ok {
		return false
	}
	if !n.handler.IsSameType(otherNorm.Handler()) {
		return false
	}
	return n.CompareCard.Compare(otherNorm.CompareCard) > 0
}

// NewNormalPlay 從驗證通過的 handler 建立 NormalPlay，持有 handler 參考並複製牌與比較牌。
func NewNormalPlay(playerIndex int, h cardpatternhandler.Handler) *NormalPlay {
	cards := h.Cards()
	cp := make([]card.Card, len(cards))
	copy(cp, cards)
	return &NormalPlay{
		handler:     h,
		PlayerIndex: playerIndex,
		Cards:       cp,
		CompareCard: h.GetComparisonCard(),
	}
}
