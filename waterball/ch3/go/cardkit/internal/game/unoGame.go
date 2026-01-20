package game

import (
	"fmt"

	"cardkit/internal/card"
	"cardkit/internal/deck"
	"cardkit/internal/player"
	"cardkit/pkg/logger"
)

type UnoGame struct {
	CardGameTemplate *CardGameTemplate
	deck             *deck.Deck[card.UnoCard]
	hasValidCard     bool // 當前玩家是否有合法的牌
	requiredPlayers  int
	tableCards       []card.UnoCard // 桌上牌堆, 用來保存打出來的牌, 最後一張牌為當前出牌
	unoPlayers       []*player.UnoPlayer
	winnerIndex      int
}

func NewUnoGame() *CardGameTemplate {
	unoGame := &UnoGame{
		deck:            deck.NewUnoDeck(),
		hasValidCard:    false,
		requiredPlayers: 4,
		tableCards:      make([]card.UnoCard, 0),
		unoPlayers:      make([]*player.UnoPlayer, 0),
		winnerIndex:     -1,
	}
	template := &CardGameTemplate{
		logger: logger.GetLogger(),
		Game:   unoGame,
	}
	unoGame.CardGameTemplate = template
	return template
}

// addTableCard adds a card to the table
func (u *UnoGame) addTableCard(card card.UnoCard) {
	if len(u.tableCards) > 37 {
		panic("桌上牌堆數量超過 37 張")
	}
	u.tableCards = append(u.tableCards, card)
}

func (u *UnoGame) getTableCards() []card.UnoCard {
	replica := make([]card.UnoCard, len(u.tableCards))
	copy(replica, u.tableCards[:len(u.tableCards)-1])
	return replica
}

// clearTableCards discards all cards except the latest one
func (u *UnoGame) clearTableCards() []card.UnoCard {
	cards := u.getTableCards()
	discardedCards := cards[:len(cards)-1]
	u.tableCards = u.tableCards[len(u.tableCards)-1:]
	return discardedCards
}

// getLatestTableCard returns the latest table card
func (u *UnoGame) getLatestTableCard() card.UnoCard {
	if len(u.tableCards) == 0 {
		panic("桌上沒有牌")
	}
	return u.tableCards[len(u.tableCards)-1]
}

func (u *UnoGame) setup() {
	aiCount, humanCount := askPlayerCount(u.requiredPlayers)

	// 創建 AI 玩家
	for i := 0; i < aiCount; i++ {
		aiPlayer := player.NewUnoPlayer(&player.AIStrategy[card.UnoCard]{})
		aiPlayer.NameHimSelf()
		u.AddUnoPlayer(aiPlayer)
	}

	// 創建 Human 玩家
	for i := 0; i < humanCount; i++ {
		fmt.Printf("請設定第 %d 位 Human 玩家名稱\n", i+1)
		humanPlayer := player.NewUnoPlayer(&player.HumanStrategy[card.UnoCard]{})
		humanPlayer.NameHimSelf()
		u.AddUnoPlayer(humanPlayer)
	}

	u.deck.Shuffle()

	fmt.Printf("為 %d 位玩家各發放 %d 張起始手牌\n", len(u.unoPlayers), 5)
	for i := 0; i < 5; i++ {
		for _, player := range u.unoPlayers {
			player.AddHandCard(u.deck.Draw())
		}
	}

	// 從牌堆中抽一張，放到牌桌上，作為起始牌
	startingCard := u.deck.Draw()
	u.addTableCard(startingCard)
}

func (u *UnoGame) AddUnoPlayer(player *player.UnoPlayer) {
	u.unoPlayers = append(u.unoPlayers, player)
}

// func (u *UnoGame) beforeTakeTurn() {
// 	// nothing to do here
// }

func (u *UnoGame) takeTurn() {
	for idx, player := range u.unoPlayers {
		fmt.Printf("\n當前牌桌上的牌 --> %s\n", u.getLatestTableCard().ToString())
		fmt.Printf("輪到玩家 %s 出牌, 手牌數量: %d\n", player.GetName(), len(player.GetHandCards()))
		hasValidCard := false
		// 檢查當前玩家是否有合法的牌
		for _, card := range player.GetHandCards() {
			if u.isCardValid(card) {
				hasValidCard = true
				break
			}
		}
		// 沒有牌可以出
		if !hasValidCard {
			fmt.Println("玩家沒牌出，從牌堆中抽一張牌...")
			if u.deck.IsEmpty() {
				fmt.Println("牌堆空了，從牌桌上回收牌...")
				discardedCards := u.clearTableCards()
				u.deck.Refill(discardedCards)
			}
			player.AddHandCard(u.deck.Draw())
			continue // 換下一位玩家出牌
		}
		for {
			card := player.PlayCard()
			if !u.isCardValid(card) {
				fmt.Printf("嘗試打出 %s 不合法的牌，請重新出牌...\n", card.ToString())
				player.AddHandCard(card)
				continue
			}
			fmt.Printf("玩家 %s 出的牌: %s\n", player.GetName(), card.ToString())
			u.addTableCard(card)

			if len(player.GetHandCards()) == 0 {
				u.winnerIndex = idx
				return
			}
			break // 出牌成功，換下一位玩家出牌
		}
	}
}

// func (u *UnoGame) afterTakeTurn() {
// 	// nothing to do here
// }

func (u *UnoGame) isGameOver() bool {
	// 檢查是否有玩家手牌為 0
	if u.winnerIndex == -1 {
		return false
	}
	return len(u.unoPlayers[u.winnerIndex].GetHandCards()) == 0
}

func (u *UnoGame) showWinner() {
	fmt.Printf("\n玩家 %s 贏得遊戲\n", u.unoPlayers[u.winnerIndex].GetName())
}

// isCardValid checks whether the card is valid compare to the latest card on the table
func (u *UnoGame) isCardValid(card card.UnoCard) bool {
	latestTableCard := u.getLatestTableCard()
	return card.Color == latestTableCard.Color || card.Number == latestTableCard.Number
}
