// Package player 定義玩家 Player：名稱、手牌、出牌決策（含選牌與 PASS）、顯示手牌、與責任鏈驗證牌型。
package player

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"big2/internal/card"
	cardpatternhandler "big2/internal/cardPatternHandler"
	handcards "big2/internal/handCards"
	"big2/internal/play"
	"big2/internal/round"
)

// Player 玩家：索引、名稱、手牌。
type Player struct {
	Index int
	Name  string
	Hand  *handcards.HandCards
}

// New 建立玩家。index 為 0..3。手牌由呼叫方以 SetHand 設定。
func New(index int, name string) *Player {
	return &Player{Index: index, Name: name, Hand: nil}
}

// SetHand 設定手牌（由 game 在發牌前呼叫）。
func (p *Player) SetHand(h *handcards.HandCards) {
	p.Hand = h
}

// AddHandCard 加入一張手牌。
func (p *Player) AddHandCard(c card.Card) {
	if p.Hand != nil {
		p.Hand.AddCard(c)
	}
}

// ShowCards 輸出「輪到<名字>了」與手牌索引、牌面兩行；索引對齊下方牌面寬度，牌面以空白分隔。
func (p *Player) ShowCards() {
	fmt.Printf("輪到%s了\n", p.Name)
	p.showHandOnly()
}

// showHandOnly 只輸出手牌索引與牌面兩行（不輸出「輪到XX了」），用於同一玩家重試時。
func (p *Player) showHandOnly() {
	if p.Hand == nil || p.Hand.Len() == 0 {
		return
	}
	cards := p.Hand.GetCards()
	for i := 0; i < len(cards); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		w := len(cards[i].String())
		if i < len(cards)-1 {
			fmt.Printf("%-*d", w, i)
		} else {
			fmt.Print(i)
		}
	}
	fmt.Println()
	for i := 0; i < len(cards); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(cards[i].String())
	}
	fmt.Println()
}

// Play 讀取一行輸入，驗證並回傳合法出牌；不合法則輸出錯誤並重試。僅在該玩家輪到時第一次顯示「輪到XX了」，重試時只顯示手牌兩行。
func (p *Player) Play(chain cardpatternhandler.Handler, r *round.Round, readLine func() string) play.Play {
	first := true
	for {
		if first {
			p.ShowCards()
			first = false
		} else {
			p.showHandOnly()
		}
		line := strings.TrimSpace(readLine())
		if line == "-1" || line == "1" {
			if r.GetTopPlay() == nil {
				fmt.Println("你不能在新的回合中喊 PASS")
				continue
			}
			fmt.Printf("玩家 %s PASS.\n", p.Name)
			return &play.PassPlay{PlayerIndex: p.Index}
		}
		indices, err := parseIndices(line)
		if err != nil {
			fmt.Println("此牌型不合法，請再嘗試一次。")
			continue
		}
		cards := make([]card.Card, 0, len(indices))
		for _, i := range indices {
			c, ok := p.Hand.CardAt(i)
			if !ok {
				fmt.Println("此牌型不合法，請再嘗試一次。")
				cards = nil
				break
			}
			cards = append(cards, c)
		}
		if len(cards) == 0 {
			continue
		}
		h := chain.Validate(cards)
		if h == nil {
			fmt.Println("此牌型不合法，請再嘗試一次。")
			continue
		}
		if !r.CheckFirstMoveRule(cards) {
			fmt.Println("首局第一手必須包含梅花 3")
			continue
		}
		norm := play.NewNormalPlay(p.Index, h)
		if !norm.IsStrongerThan(r.GetTopPlay()) {
			fmt.Println("此牌型不合法，請再嘗試一次。")
			// fmt.Println("牌不夠大，請重新出牌")
			continue
		}
		fmt.Printf("玩家 %s 打出了 %s %s\n", p.Name, norm.PatternName, formatCards(norm.GetCards()))
		p.Hand.RemoveCards(norm.GetCards())
		return norm
	}
}

func formatCards(cards []card.Card) string {
	var b strings.Builder
	for i, c := range cards {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(c.String())
	}
	return b.String()
}

func parseIndices(line string) ([]int, error) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return nil, strconv.ErrSyntax
	}
	out := make([]int, 0, len(parts))
	for _, s := range parts {
		i, err := strconv.Atoi(s)
		if err != nil || i < 0 {
			return nil, err
		}
		out = append(out, i)
	}
	return out, nil
}

// ReadLineFromScanner 供 game 傳入：從 scanner 讀一行。
func ReadLineFromScanner(scanner *bufio.Scanner) func() string {
	return func() string {
		if scanner.Scan() {
			return scanner.Text()
		}
		return ""
	}
}
