package game

import (
	"cardkit/internal/deck"
	"cardkit/internal/player"
	"cardkit/pkg/logger"
	"fmt"
)

type ShowdownGame struct {
	// Bidirectional reference to the template, able to access the Players, Deck and its methods...
	CardGameTemplate *CardGameTemplate[struct{}]
	rounds           int
	currentRound     int
}

func NewShowdownGame() *CardGameTemplate[struct{}] {
	showdownGame := &ShowdownGame{
		rounds:       13,
		currentRound: 1,
	}
	template := &CardGameTemplate[struct{}]{
		logger:  logger.GetLogger(),
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

	// 創建 AI 玩家
	for i := 0; i < aiCount; i++ {
		aiPlayer := player.NewAIShowdownPlayer()
		aiPlayer.NameHimself()
		s.CardGameTemplate.AddPlayer(&aiPlayer.Player)
	}

	// 創建 Human 玩家
	for i := 0; i < humanCount; i++ {
		fmt.Printf("請設定第 %d 位 Human 玩家名稱\n", i+1)
		humanPlayer := player.NewHumanShowdownPlayer()
		humanPlayer.NameHimself()
		s.CardGameTemplate.AddPlayer(&humanPlayer.Player)
	}

	s.CardGameTemplate.showPlayers()
	// Dispacth cards to players
	dealCards(s.CardGameTemplate.Players, s.CardGameTemplate.Deck, s.rounds)
}

func (s *ShowdownGame) takeTurn() {
	// TODO: Implement this method
}

func (s *ShowdownGame) isGameOver() bool {
	s.CardGameTemplate.logger.Debug("檢查是否遊戲結束...")
	if s.currentRound > s.rounds {
		return true
	}
	return false
}

func (s *ShowdownGame) showWinner() {
	s.CardGameTemplate.logger.Debug("顯示贏家...")
	players := s.CardGameTemplate.GetPlayers()
	winner := players[0]
	for idx, player := range players {
		fmt.Printf("玩家%d: %s, 得分: %d\n", idx+1, player.GetName(), player.GetPoints())
		if player.GetPoints() > winner.GetPoints() {
			winner = player
		}
	}
	fmt.Printf("贏家是: %s\n", winner.GetName())
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
