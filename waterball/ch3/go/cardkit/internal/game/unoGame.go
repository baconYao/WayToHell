package game

import (
	"fmt"

	"cardkit/internal/deck"
	"cardkit/internal/player"
	"cardkit/pkg/logger"
)

type UnoGame struct {
	CardGameTemplate *CardGameTemplate[bool]
	unoPlayers       []*player.UnoPlayer
}

func NewUnoGame() *CardGameTemplate[bool] {
	unoGame := &UnoGame{
		unoPlayers: make([]*player.UnoPlayer, 0),
	}
	template := &CardGameTemplate[bool]{
		logger: logger.GetLogger(),
		Game:   unoGame,
		Deck:   deck.NewUnoDeck(),
	}
	unoGame.CardGameTemplate = template
	return template
}

func (u *UnoGame) setup() {
	u.CardGameTemplate.Deck.Shuffle()
}

func (u *UnoGame) AddUnoPlayer(player *player.UnoPlayer) {
	u.unoPlayers = append(u.unoPlayers, player)
}

func (u *UnoGame) setupPlayers() {
	const requiredPlayers = 4
	aiCount, humanCount := askPlayerCount(requiredPlayers)

	// 創建 AI 玩家
	for i := 0; i < aiCount; i++ {
		aiPlayer := player.NewUnoPlayer(&player.AIStrategy{})
		aiPlayer.NameHimSelf()
		u.AddUnoPlayer(aiPlayer)
	}

	// 創建 Human 玩家
	for i := 0; i < humanCount; i++ {
		fmt.Printf("請設定第 %d 位 Human 玩家名稱\n", i+1)
		humanPlayer := player.NewUnoPlayer(&player.HumanStrategy{})
		humanPlayer.NameHimSelf()
		u.AddUnoPlayer(humanPlayer)
	}

	// u.CardGameTemplate.showPlayers()
	// Dispacth cards to players
	// dealCards(u.CardGameTemplate.Players, u.CardGameTemplate.Deck, 5)
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
