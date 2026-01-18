package game

import (
	"cardkit/internal/deck"
	"cardkit/internal/player"
	"cardkit/pkg/logger"
)

type ShowdownGame struct {
	// Bidirectional reference to the template, able to access the Players, Deck and its methods...
	CardGameTemplate *CardGameTemplate[struct{}]
	rounds           int
	currentRound     int
	showdownPlayers  []*player.ShowdownPlayer
}

func NewShowdownGame() *CardGameTemplate[struct{}] {
	showdownGame := &ShowdownGame{
		rounds:          13,
		currentRound:    1,
		showdownPlayers: make([]*player.ShowdownPlayer, 0),
	}
	template := &CardGameTemplate[struct{}]{
		logger: logger.GetLogger(),
		Game:   showdownGame,
		Deck:   deck.NewShowdownDeck(),
	}
	showdownGame.CardGameTemplate = template
	return template
}

func (s *ShowdownGame) AddShowdownPlayer(player *player.ShowdownPlayer) {
	s.showdownPlayers = append(s.showdownPlayers, player)
}

func (s *ShowdownGame) setup() {
	const requiredPlayers = 4
	aiCount, humanCount := askPlayerCount(requiredPlayers)

	// 創建 AI 玩家
	for i := 0; i < aiCount; i++ {
		aiPlayer := player.NewShowdownPlayer(&player.AIStrategy{})
		aiPlayer.NameHimSelf()
		s.AddShowdownPlayer(aiPlayer)
	}

	// 創建 Human 玩家
	for i := 0; i < humanCount; i++ {
		humanPlayer := player.NewShowdownPlayer(&player.HumanStrategy{})
		humanPlayer.NameHimSelf()
		s.AddShowdownPlayer(humanPlayer)
	}

	// s.CardGameTemplate.showPlayers()
	// // Dispacth cards to players
	// dealCards(s.CardGameTemplate.Players, s.CardGameTemplate.Deck, s.rounds)

	s.CardGameTemplate.Deck.Shuffle()
}

func (s *ShowdownGame) setupPlayers() {

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
	// s.CardGameTemplate.logger.Debug("顯示贏家...")
	// players := s.CardGameTemplate.GetPlayers()
	// winner := players[0]
	// // FIXME: 需要修正為使用 showdownPlayers 而不是 players
	// for idx, player := range players {
	// 	fmt.Printf("玩家%d: %s, 得分: %d\n", idx+1, player.GetName(), player.GetPoints())
	// 	if player.GetPoints() > winner.GetPoints() {
	// 		winner = player
	// 	}
	// }
	// fmt.Printf("贏家是: %s\n", winner.GetName())
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

func (s *ShowdownGame) addShowdownPlayer(player *player.ShowdownPlayer) {
	s.showdownPlayers = append(s.showdownPlayers, player)
}
