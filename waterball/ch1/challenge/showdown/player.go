package showdown

import (
	"errors"
	"fmt"
)

const MAXCARDS int = ROUNDS

type Player interface {
	Show() (Card, error)
	GainCard(card Card) error
	GetName() []byte
	GainPoint(point int) error
	GetPoint() int
	GetPrivilege() bool
	SetPrivilege(p bool) error
	GetExchangedHand() *ExchangedHand
	NameHimSelf() error
}

type BasePlayer struct {
	name          []byte
	point         int
	cards         []Card
	privilege     bool // the ability to perform exchange hand privilege
	exchangedHand *ExchangedHand
}

func (bp *BasePlayer) Show() (Card, error) {
	panic("NameHimSelf method must be implemented by concrete types")
}

// GainCard adds a card into players' hand
func (bp *BasePlayer) GainCard(card Card) error {
	if len(bp.cards) > MAXCARDS {
		return errors.New("cannot gain cards: maximum card limit reached")
	}
	bp.cards = append(bp.cards, card)
	fmt.Printf("Player %s gains a new card: %s-%s\n", bp.GetName(), card.GetSuit(), card.GetRank())
	return nil
}

func (bp BasePlayer) GetPoint() int {
	return bp.point
}

func (bp *BasePlayer) GainPoint(point int) error {
	if point < 1 {
		return errors.New("cannot gain the negative point")
	}
	bp.point = point
	return nil
}

func (bp BasePlayer) NameHimSelf() error {
	panic("NameHimSelf method must be implemented by concrete types")
}

func (bp BasePlayer) GetName() []byte {
	return bp.name
}

func (bp *BasePlayer) setName(name []byte) error {
	if len(name) < 4 || len(name) > 8 {
		return errors.New("name length must be between 4 and 8 characters")
	}
	bp.name = name
	return nil
}

func (bp BasePlayer) GetPrivilege() bool {
	return bp.privilege
}

func (bp *BasePlayer) SetPrivilege(p bool) error {
	bp.privilege = p
	return nil
}

func exchangedHands(candidate BasePlayer, switchBackIteration int) (*ExchangedHand, error) {
	return nil, nil
}

func (bp BasePlayer) GetExchangedHand() *ExchangedHand {
	return bp.exchangedHand
}
