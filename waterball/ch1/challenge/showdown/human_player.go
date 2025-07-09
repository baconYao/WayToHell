package showdown

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"showdown/logger"
	"strconv"
	"strings"
)

type Human struct {
	BasePlayer
}

func NewHumanPlayer() *Human {
	return &Human{BasePlayer: BasePlayer{
		showdown:      nil,
		name:          nil,
		point:         0,
		cards:         make([]Card, 0),
		exchangedHand: nil,
	}}
}

// NameHimSelf assigns the name for human player by CLI
func (h *Human) NameHimSelf() error {
	log := logger.GetLogger()
	log.Info("Please give a name for this human player")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	err := h.setName([]byte(input))
	if err != nil {
		return err
	}

	return nil
}

func (h *Human) Show() (Card, error) {
	log := logger.GetLogger()
	cards := h.GetCards()

	// Check if hand is empty
	if len(cards) == 0 {
		log.Info("No cards available to show")
		return Card{}, fmt.Errorf("no cards available")
	}

	log.Info("Show a card, please enter the index from the following cards...")
	log.Info("Cards: %s", PrettyCardsHelper(cards))

	// Retry loop for valid input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		index, err := strconv.Atoi(input)
		if err != nil {
			log.Info("Invalid input: please enter a number")
			continue
		}

		if index < 0 || index >= len(cards) {
			log.Info("Invalid index: please enter a number between 0 and %d", len(cards)-1)
			continue
		}

		// Get selected card
		selectedCard := cards[index]
		log.Info("Selected card: %s-%s", selectedCard.GetSuit().String(), selectedCard.GetRank().String())

		// Remove the selected card from the hand
		newCards := append(cards[:index], cards[index+1:]...)
		h.SetCards(newCards)

		return selectedCard, nil
	}
}

func (h *Human) ExchangeHands(switchBackIteration int) error {
	log := logger.GetLogger()
	if h.GetExchangedHand() != nil {
		return errors.New("cannot exchange hands again")
	}

	candidate, err := h.selectExchangeCandidate()
	if err != nil {
		return err
	}

	log.Debug("Player '%s' exchanging hands with player '%s'", h.GetName(), candidate.GetName())
	eh, err := NewExchangedHand(h, candidate, switchBackIteration)
	if err != nil {
		return err
	}

	h.SetExchangedHand(eh)
	err = eh.Exchange()
	if err != nil {
		return err
	}

	return nil
}

func (h *Human) WantExchangeHands() (bool, error) {
	log := logger.GetLogger()
	log.Info("Prompting human player %s to decide on exchanging hands", h.GetName())

	fmt.Printf("Player %s, do you want to exchange hands? (yes/no): ", h.GetName())
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.ToLower(strings.TrimSpace(scanner.Text()))

	log.Debug("Human player %s input: %s", h.GetName(), input)

	switch input {
	case "yes", "y":
		log.Debug("Human player %s chose to exchange hands", h.GetName())
		return true, nil
	case "no", "n":
		log.Debug("Human player %s chose not to exchange hands", h.GetName())
		return false, nil
	default:
		log.Error("Invalid input by human player %s: %s", h.GetName(), input)
		return false, fmt.Errorf("invalid input: %s, please enter 'yes' or 'no'", input)
	}
}

// selectExchangeCandidate selects a player for Human player to exchange hands with by CLI
func (h *Human) selectExchangeCandidate() (Player, error) {
	log := logger.GetLogger()
	var candidate Player

	// CLI choose other players
	players := h.GetShowdown().GetPlayers()
	for _, p := range players {
		if yes := bytes.Equal(h.GetName(), p.GetName()); !yes {
			log.Info("Player '%s'", string(p.GetName()))
		}
	}
	log.Info("Choose a player from above and exchange hands with them...")

	ask := true
	for ask {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := []byte(strings.ToLower(strings.TrimSpace(scanner.Text())))
		for _, p := range players {
			if yes := bytes.Equal(input, p.GetName()); yes {
				if yes := bytes.Equal(input, h.GetName()); yes {
					log.Info("Cannot exchange hands with yourself. Please try again")
					break
				} else {
					candidate = p
					ask = false
					break
				}
			}
		}
	}

	return candidate, nil
}
