package showdown

import (
	"math/rand"
	"time"
)

type AI struct {
	BasePlayer
}

func NewAIPlayer() *AI {
	return &AI{BasePlayer: BasePlayer{
		name:          nil,
		point:         0,
		cards:         make([]Card, 0),
		privilege:     true,
		exchangedHand: nil,
	}}
}

// NameHimSelf randomly assigns a name for AI player
func (a *AI) NameHimSelf() error {
	letters := "abcdefghijklmnopqrstuvwxyz"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := r.Intn(5) + 4
	result := make([]byte, length)
	for i := range length {
		result[i] = letters[r.Intn(len(letters))]
	}
	a.setName(result)
	return nil
}
