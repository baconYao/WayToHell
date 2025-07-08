package showdown

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"showdown/logger"
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
		privilege:     true,
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
	panic("Show method not be implemented yet")
}

func (h *Human) ExchangeHands(switchBackIteration int) error {
	log := logger.GetLogger()
	if !h.GetPrivilege() && h.GetExchangedHand() != nil {
		return errors.New("cannot exchange hands again")
	}

	candidate, err := h.selectExchangeCandidate()
	if err != nil {
		return err
	}

	log.Info("Player '%s' exchanging hands with player '%s'", h.GetName(), candidate.GetName())
	eh, err := NewExchangedHand(h, candidate, switchBackIteration)
	if err != nil {
		return err
	}

	h.SetExchangedHand(eh)
	err = eh.Exchange()
	if err != nil {
		return err
	}

	err = h.SetPrivilege(false)
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
