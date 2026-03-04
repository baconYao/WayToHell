// Package game 負責遊戲流程：初始化（玩家、洗牌發牌、責任鏈）、尋找梅花 3 玩家、
// 回合循環、判定勝負，以及從標準輸入讀取並驅動整個流程。
package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"big2/internal/card"
	cardpatternhandler "big2/internal/cardPatternHandler"
	"big2/internal/deck"
	handcards "big2/internal/handCards"
	"big2/internal/play"
	"big2/internal/player"
	"big2/internal/round"
)

// Game 一局大老二：牌堆、四位玩家、牌型責任鏈、是否首回合、下一回合先出牌者索引。
type Game struct {
	deck         *deck.Deck
	players      []*player.Player
	patternChain cardpatternhandler.Handler
	winner       *player.Player
	isFirstRound bool
	starterIndex int
	scanner      *bufio.Scanner
}

// NewGame 從標準輸入讀取前五行（牌堆 + 四名玩家），建立玩家、牌堆、發牌、責任鏈，
// 並設定首回合由擁有梅花 3 的玩家先出牌；後續動作由 Run 從同一輸入讀取。
func NewGame() (*Game, error) {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 5)
	for i := 0; i < 5 && scanner.Scan(); i++ {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	if len(lines) < 5 {
		return nil, fmt.Errorf("need 5 lines: deck and 4 player names")
	}
	d, err := deck.NewFromShuffledCards(lines[0])
	if err != nil {
		return nil, err
	}
	names := []string{lines[1], lines[2], lines[3], lines[4]}
	players := make([]*player.Player, 4)
	hands := make([]*handcards.HandCards, 4)
	for i := 0; i < 4; i++ {
		players[i] = player.New(i, names[i])
		hands[i] = handcards.New()
		players[i].SetHand(hands[i])
		players[i].SetScanner(scanner)
	}
	for i := 0; !d.Empty(); i++ {
		c, ok := d.Deal()
		if !ok {
			break
		}
		players[i%4].AddHandCard(c)
	}
	g := &Game{
		deck:         d,
		players:      players,
		patternChain: cardpatternhandler.BuildChain(),
		isFirstRound: true,
		scanner:      scanner,
	}
	g.starterIndex = g.findPlayerWithClub3()
	return g, nil
}

// findPlayerWithClub3 回傳手牌中有梅花 3 的玩家索引（0..3）。
func (g *Game) findPlayerWithClub3() int {
	club3 := card.MustParseCard("C[3]")
	for i, p := range g.players {
		for _, c := range p.Hand.GetCards() {
			if c.Compare(club3) == 0 {
				return i
			}
		}
	}
	return 0
}

// Run 執行遊戲主循環：每回合由 starter 先出牌，直到有人手牌出完則宣告贏家。
func (g *Game) Run() {
	starterIndex := g.starterIndex
	for g.winner == nil {
		fmt.Println("新的回合開始了。")
		r := round.New(g.isFirstRound)
		r.Start(starterIndex)
		for !r.IsRoundEnded() {
			idx := r.CurrentPlayerIndex()
			p := g.players[idx]
			playResult := p.Play(g.patternChain, r)
			r.AcceptPlay(playResult)
			if _, ok := playResult.(*play.NormalPlay); ok {
				if p.Hand.IsEmpty() {
					g.winner = p
					break
				}
			}
			r.NextPlayer()
		}
		if g.winner != nil {
			break
		}
		starterIndex = r.RoundWinnerIndex()
		g.isFirstRound = false
	}
	fmt.Printf("遊戲結束，遊戲的勝利者為 %s\n", g.winner.Name)
}
