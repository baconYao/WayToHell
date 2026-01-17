package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/baconYao/WayToHell/waterball/ch3/go/cardkit/internal/deck"
	"github.com/baconYao/WayToHell/waterball/ch3/go/cardkit/internal/player"
)

// TODO: 遊戲框架核心 - 樣板方法模式
// 定義 Game 介面和樣板方法實作
// 包含：start(), setup(), beforeTakeTurn(), takeTurn(), afterTakeTurn(), isGameOver(), showWinner(), compare()
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
	Players []*player.Player
	Deck    *deck.Deck
	Game    GameInterface[T]
}

// Setup is a placeholder (abstract method) for the subclass to implement
// func (g *CardGameTemplate[T]) setup() {
// 	// not implemented
// 	panic("Setup is not implemented")
// }

// // SetupPlayers is a placeholder (abstract method) for the subclass to implement
// func (g *CardGameTemplate[T]) setupPlayers() {
// 	panic("SetupPlayers is not implemented")
// }

// // BeforeTakeTurn is a placeholder (abstract method) for the subclass to implement
// func (g *CardGameTemplate[T]) beforeTakeTurn() {
// 	// not implemented
// 	panic("BeforeTakeTurn is not implemented")
// }

// // TakeTurn is a placeholder (abstract method) for the subclass to implement
// func (g *CardGameTemplate[T]) takeTurn() {
// 	panic("TakeTurn is not implemented")
// }

// // AfterTakeTurn is a placeholder (abstract method) for the subclass to implement
// func (g *CardGameTemplate[T]) afterTakeTurn() {
// 	panic("AfterTakeTurn is not implemented")
// }

// // IsGameOver is a placeholder (abstract method) for the subclass to implement
// func (g *CardGameTemplate[T]) isGameOver() bool {
// 	// not implemented
// 	panic("IsGameOver is not implemented")
// }

// // ShowWinner is a placeholder (abstract method) for the subclass to implement
// func (g *CardGameTemplate[T]) showWinner() {
// 	panic("ShowWinner is not implemented")
// }

// // Compare is a placeholder (abstract method) for the subclass to implement
// func (g *CardGameTemplate[T]) compare() T {
// 	panic("Compare is not implemented")
// }

func (g *CardGameTemplate[T]) GetPlayers() []*player.Player {
	return g.Players
}

func (g *CardGameTemplate[T]) AddPlayer(players *player.Player) {
	g.Players = append(g.Players, players)
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

	// 詢問 Human 玩家數量
	for {
		fmt.Print("請輸入要幾位 human 玩家: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		count, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("無效的輸入，請輸入數字")
			continue
		}
		if count < 0 {
			fmt.Println("Human 玩家數量不能為負數")
			continue
		}
		if count > requiredPlayers {
			fmt.Printf("Human 玩家數量不能超過 %d 位\n", requiredPlayers)
			continue
		}
		humanCount = count
		break
	}

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
func (g *CardGameTemplate[T]) Start() {
	fmt.Println("--- Starting the game ---")
	g.Game.setup()
	g.Game.setupPlayers()
	g.Game.beforeTakeTurn()
	for !g.Game.isGameOver() {
		g.Game.takeTurn()
	}
	g.Game.afterTakeTurn()
	g.Game.showWinner()
}
