package game

import (
	"cardkit/internal/deck"
	"cardkit/pkg/logger"
)

type GameInterface[T any] interface {
	setup()
	beforeTakeTurn()
	takeTurn()
	afterTakeTurn()
	isGameOver() bool
	showWinner()
	compare() T
}

type CardGameTemplate[T any] struct {
	logger *logger.Logger
	Deck   *deck.Deck
	Game   GameInterface[T]
}

// Start 是核心 Template Method。其執行邏輯遵循 OOD 中粉紅色便條紙的定義:
// 1. Setup
// 2. 迴圈執行 TakeTurn，進行 Compare 直到 GameOver
// 3. ShowWinner
func (c *CardGameTemplate[T]) Start() {
	c.logger.Debug("開始進行遊戲...")
	c.Game.setup()
	for !c.Game.isGameOver() {
		c.Game.beforeTakeTurn()
		c.Game.takeTurn()
		c.Game.compare()
		c.Game.afterTakeTurn()
	}
	c.Game.showWinner()
}
