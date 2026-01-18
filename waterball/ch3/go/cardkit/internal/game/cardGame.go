package game

import "cardkit/pkg/logger"

const (
	SHOWDOWN_GAME = "showdown"
	UNO_GAME      = "uno"
)

type CardGame struct {
	game   CardGameStrategy
	logger *logger.Logger
}

func NewCardGame(strategy CardGameStrategy) *CardGame {
	return &CardGame{
		game:   strategy,
		logger: logger.GetLogger(),
	}
}

func (c *CardGame) Start() {
	c.logger.Info("Starting the game...")
	c.game.Start()
}

// func (c *CardGame) SetGame(game CardGameStrategy) {
// 	c.game = game
// }
