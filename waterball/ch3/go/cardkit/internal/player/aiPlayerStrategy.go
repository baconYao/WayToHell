package player

import (
	"cardkit/internal/card"
	"math/rand"
	"time"
)

type AIStrategy[T card.Card] struct{}

func (s *AIStrategy[T]) DecideName() string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nameLength := r.Intn(3) + 3
	result := make([]byte, nameLength)
	for i := 0; i < nameLength; i++ {
		result[i] = letters[r.Intn(len(letters))]
	}
	return string(result)
}

func (s *AIStrategy[T]) DecideCard(handCards []T) int {
	if len(handCards) == 0 {
		return -1
	}
	randomIndex := rand.Intn(len(handCards))
	return randomIndex
}
