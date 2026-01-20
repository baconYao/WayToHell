package player

import (
	"errors"
	"fmt"

	"cardkit/internal/card"
	"cardkit/internal/hand"
)

// PlayerBehaviorStrategy 定義了玩家如何「做決定」的行為
type PlayerBehaviorStrategy[T card.Card] interface {
	DecideName() string
	DecideCard(handCards []T) int
}

// T 必須實作 card.Card 介面
type BasePlayer[T card.Card] struct {
	name             string
	hand             *hand.Hand[T]
	behaviorStrategy PlayerBehaviorStrategy[T]
}

func NewBasePlayer[T card.Card](behaviorStrategy PlayerBehaviorStrategy[T]) BasePlayer[T] {
	return BasePlayer[T]{
		hand:             &hand.Hand[T]{Cards: make([]T, 0)},
		behaviorStrategy: behaviorStrategy,
	}
}

func (b *BasePlayer[T]) AddHandCard(c T) {
	b.hand.Add(c)
}

func (b *BasePlayer[T]) GetHandCards() []T {
	return b.hand.GetCards()
}

func (b *BasePlayer[T]) RemoveHandCard(index int) {
	err := b.hand.Remove(index)
	if err != nil {
		fmt.Println(err)
	}
}

// PlayCard is a placeholder (abstract method) for the subclass to implement
func (b *BasePlayer[T]) PlayCard() T {
	cards := b.hand.GetCards()
	index := b.behaviorStrategy.DecideCard(cards)
	if index < 0 || index >= len(cards) {
		// 回傳 T 的零值
		var zero T
		return zero
	}
	chosenCard := cards[index]
	b.hand.Remove(index)
	return chosenCard
}

// NameHimSelf is a placeholder (abstract method) for the subclass to implement
func (b *BasePlayer[T]) NameHimSelf() {
	name := b.behaviorStrategy.DecideName()
	for b.SetName(name) != nil {
		// 如果名稱不合法（針對真人），就重新要求再次輸入名稱
		name = b.behaviorStrategy.DecideName()
	}
}

func (b *BasePlayer[T]) SetName(name string) error {
	if len(name) < 3 || len(name) > 5 {
		return errors.New("name must be between 3 and 5 characters")
	}
	b.name = name
	return nil
}

func (b *BasePlayer[T]) GetName() string {
	return b.name
}
