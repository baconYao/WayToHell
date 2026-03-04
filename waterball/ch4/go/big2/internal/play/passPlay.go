package play

import "big2/internal/card"

// PassPlay 放棄出牌。
type PassPlay struct {
	PlayerIndex int
}

func (p *PassPlay) GetPlayerIndex() int { return p.PlayerIndex }

func (p *PassPlay) GetCards() []card.Card { return nil }

// IsStrongerThan 放棄出牌的牌型比任何牌型都弱。
func (p *PassPlay) IsStrongerThan(other Play) bool {
	return false
}
