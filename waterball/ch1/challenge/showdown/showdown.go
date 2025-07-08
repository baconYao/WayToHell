package showdown

import (
	"errors"
	"fmt"
	"showdown/logger"
)

const ROUNDS int = 13

type Showdown struct {
	players []Player
	deck    *Deck
}

// NewShowdown initiates a showdown instance
func NewShowdown(players []Player) (*Showdown, error) {
	if len(players) != 4 {
		return nil, errors.New("must need four players")
	}
	game := Showdown{players: players, deck: nil}
	for _, p := range players {
		p.SetShowdown(&game)
	}
	return &game, nil
}

// GetPlayers returns all players instance
func (s Showdown) GetPlayers() []Player {
	return s.players
}

// Start prepares, shuffles and dispatches the cards to players
func (s *Showdown) Start() error {
	log := logger.GetLogger()
	log.Info("Start Showdown Game...")
	// Naming players
	for idx, p := range s.players {
		if err := p.NameHimSelf(); err != nil {
			return err
		}
		log.Info("Player %d: %s joins game\n", idx+1, p.GetName())
	}
	// Shuffle the deck
	deck := NewDeck()
	deck.Shuffle()
	s.deck = deck

	// Dispatch cards
	for r := range ROUNDS {
		log.Debug("Dispatch card iteration %d ...", r+1)
		for _, p := range s.players {
			card, err := deck.DrawCard()
			if err != nil {
				return err
			}
			if err := p.GainCard(card); err != nil {
				return err
			}
		}
	}

	return nil
}

// TakeTurns handles the procedure during each round
func (s Showdown) TakeTurns() error {
	// log := logger.GetLogger()
	for r := range ROUNDS {
		fmt.Printf("============ Round %d ============\n", r+1)
		// showedCards := make([]Card, 0)
		for _, p := range s.players {
			// swap card back if the exchangedHand exists and matchs the iteration
			if !p.GetPrivilege() && p.GetExchangedHand().GetSwapBackIteration() == r {
				p.GetExchangedHand().Exchange()
			}
			// Ask player wether to exchange hands privilege
			if p.GetPrivilege() {
				yes, err := p.WantExchangeHands()
				if err != nil {
					return err
				}
				if yes {
					err := p.ExchangeHands(r + 3)
					if err != nil {
						return err
					}
				}
			}
			// // Player shows a card
			// card, err := p.Show()
			// if err != nil {
			// 	return err
			// }
			// // TODO: handle empty card (nil)
			// showedCards = append(showedCards, card)

		}
	}
	return nil
}

// printCard prints all players' card in each round
func (s Showdown) printCard() error {
	return nil
}

func (s Showdown) compareCard() error {
	return nil
}

// ShowWinner prints the final winner who has the highest point
func (s Showdown) ShowWinner() error {
	// maxPoint := 0
	// for _, p := range s.players {
	// 	// switch card back if the exchangedHand exists and matchs the iteration
	// 	pp := p.GetPoint()

	// }
	return nil
}
