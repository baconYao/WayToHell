package game

import (
	"cardkit/internal/deck"
	"cardkit/pkg/logger"
)

type GameInterface[T any] interface {
	setup()
	setupPlayers()
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

// func (c *CardGameTemplate[T]) GetPlayers() []*player.Player {
// 	return c.Players
// }

// func (c *CardGameTemplate[T]) AddPlayer(players *player.Player) {
// 	c.Players = append(c.Players, players)
// }

// func (c *CardGameTemplate[T]) showPlayers() {
// 	for idx, player := range c.Players {
// 		fmt.Printf("玩家%d: %s\n", idx+1, player.GetName())
// 	}
// }

// Start 是核心 Template Method。其執行邏輯遵循 OOD 中粉紅色便條紙的定義:
// 1. Setup
// 2. 迴圈執行 TakeTurn，進行 Compare 直到 GameOver
// 3. ShowWinner
func (c *CardGameTemplate[T]) Start() {
	c.logger.Debug("開始進行遊戲...")
	c.Game.setup()
	c.Game.setupPlayers()
	c.Game.beforeTakeTurn()
	for !c.Game.isGameOver() {
		c.Game.takeTurn()
	}
	c.Game.afterTakeTurn()
	c.Game.showWinner()
}
