package game

import (
	"cardkit/pkg/logger"
)

type GameInterface interface {
	setup()
	// beforeTakeTurn()
	takeTurn()
	// afterTakeTurn()
	isGameOver() bool
	showWinner()
}

type CardGameTemplate struct {
	logger *logger.Logger
	Game   GameInterface
}

// Start 是核心 Template Method。其執行邏輯遵循 OOD 中粉紅色便條紙的定義:
// 1. Setup
// 2. 迴圈執行 TakeTurn，進行 Compare 直到 GameOver
// 3. ShowWinner
func (c *CardGameTemplate) Start() {
	c.logger.Debug("開始進行遊戲...")
	c.Game.setup()
	for !c.Game.isGameOver() {
		// c.Game.beforeTakeTurn() // 每回合的前處理
		c.Game.takeTurn() // 每回合的主要邏輯處理
		// c.Game.afterTakeTurn()  // 每回合的後處理
	}
	c.Game.showWinner()
}
