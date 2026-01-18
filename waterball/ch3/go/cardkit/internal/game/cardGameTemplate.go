package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"cardkit/internal/deck"
	"cardkit/internal/player"
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
	logger  *logger.Logger
	Players []*player.Player
	Deck    *deck.Deck
	Game    GameInterface[T]
}

func (c *CardGameTemplate[T]) GetPlayers() []*player.Player {
	return c.Players
}

func (c *CardGameTemplate[T]) AddPlayer(players *player.Player) {
	c.Players = append(c.Players, players)
}

func (c *CardGameTemplate[T]) showPlayers() {
	for idx, player := range c.Players {
		fmt.Printf("玩家%d: %s\n", idx+1, player.GetName())
	}
}

// askPlayerCount asks and returns the number of AI 和 Human Players
// requiredPlayers: total players required is this game
// Returns: aiCount, humanCount
func askPlayerCount(requiredPlayers int) (aiCount, humanCount int) {
	fmt.Printf("此遊戲需要 %d 位玩家。\n", requiredPlayers)

	scanner := bufio.NewScanner(os.Stdin)

	// 詢問 AI 玩家數量
	for {
		fmt.Print("請輸入要幾位 AI 玩家: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		count, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("無效的輸入，請輸入數字")
			continue
		}
		if count < 0 {
			fmt.Println("AI 玩家數量不能為負數")
			continue
		}
		if count > requiredPlayers {
			fmt.Printf("AI 玩家數量不能超過 %d 位\n", requiredPlayers)
			continue
		}
		aiCount = count
		break
	}

	humanCount = requiredPlayers - aiCount

	fmt.Printf("AI 玩家數量: %d, Human 玩家數量: %d\n", aiCount, humanCount)

	return aiCount, humanCount
}

// dealCards deals cards to all players
// players: list of players to deal cards to
// deck: the deck to draw cards from
// rounds: number of rounds to deal (each round gives one card to each player)
func dealCards(players []*player.Player, deck *deck.Deck, rounds int) {
	fmt.Printf("為 %d 位玩家各發放 %d 張起始手牌\n", len(players), rounds)
	for i := 0; i < rounds; i++ {
		for _, player := range players {
			player.AddHandCard(deck.Draw())
		}
	}
}

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
