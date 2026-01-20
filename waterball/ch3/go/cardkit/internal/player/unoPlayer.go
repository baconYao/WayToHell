package player

import (
	"cardkit/internal/card"
)

type UnoPlayer struct {
	BasePlayer[card.UnoCard]
}

type AIUnoPlayer struct {
	UnoPlayer
}

func NewUnoPlayer(playerBehaviorStrategy PlayerBehaviorStrategy[card.UnoCard]) *UnoPlayer {
	return &UnoPlayer{
		BasePlayer: NewBasePlayer[card.UnoCard](playerBehaviorStrategy),
	}
}
