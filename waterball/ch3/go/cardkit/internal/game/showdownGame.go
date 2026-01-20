package game

import (
	"cardkit/internal/card"
	"cardkit/internal/deck"
	"cardkit/internal/player"
	"cardkit/pkg/logger"
	"fmt"
)

type ShowdownGame struct {
	// Bidirectional reference to the template, able to access the Players, Deck and its methods...
	CardGameTemplate *CardGameTemplate
	currentRound     int
	deck             *deck.Deck[card.PokerCard]
	requiredPlayers  int
	rounds           int
	showdownPlayers  []*player.ShowdownPlayer
	tableCards       []card.PokerCard // 桌上牌堆, 用來保存每回合打出來的牌
}

func NewShowdownGame() *CardGameTemplate {
	showdownGame := &ShowdownGame{
		currentRound:    1,
		deck:            deck.NewShowdownDeck(),
		requiredPlayers: 4,
		rounds:          13,
		showdownPlayers: make([]*player.ShowdownPlayer, 0),
		tableCards:      make([]card.PokerCard, 0),
	}
	template := &CardGameTemplate{
		logger: logger.GetLogger(),
		Game:   showdownGame,
	}
	showdownGame.CardGameTemplate = template
	return template
}

func (s *ShowdownGame) addShowdownPlayer(player *player.ShowdownPlayer) {
	s.showdownPlayers = append(s.showdownPlayers, player)
}

func (s *ShowdownGame) addTableCard(card card.PokerCard) {
	s.tableCards = append(s.tableCards, card)
}

func (s *ShowdownGame) getTableCards() []card.PokerCard {
	replica := make([]card.PokerCard, len(s.tableCards))
	copy(replica, s.tableCards[:len(s.tableCards)-1])
	return replica
}

func (s *ShowdownGame) clearTableCards() {
	s.tableCards = make([]card.PokerCard, 0)
}

func (s *ShowdownGame) setup() {
	aiCount, humanCount := askPlayerCount(s.requiredPlayers)

	// 創建 AI 玩家
	for i := 0; i < aiCount; i++ {
		aiPlayer := player.NewShowdownPlayer(&player.AIStrategy[card.PokerCard]{})
		aiPlayer.NameHimSelf()
		s.addShowdownPlayer(aiPlayer)
	}

	// 創建 Human 玩家
	for i := 0; i < humanCount; i++ {
		humanPlayer := player.NewShowdownPlayer(&player.HumanStrategy[card.PokerCard]{})
		humanPlayer.NameHimSelf()
		s.addShowdownPlayer(humanPlayer)
	}

	s.deck.Shuffle()

	fmt.Printf("為 %d 位玩家各發放 %d 張起始手牌\n", len(s.showdownPlayers), s.rounds)
	for i := 0; i < s.rounds; i++ {
		for _, sPlayer := range s.showdownPlayers {
			sPlayer.AddHandCard(s.deck.Draw())
		}
	}

}

func (s *ShowdownGame) takeTurn() {
	for _, sPlayer := range s.showdownPlayers {
		card := sPlayer.PlayCard()
		s.addTableCard(card)
		fmt.Printf("玩家 %s 出一張牌 %s\n", sPlayer.GetName(), card.ToString())
	}
	s.compare()
	s.clearTableCards()
	s.currentRound++
}

func (s *ShowdownGame) isGameOver() bool {
	s.CardGameTemplate.logger.Debug("檢查是否遊戲結束...")
	if s.currentRound > s.rounds {
		fmt.Printf("\n已經完成 %d 輪遊戲，遊戲結束\n", s.rounds)
		return true
	}
	fmt.Printf("\n進行第 %d 輪遊戲...\n", s.currentRound)
	return false
}

func (s *ShowdownGame) showWinner() {
	s.CardGameTemplate.logger.Debug("顯示贏家...")
	winner := s.showdownPlayers[0]
	for idx, player := range s.showdownPlayers {
		fmt.Printf("玩家%d: %s, 得分: %d\n", idx+1, player.GetName(), player.GetPoints())
		if player.GetPoints() > winner.GetPoints() {
			winner = player
		}
	}
	fmt.Printf("贏家是: %s\n", winner.GetName())
}

// compare handles the comparison logic of showdowm game
func (s *ShowdownGame) compare() {
	cards := s.getTableCards()
	maxCardIndex := 0
	for idx, card := range cards {
		if card.GetRank() > cards[maxCardIndex].GetRank() {
			maxCardIndex = idx
		} else if card.GetRank() == cards[maxCardIndex].GetRank() {
			if card.GetSuit() > cards[maxCardIndex].GetSuit() {
				maxCardIndex = idx
			}
		}
	}

	s.showdownPlayers[maxCardIndex].GainPoint()
	fmt.Printf("玩家 %s 獲得一分，目前得分: %d\n", s.showdownPlayers[maxCardIndex].GetName(), s.showdownPlayers[maxCardIndex].GetPoints())
}

// // BeforeTakeTurn is a placeholder since we don't need to implement anything here
// func (s *ShowdownGame) beforeTakeTurn() {
// 	s.CardGameTemplate.logger.Debug("before take turn...")
// }

// // afterTakeTurn is a placeholder since we don't need to implement anything here
// func (s *ShowdownGame) afterTakeTurn() {
// 	s.CardGameTemplate.logger.Debug("after take turn...")
// 	s.clearTableCards()
// 	s.currentRound++
// }
