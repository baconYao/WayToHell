package showdown

import (
	"errors"
	"fmt"
)

const ROUNDS int = 13

type Showdown struct {
	players []Player
	deck    *Deck
}

func NewShowdown(players []Player) (*Showdown, error) {
	if len(players) != 4 {
		return nil, errors.New("must need four players")
	}
	return &Showdown{players: players}, nil
}

// GetPlayers returns all players instance
func (s Showdown) GetPlayers() []Player {
	return s.players
}

// Start
func (s *Showdown) Start() error {
	fmt.Println("Start Showdown Game...")
	// Naming players
	for idx, p := range s.players {
		if err := p.NameHimSelf(); err != nil {
			return err
		}
		fmt.Printf("Player %d: %s\n", idx+1, p.GetName())
	}
	// Shuffle the deck
	deck := NewDeck()
	deck.Shuffle()
	s.deck = deck

	// Dispatch cards
	for r := range ROUNDS {
		fmt.Printf("Dispatch card iteration %d ...", r+1)
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

func (s Showdown) TakeTurns() error {
	for r := range ROUNDS {
		for _, p := range s.players {
			// switch card back if the exchangedHand exists and matchs the iteration
			if !p.GetPrivilege() && p.GetExchangedHand().GetSwitchBackIteration() == r {
				// TODO
				fmt.Println("Swtich Back")
				continue
			}

		}
	}
	return nil
}

func (s Showdown) printCard() error {
	return nil
}

func (s Showdown) compareCard() error {
	return nil
}

func (s Showdown) ShowWinner() error {
	// maxPoint := 0
	// for _, p := range s.players {
	// 	// switch card back if the exchangedHand exists and matchs the iteration
	// 	pp := p.GetPoint()

	// }
	return nil
}
