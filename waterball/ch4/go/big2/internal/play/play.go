// Package play 定義出牌行為：Play 介面、NormalPlay（含牌型）、PassPlay，以及與頂牌比較 isStrongerThan。
package play

import "big2/internal/card"

// Play 出牌：可能是正常牌型或 PASS。
type Play interface {
	GetPlayerIndex() int
	IsStrongerThan(other Play) bool
	GetCards() []card.Card
}
