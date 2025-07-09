package showdown

import (
	"errors"
	"fmt"
	"showdown/logger"
	"strings"
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
	log := logger.GetLogger()
	for r := range ROUNDS {
		fmt.Printf("============ Round %d ============\n", r+1)
		showedCards := make(map[string]Card, 0)
		for _, p := range s.players {
			// swap card back if the exchangedHand exists and matchs the iteration
			if p.GetExchangedHand() != nil && p.GetExchangedHand().GetSwapBackIteration() == r {
				log.Info("Exchanging cards back...")
				p.GetExchangedHand().Exchange()
			}
			// Ask player wether to exchange hands privilege
			if p.GetExchangedHand() == nil {
				yes, err := p.WantExchangeHands()
				if err != nil {
					return err
				}
				if yes {
					log.Info("Exchanging cards ...")
					err := p.ExchangeHands(r + 3)
					if err != nil {
						return err
					}
				}
			}
		}

		// Players show a card
		for _, p := range s.players {
			if len(p.GetCards()) > 0 {
				card, err := p.Show()
				if err != nil {
					return err
				}
				showedCards[string(p.GetName())] = card
			}
		}

		// Print the showed card
		s.printCard(showedCards)

		// Get the winner in this round and the winner gains 1 point
		wpn, err := s.compareCard(showedCards)
		if err != nil {
			return err
		}
		for _, p := range s.players {
			if string(p.GetName()) == wpn {
				log.Info("Player '%s' got 1 point", wpn)
				p.GainPoint(1)
				break
			}
		}
	}
	return nil
}

// printCard prints all players' card in each round
func (s Showdown) printCard(showedCards map[string]Card) {
	log := logger.GetLogger()
	for k, v := range showedCards {
		log.Info("Player '%s' shows '%s-%s'", k, v.GetSuit().String(), v.GetRank().String())
	}
}

// compareCard compares up to 4 unique cards in showedCards and returns the winner's name and an error
func (s Showdown) compareCard(showedCards map[string]Card) (string, error) {
	if len(showedCards) < 1 || len(showedCards) > 4 {
		return "", fmt.Errorf("showedCards must contain 1 to 4 cards, got %d", len(showedCards))
	}

	maxCardSuit := Suit(0)
	maxCardRank := Rank(0)
	playerName := ""

	for name, card := range showedCards {
		if !card.IsValid() {
			return "", fmt.Errorf("invalid card for %s: rank=%v, suit=%v", name, card.GetRank(), card.GetSuit())
		}

		// Compare rank and suit
		if card.GetRank() > maxCardRank || (card.GetRank() == maxCardRank && card.GetSuit() > maxCardSuit) {
			maxCardRank = card.GetRank()
			maxCardSuit = card.GetSuit()
			playerName = name
		}
	}

	return playerName, nil
}

// ShowWinner prints the final winner(s) who have the highest point
func (s Showdown) ShowWinner() {
	maxPoint := 0
	var winners []Player

	for _, p := range s.players {
		if p.GetPoint() > maxPoint {
			// New highest point, reset winners
			maxPoint = p.GetPoint()
			winners = []Player{p}
		} else if p.GetPoint() == maxPoint {
			// Equal to highest point, add to winners
			winners = append(winners, p)
		}
	}

	if len(winners) == 1 {
		fmt.Printf("Winner: %s with %d points\n", winners[0].GetName(), maxPoint)
	} else {
		names := make([]string, len(winners))
		for i, p := range winners {
			names[i] = string(p.GetName())
		}
		fmt.Printf("Tie: %s share %d points\n", strings.Join(names, " and "), maxPoint)
	}
}
