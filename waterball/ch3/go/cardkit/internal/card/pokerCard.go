package card

import "fmt"

type Suit int

const (
	Club    Suit = iota // 0
	Diamond             // 1
	Heart               // 2
	Spade               // 3
)

type Rank int

const (
	Two   Rank = iota + 2 // 2
	Three                 // 3
	Four                  // 4
	Five                  // 5
	Six                   // 6
	Seven                 // 7
	Eight                 // 8
	Nine                  // 9
	Ten                   // 10
	Jack                  // 11
	Queen                 // 12
	King                  // 13
	Ace                   // 14
)

type PokerCard struct {
	Suit Suit
	Rank Rank
}

func (c PokerCard) ToString() string {
	suitName := map[Suit]string{
		Club:    "Club",
		Diamond: "Diamond",
		Heart:   "Heart",
		Spade:   "Spade",
	}
	rankName := map[Rank]string{
		Jack:  "J",
		Queen: "Q",
		King:  "K",
		Ace:   "A",
	}
	rName, ok := rankName[c.Rank]
	if !ok {
		rName = fmt.Sprintf("%d", c.Rank)
	}

	return fmt.Sprintf("[%s %s]", suitName[c.Suit], rName)
}

func (c PokerCard) GetRank() Rank {
	return c.Rank
}

func (c PokerCard) GetSuit() Suit {
	return c.Suit
}
