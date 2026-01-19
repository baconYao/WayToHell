package player

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"cardkit/internal/card"
	"cardkit/internal/hand"
)

// PlayerStrategy 定義了玩家如何「做決定」的行為
type PlayerBehaviorStrategy interface {
	DecideName() string
	DecideCard(handCards []card.Card) int
}

type AIStrategy struct{}

func (s *AIStrategy) DecideName() string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nameLength := r.Intn(3) + 3
	result := make([]byte, nameLength)
	for i := 0; i < nameLength; i++ {
		result[i] = letters[r.Intn(len(letters))]
	}
	return string(result)
}

func (s *AIStrategy) DecideCard(handCards []card.Card) int {
	if len(handCards) == 0 {
		return -1
	}
	randomIndex := rand.Intn(len(handCards))
	return randomIndex
}

type HumanStrategy struct{}

func (s *HumanStrategy) DecideName() string {
	fmt.Println("請設定該人類玩家名稱 (3 到 5 個字元):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func (s *HumanStrategy) DecideCard(handCards []card.Card) int {
	if len(handCards) == 0 {
		fmt.Println("無牌可出")
		return -1
	}
	fmt.Println("你的手牌如下:")
	for i, c := range handCards {
		fmt.Printf("%d. %s\n", i+1, c.ToString())
	}
	fmt.Print("請選擇出牌編號: ")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		index, err := strconv.Atoi(input)
		if err != nil || index < 1 || index > len(handCards) {
			fmt.Printf("無效輸入，請輸入 1 到 %d: ", len(handCards))
			continue
		}
		return index - 1
	}
}

// T 必須實作 card.Card 介面
type BasePlayer[T card.Card] struct {
	name             string
	hand             *hand.Hand
	behaviorStrategy PlayerBehaviorStrategy
}

func NewBasePlayer[T card.Card](behaviorStrategy PlayerBehaviorStrategy) BasePlayer[T] {
	return BasePlayer[T]{
		hand:             &hand.Hand{Cards: make([]card.Card, 0)},
		behaviorStrategy: behaviorStrategy,
	}
}

func (b *BasePlayer[T]) AddHandCard(c card.Card) {
	b.hand.Add(c)
}

func (b *BasePlayer[T]) GetHandCards() []card.Card {
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
	// 在 New 時保證了傳入的會是 T，將 card.Card 轉回具體的 T
	return chosenCard.(T)
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
