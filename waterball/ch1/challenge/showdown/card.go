package showdown

type Card struct {
	rank Rank
	suit Suit
}

// Rank 表示撲克牌的階級
type Rank int

const (
	Rank2 Rank = iota
	Rank3
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
)

func (r Rank) String() string {
	switch r {
	case Rank2:
		return "2"
	case Rank3:
		return "3"
	case Rank4:
		return "4"
	case Rank5:
		return "5"
	case Rank6:
		return "6"
	case Rank7:
		return "7"
	case Rank8:
		return "8"
	case Rank9:
		return "9"
	case Rank10:
		return "10"
	case RankJ:
		return "J"
	case RankQ:
		return "Q"
	case RankK:
		return "K"
	case RankA:
		return "A"
	default:
		return "Unknown"
	}
}

// Suit 表示撲克牌的花色
type Suit int

const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)

func (s Suit) String() string {
	switch s {
	case Club:
		return "Club"
	case Diamond:
		return "Diamond"
	case Heart:
		return "Heart"
	case Spade:
		return "Spade"
	default:
		return "Unknown"
	}
}

// GetRank returns the rank of the card
func (c Card) GetRank() Rank {
	return c.rank
}

// GetSuit returns the suit of the card
func (c Card) GetSuit() Suit {
	return c.suit
}

func (c Card) IsValid() bool {
	return c.rank >= Rank2 && c.rank <= RankA && c.suit >= Club && c.suit <= Spade
}
