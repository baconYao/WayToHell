package showdown

import (
	"errors"
	"fmt"
	"showdown/logger"
)

type ExchangedHand struct {
	swapBackIteration int
	issuer            Player
	candidate         Player
}

func NewExchangedHand(issuer, candidate Player, swapBackIteration int) (*ExchangedHand, error) {
	if issuer == nil || candidate == nil {
		return nil, errors.New("issuer or candidate cannot be nil")
	}

	if swapBackIteration < 2 || swapBackIteration > ROUNDS {
		return nil, errors.New("invalid switch back iteration")
	}

	return &ExchangedHand{
		swapBackIteration: 0,
		issuer:            issuer,
		candidate:         candidate,
	}, nil
}

// Exchange swaps the hands card between issuer and candidate.
func (e *ExchangedHand) Exchange() error {
	log := logger.GetLogger()
	log.Info("Starting card swap between issuer and candidate...")

	issuerCards := e.issuer.GetCards()
	log.Debug("Cards of issuer '%s'", e.issuer.GetName())
	PrintCardsHelper(issuerCards)

	candidateCards := e.candidate.GetCards()
	log.Debug("Cards of candidate '%s'", e.candidate.GetName())
	PrintCardsHelper(candidateCards)

	if issuerCards == nil || candidateCards == nil {
		return fmt.Errorf("cannot swap: one or both players have nil cards")
	}

	log.Debug("Swaping...")
	if err := e.issuer.SetCards(candidateCards); err != nil {
		return fmt.Errorf("failed to set cards for issuer: %v", err)
	}
	log.Debug("After swaping, cards of issuer '%s'", e.issuer.GetName())
	PrintCardsHelper(e.issuer.GetCards())

	log.Debug("After swaping, cards of candidate '%s'", e.candidate.GetName())
	if err := e.candidate.SetCards(issuerCards); err != nil {

		return fmt.Errorf("failed to set cards for candidate: %v", err)
	}
	PrintCardsHelper(e.candidate.GetCards())

	return nil
}

func (e ExchangedHand) GetSwapBackIteration() int {
	return e.swapBackIteration
}

func (e ExchangedHand) GetIssuer() Player {
	return e.issuer
}

func (e ExchangedHand) GetCandidate() Player {
	return e.candidate
}
