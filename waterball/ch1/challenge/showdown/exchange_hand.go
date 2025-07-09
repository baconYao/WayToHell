package showdown

import (
	"errors"
	"fmt"
	"showdown/logger"
)

type ExchangedHand struct {
	swapBackRound int
	issuer        Player
	candidate     Player
}

func NewExchangedHand(issuer, candidate Player, swapBackRound int) (*ExchangedHand, error) {
	if issuer == nil || candidate == nil {
		return nil, errors.New("issuer or candidate cannot be nil")
	}

	if swapBackRound < 2 || swapBackRound > ROUNDS {
		return nil, errors.New("invalid switch back iteration")
	}

	return &ExchangedHand{
		swapBackRound: swapBackRound,
		issuer:        issuer,
		candidate:     candidate,
	}, nil
}

// Exchange swaps the hands card between issuer and candidate.
func (e *ExchangedHand) Exchange() error {
	log := logger.GetLogger()
	log.Info("Exchanging cards between '%s' and '%s'...", e.issuer.GetName(), e.candidate.GetName())

	issuerCards := e.issuer.GetCards()
	log.Debug("Cards of issuer '%s' -> %s", e.issuer.GetName(), PrettyCardsHelper(issuerCards))

	candidateCards := e.candidate.GetCards()
	log.Debug("Cards of candidate '%s' -> %s", e.candidate.GetName(), PrettyCardsHelper(candidateCards))

	if issuerCards == nil || candidateCards == nil {
		return fmt.Errorf("cannot exchnage: one or both players have nil cards")
	}

	if err := e.issuer.SetCards(candidateCards); err != nil {
		return fmt.Errorf("failed to set cards for issuer: %v", err)
	}
	log.Debug("After exchanging, cards of issuer '%s' -> %s", e.issuer.GetName(), e.candidate.GetName())

	if err := e.candidate.SetCards(issuerCards); err != nil {

		return fmt.Errorf("failed to set cards for candidate: %v", err)
	}
	log.Debug("After exchanging, cards of candidate '%s' -> %s", e.candidate.GetName(), PrettyCardsHelper(e.candidate.GetCards()))

	return nil
}

func (e ExchangedHand) GetswapBackRound() int {
	return e.swapBackRound
}

func (e ExchangedHand) GetIssuer() Player {
	return e.issuer
}

func (e ExchangedHand) GetCandidate() Player {
	return e.candidate
}
