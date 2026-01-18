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

type BasePlayer struct {
	name             string
	hand             *hand.Hand
	behaviorStrategy PlayerBehaviorStrategy
}

func (b *BasePlayer) AddHandCard(c card.Card) {
	b.hand.Add(c)
}

func (b *BasePlayer) GetHandCards() []card.Card {
	return b.hand.GetCards()
}

func (b *BasePlayer) RemoveHandCard(index int) {
	err := b.hand.Remove(index)
	if err != nil {
		fmt.Println(err)
	}
}

// PlayCard is a placeholder (abstract method) for the subclass to implement
func (b *BasePlayer) PlayCard() card.Card {
	cards := b.hand.GetCards()
	index := b.behaviorStrategy.DecideCard(cards)
	if index < 0 || index >= len(cards) {
		return nil
	}
	card := cards[index]
	b.hand.Remove(index)
	return card
}

// NameHimSelf is a placeholder (abstract method) for the subclass to implement
func (b *BasePlayer) NameHimSelf() {
	name := b.behaviorStrategy.DecideName()
	for b.SetName(name) != nil {
		// 如果名稱不合法（針對真人），就重新要求再次輸入名稱
		name = b.behaviorStrategy.DecideName()
	}
}

func (b *BasePlayer) SetName(name string) error {
	if len(name) < 3 || len(name) > 5 {
		return errors.New("name must be between 3 and 5 characters")
	}
	b.name = name
	return nil
}

func (b *BasePlayer) GetName() string {
	return b.name
}

// AskPlayerCount asks and returns the number of AI 和 Human Players
// requiredPlayers: total players required is this game
// Returns: aiCount, humanCount
func AskPlayerCount(requiredPlayers int) (aiCount, humanCount int) {
	fmt.Printf("此遊戲需要 %d 位玩家。\n", requiredPlayers)

	scanner := bufio.NewScanner(os.Stdin)

	// 詢問 AI 玩家數量
	for {
		fmt.Print("請輸入要幾位 AI 玩家: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		count, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("無效的輸入，請輸入數字")
			continue
		}
		if count < 0 {
			fmt.Println("AI 玩家數量不能為負數")
			continue
		}
		if count > requiredPlayers {
			fmt.Printf("AI 玩家數量不能超過 %d 位\n", requiredPlayers)
			continue
		}
		aiCount = count
		break
	}

	humanCount = requiredPlayers - aiCount

	fmt.Printf("AI 玩家數量: %d, Human 玩家數量: %d\n", aiCount, humanCount)

	return aiCount, humanCount
}
