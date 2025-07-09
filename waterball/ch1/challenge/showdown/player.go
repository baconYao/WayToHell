package showdown

import (
	"errors"
	"showdown/logger"
)

const MAXCARDS int = ROUNDS

type Player interface {
	Show() (Card, error)
	GainCard(card Card) error
	GetName() []byte
	GainPoint(point int) error
	GetPoint() int
	GetCards() []Card
	SetCards(cards []Card) error
	ExchangeHands(switchBackIteration int) error
	SetExchangedHand(eh *ExchangedHand) error
	GetExchangedHand() *ExchangedHand
	GetShowdown() *Showdown
	SetShowdown(*Showdown) error
	WantExchangeHands() (bool, error)
	NameHimSelf() error
	selectExchangeCandidate() (Player, error)
}

type BasePlayer struct {
	showdown      *Showdown
	name          []byte
	point         int
	cards         []Card
	exchangedHand *ExchangedHand
}

// GainCard adds a card into players' hand
func (bp *BasePlayer) GainCard(card Card) error {
	log := logger.GetLogger()
	if len(bp.cards) > MAXCARDS {
		return errors.New("cannot gain cards: maximum card limit reached")
	}
	bp.cards = append(bp.cards, card)
	log.Debug("Player %s gains a new card: %s-%s\n", bp.GetName(), card.GetSuit(), card.GetRank())
	return nil
}

func (bp BasePlayer) GetPoint() int {
	return bp.point
}

func (bp *BasePlayer) GainPoint(point int) error {
	if point < 1 {
		return errors.New("cannot gain the negative point")
	}
	bp.point = bp.point + point
	return nil
}

func (bp BasePlayer) GetShowdown() *Showdown {
	return bp.showdown
}

func (bp *BasePlayer) SetShowdown(showdown *Showdown) error {
	if showdown == nil {
		return errors.New("player cannot join a non-existent game")
	}
	bp.showdown = showdown
	return nil
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

func (bp BasePlayer) GetCards() []Card {
	return append([]Card(nil), bp.cards...)
}

func (bp *BasePlayer) SetCards(cards []Card) error {
	bp.cards = append([]Card(nil), cards...)
	return nil
}

func (bp *BasePlayer) SetExchangedHand(eh *ExchangedHand) error {
	if eh == nil {
		return errors.New("cannot set to exchangedHand as nil")
	}
	bp.exchangedHand = eh
	return nil
}

func (bp BasePlayer) GetExchangedHand() *ExchangedHand {
	return bp.exchangedHand
}
