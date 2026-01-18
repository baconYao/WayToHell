package game

import (
	"bufio"
	"cardkit/internal/deck"
	"cardkit/internal/player"
	"cardkit/pkg/logger"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	SHOWDOWN_GAME = "showdown"
	UNO_GAME      = "uno"
)

type CardGame struct {
	game   CardGameStrategy
	logger *logger.Logger
}

func NewCardGame(strategy CardGameStrategy) *CardGame {
	return &CardGame{
		game:   strategy,
		logger: logger.GetLogger(),
	}
}

func (c *CardGame) Start() {
	c.logger.Info("Starting the game...")
	c.game.Start()
}

// dispatchCards deals cards to all players
// players: list of players to deal cards to
// deck: the deck to draw cards from
// rounds: number of rounds to deal (each round gives one card to each player)
func dispatchCards(players []*player.BasePlayer, deck *deck.Deck, rounds int) {
	fmt.Printf("為 %d 位玩家各發放 %d 張起始手牌\n", len(players), rounds)
	for i := 0; i < rounds; i++ {
		for _, player := range players {
			player.AddHandCard(deck.Draw())
		}
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
