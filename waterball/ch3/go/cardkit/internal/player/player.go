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

type Player struct {
	name string
	hand *hand.Hand
}

func (p *Player) AddHandCard(c card.Card) {
	p.hand.Add(c)
}

func (p *Player) GetHandCards() []card.Card {
	return p.hand.GetCards()
}

func (p *Player) RemoveHandCard(index int) {
	err := p.hand.Remove(index)
	if err != nil {
		fmt.Println(err)
	}
}

// PlayCard is a placeholder (abstract method) for the subclass to implement
func (p *Player) PlayCard() card.Card {
	// Implement this method in the subclass to play a card
	return nil
}

// NameHimself is a placeholder (abstract method) for the subclass to implement
func (p *Player) NameHimself() {
	// Implement this method in the subclass to name itself
}

func (p *Player) SetName(name string) error {
	if len(name) < 3 || len(name) > 5 {
		return errors.New("name must be between 3 and 5 characters")
	}
	p.name = name
	return nil
}

func (p *Player) GetName() string {
	return p.name
}

// AISelectName is a method to select a random name for the AI player
func AISelectName() string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nameLength := r.Intn(3) + 3
	result := make([]byte, nameLength)
	for i := 0; i < nameLength; i++ {
		result[i] = letters[r.Intn(len(letters))]
	}
	return string(result)
}

// HumanSelectName is a method to ask user to input a name for the human player
func HumanSelectName() string {
	fmt.Println("請設定該人類玩家名稱, 名稱必須在 3 到 5 個字元之間")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	return input
}

// AISelectCard is a method to randomly select a card from the hand cards for the AI player
func AISelectCard(handCards []card.Card) (card.Card, int) {
	randomIndex := rand.Intn(len(handCards))
	return handCards[randomIndex], randomIndex
}

// HumanSelectCard is a method to ask user to select a card from the hand cards for the human player
func HumanSelectCard(handCards []card.Card) (card.Card, int) {
	if len(handCards) == 0 {
		fmt.Println("No cards to play")
		return nil, -1
	}
	// print the hand cards
	fmt.Println("你的手牌如下:")
	for i, card := range handCards {
		fmt.Printf("%d. %s\n", i+1, card.ToString())
	}
	// ask the player to select a card to play
	fmt.Println("請選擇一張牌出牌")
	// Retry loop for valid input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		index, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("無效的輸入")
			continue
		}
		if index < 1 || index > len(handCards) {
			fmt.Println("請選擇 1 到", len(handCards), "之間的牌")
			continue
		}
		card := handCards[index-1]
		return card, index - 1
	}
}
