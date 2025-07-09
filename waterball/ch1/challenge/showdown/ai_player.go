package showdown

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"showdown/logger"
	"time"
)

type AI struct {
	BasePlayer
}

func NewAIPlayer() *AI {
	return &AI{BasePlayer: BasePlayer{
		showdown:      nil,
		name:          nil,
		point:         0,
		cards:         make([]Card, 0),
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

func (a *AI) Show() (Card, error) {
	log := logger.GetLogger()
	cards := a.GetCards()

	// Check if hand is empty
	if len(cards) == 0 {
		log.Info("No cards available to show")
		return Card{}, fmt.Errorf("no cards available")
	}

	// Create a new random number generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Randomly select a card
	index := rng.Intn(len(cards))
	selectedCard := cards[index]
	log.Debug("AI %s selected card: %s-%s", a.GetName(), selectedCard.GetSuit().String(), selectedCard.GetRank().String())

	// Remove the selected card from the hand
	newCards := append(cards[:index], cards[index+1:]...)
	a.SetCards(newCards)

	return selectedCard, nil
}

func (a *AI) WantExchangeHands() (bool, error) {
	log := logger.GetLogger()
	log.Debug("AI player %s is deciding on exchanging hands...", a.GetName())

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	wantExchange := r.Intn(2) == 1

	if wantExchange {
		log.Debug("AI player '%s' chose to exchange hands!", a.GetName())
	} else {
		log.Debug("AI player '%s' chose not to exchange hands!", a.GetName())
	}

	return wantExchange, nil
}

func (a *AI) ExchangeHands(switchBackIteration int) error {
	log := logger.GetLogger()
	if a.GetExchangedHand() != nil {
		return errors.New("cannot exchange hands again")
	}

	// Randomly choose a player and swap card with them
	candidate, err := a.selectExchangeCandidate()
	if err != nil {
		return err
	}

	log.Info("Player '%s' exchanging hands with player '%s'", a.GetName(), candidate.GetName())
	eh, err := NewExchangedHand(a, candidate, switchBackIteration)
	if err != nil {
		return err
	}

	a.SetExchangedHand(eh)
	err = eh.Exchange()
	if err != nil {
		return err
	}

	return nil
}

func (a *AI) ChooseExchangedHandCandidate() (Player, error) {
	log := logger.GetLogger()
	players := a.GetShowdown().GetPlayers()
	for _, p := range players {
		log.Info(string(p.GetName()))
	}
	return nil, nil
}

// selectExchangeCandidate selects a player for AI to exchange hands with
func (a *AI) selectExchangeCandidate() (Player, error) {
	var candidate Player

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Filter out valid candidates (exclude self)
	var validCandidates []Player
	for _, p := range a.GetShowdown().GetPlayers() {
		if !bytes.Equal(p.GetName(), a.GetName()) {
			validCandidates = append(validCandidates, p)
		}
	}

	// If no valid candidates, log and return nil
	if len(validCandidates) == 0 {
		return nil, errors.New("no valid players to exchange hands with")
	}

	// Randomly select a candidate
	candidate = validCandidates[rng.Intn(len(validCandidates))]

	return candidate, nil
}
