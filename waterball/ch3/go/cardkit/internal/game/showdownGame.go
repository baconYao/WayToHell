package game

import (
	"fmt"

	"github.com/baconYao/WayToHell/waterball/ch3/go/cardkit/internal/deck"
	"github.com/baconYao/WayToHell/waterball/ch3/go/cardkit/internal/player"
)

type ShowdownGame struct {
	// Bidirectional reference to the template, able to access the Players, Deck and its methods...
	CardGameTemplate *CardGameTemplate[struct{}]
	Rounds           int
}

func NewShowdownGame() *CardGameTemplate[struct{}] {
	showdownGame := &ShowdownGame{
		Rounds: 13,
	}
	template := &CardGameTemplate[struct{}]{
		Game:    showdownGame,
		Players: nil,
		Deck:    deck.NewShowdownDeck(),
	}
	showdownGame.CardGameTemplate = template
	return template
}

func (s *ShowdownGame) setup() {
	s.CardGameTemplate.Deck.Shuffle()
}

func (s *ShowdownGame) setupPlayers() {
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
		aiPlayer := player.NewAIShowdownPlayer()
		aiPlayer.NameHimself()
		s.CardGameTemplate.AddPlayer(&aiPlayer.Player)
	}

	// 創建 Human 玩家
	for i := 0; i < humanCount; i++ {
		humanPlayer := player.NewHumanShowdownPlayer()
		humanPlayer.NameHimself()
		s.CardGameTemplate.AddPlayer(&humanPlayer.Player)
	}

	// Dispacth cards to players
	dealCards(s.CardGameTemplate.Players, s.CardGameTemplate.Deck, s.Rounds)
}

func (s *ShowdownGame) takeTurn() {
	// TODO: Implement this method
}

func (s *ShowdownGame) isGameOver() bool {
	return false
}

func (s *ShowdownGame) showWinner() {
	// TODO: Implement this method
}

func (s *ShowdownGame) compare() struct{} {
	return struct{}{}
}

// BeforeTakeTurn is a placeholder since we don't need to implement anything here
func (s *ShowdownGame) beforeTakeTurn() {
}

// afterTakeTurn is a placeholder since we don't need to implement anything here
func (s *ShowdownGame) afterTakeTurn() {
}
