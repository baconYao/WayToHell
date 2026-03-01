// Package play 定義出牌行為：Play 介面、NormalPlay（含牌型）、PassPlay，以及與頂牌比較 isStrongerThan。
package play

import (
	"big2/internal/card"
	cardpatternhandler "big2/internal/cardPatternHandler"
)

// Play 出牌：可能是正常牌型或 PASS。
type Play interface {
	GetPlayerIndex() int
	IsStrongerThan(other Play) bool
	GetCards() []card.Card
}

// NormalPlay 打出的牌型。儲存牌與比較用牌的快照，避免責任鏈被下一位玩家覆寫。
type NormalPlay struct {
	PlayerIndex   int
	CardsSnapshot []card.Card
	CompareCard   card.Card
	PatternName   string
}

func (n *NormalPlay) GetPlayerIndex() int { return n.PlayerIndex }

func (n *NormalPlay) GetCards() []card.Card {
	return n.CardsSnapshot
}

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
		PlayerIndex:   playerIndex,
		CardsSnapshot: cp,
		CompareCard:   h.GetComparisonCard(),
		PatternName:   h.Name(),
	}
}

// PassPlay 放棄出牌。
type PassPlay struct {
	PlayerIndex int
}

func (p *PassPlay) GetPlayerIndex() int   { return p.PlayerIndex }
func (p *PassPlay) GetCards() []card.Card { return nil }

func (p *PassPlay) IsStrongerThan(other Play) bool {
	return false
}
