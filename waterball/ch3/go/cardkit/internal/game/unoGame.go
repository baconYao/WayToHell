package game

import (
	"fmt"

	"github.com/baconYao/WayToHell/waterball/ch3/go/cardkit/internal/deck"
	"github.com/baconYao/WayToHell/waterball/ch3/go/cardkit/internal/player"
)

type UnoGame struct {
	CardGameTemplate *CardGameTemplate[bool]
}

func NewUnoGame() *CardGameTemplate[bool] {
	unoGame := &UnoGame{}
	template := &CardGameTemplate[bool]{
		Game:    unoGame,
		Players: make([]*player.Player, 0),
		Deck:    deck.NewUnoDeck(),
	}
	unoGame.CardGameTemplate = template
	return template
}

func (u *UnoGame) setup() {
	u.CardGameTemplate.Deck.Shuffle()
}

func (u *UnoGame) setupPlayers() {
	const requiredPlayers = 4
	aiCount, humanCount := askPlayerCount(requiredPlayers)

	// 驗證總數
	totalPlayers := aiCount + humanCount
	if totalPlayers != requiredPlayers {
		fmt.Printf("玩家總數必須為 %d 位，目前為 %d 位（AI: %d, Human: %d）\n", requiredPlayers, totalPlayers, aiCount, humanCount)
		return
	}

	// 創建 AI 玩家
	for i := 0; i < aiCount; i++ {
		aiPlayer := player.NewAIUnoPlayer()
		aiPlayer.NameHimself()
		u.CardGameTemplate.AddPlayer(&aiPlayer.Player)
	}

	// 創建 Human 玩家
	for i := 0; i < humanCount; i++ {
		humanPlayer := player.NewHumanUnoPlayer()
		humanPlayer.NameHimself()
		u.CardGameTemplate.AddPlayer(&humanPlayer.Player)
	}

	// Dispacth cards to players
	dealCards(u.CardGameTemplate.Players, u.CardGameTemplate.Deck, 5)
}

func (u *UnoGame) beforeTakeTurn() {
	fmt.Println("Uno: 確定當前出牌順序")
}

func (u *UnoGame) takeTurn() {
	fmt.Println("Uno: 玩家嘗試出一張合法的牌")
}

func (u *UnoGame) afterTakeTurn() {
	// 檢查是否有玩家喊 Uno
}

func (u *UnoGame) isGameOver() bool {
	// 檢查是否有玩家手牌為 0
	return false
}

func (u *UnoGame) showWinner() {
	fmt.Println("Uno: 顯示第一位出完牌的贏家")
}

func (u *UnoGame) compare() bool {
	fmt.Println("Uno: 檢查出牌是否與檯面牌顏色或數字相符")
	return true
}
