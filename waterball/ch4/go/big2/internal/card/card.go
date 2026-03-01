// Package card 定義撲克牌領域概念：Card（牌）、Rank（數字）、Suit（花色），以及牌的比較邏輯。
package card

import (
	"fmt"
	"strings"
)

// Suit 花色。大小順序：梅花 < 菱形 < 愛心 < 黑桃。
type Suit int

const (
	Club    Suit = iota // C 梅花
	Diamond             // D 菱形
	Heart               // H 愛心
	Spade               // S 黑桃
)

var suitRunes = map[Suit]string{Club: "C", Diamond: "D", Heart: "H", Spade: "S"}

func (s Suit) String() string { return suitRunes[s] }

// ParseSuit 從單字元解析花色（C, D, H, S）。
func ParseSuit(r rune) (Suit, bool) {
	switch r {
	case 'C':
		return Club, true
	case 'D':
		return Diamond, true
	case 'H':
		return Heart, true
	case 'S':
		return Spade, true
	default:
		return 0, false
	}
}

// Rank 數字。大小順序：3 < 4 < ... < 10 < J < Q < K < A < 2。
type Rank int

const (
	Rank3 Rank = iota
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
	Rank9
	Rank10
	RankJ
	RankQ
	RankK
	RankA
	Rank2
)

var rankStrings = map[Rank]string{
	Rank3: "3", Rank4: "4", Rank5: "5", Rank6: "6", Rank7: "7", Rank8: "8", Rank9: "9",
	Rank10: "10", RankJ: "J", RankQ: "Q", RankK: "K", RankA: "A", Rank2: "2",
}

func (r Rank) String() string { return rankStrings[r] }

// ParseRank 從字串解析數字（3,4,...,9,10,J,Q,K,A,2）。
func ParseRank(s string) (Rank, bool) {
	switch s {
	case "3":
		return Rank3, true
	case "4":
		return Rank4, true
	case "5":
		return Rank5, true
	case "6":
		return Rank6, true
	case "7":
		return Rank7, true
	case "8":
		return Rank8, true
	case "9":
		return Rank9, true
	case "10":
		return Rank10, true
	case "J":
		return RankJ, true
	case "Q":
		return RankQ, true
	case "K":
		return RankK, true
	case "A":
		return RankA, true
	case "2":
		return Rank2, true
	default:
		return 0, false
	}
}

// Card 撲克牌，(數字, 花色) 的組合。
type Card struct {
	Rank Rank
	Suit Suit
}

// Compare 與另一張牌比較。先比數字再比花色。
// 回傳：負數表示 this < other，0 表示相等，正數表示 this > other。
func (c Card) Compare(other Card) int {
	if d := int(c.Rank) - int(other.Rank); d != 0 {
		return d
	}
	return int(c.Suit) - int(other.Suit)
}

func (c Card) String() string {
	return fmt.Sprintf("%s[%s]", c.Suit.String(), c.Rank.String())
}

// ParseCard 從規格格式解析一張牌，例如 "C[3]"、"S[10]"、"H[A]"。
func ParseCard(s string) (Card, error) {
	s = strings.TrimSpace(s)
	if len(s) < 4 {
		return Card{}, fmt.Errorf("invalid card format: %q", s)
	}
	runes := []rune(s)
	if runes[1] != '[' || runes[len(runes)-1] != ']' {
		return Card{}, fmt.Errorf("invalid card format: %q", s)
	}
	suit, ok := ParseSuit(runes[0])
	if !ok {
		return Card{}, fmt.Errorf("invalid suit: %q", string(runes[0]))
	}
	rankStr := string(runes[2 : len(runes)-1])
	rank, ok := ParseRank(rankStr)
	if !ok {
		return Card{}, fmt.Errorf("invalid rank: %q", rankStr)
	}
	return Card{Rank: rank, Suit: suit}, nil
}

// ParseCards 解析多張牌，以空白分隔，例如 "S[8] S[9] C[3]"。
func ParseCards(line string) ([]Card, error) {
	parts := strings.Fields(line)
	cards := make([]Card, 0, len(parts))
	for _, p := range parts {
		c, err := ParseCard(p)
		if err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}
	return cards, nil
}

// MustParseCard 解析單張牌，解析失敗時 panic（用於測試或確定格式正確時）。
func MustParseCard(s string) Card {
	c, err := ParseCard(s)
	if err != nil {
		panic(err)
	}
	return c
}
