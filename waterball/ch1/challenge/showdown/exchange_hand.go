package showdown

import "errors"

type ExchangedHand struct {
	switchBackIteration int
	issuer              *Player
	candidate           *Player
}

func NewExchangedHand(switchBackIteration int, issuer, candidate *Player) *ExchangedHand {
	return &ExchangedHand{
		switchBackIteration: switchBackIteration,
		issuer:              issuer,
		candidate:           candidate,
	}
}

func (e ExchangedHand) GetSwitchBackIteration() int {
	return e.switchBackIteration
}

func (e *ExchangedHand) SetSwitchBackIteration(iteration int) error {
	if iteration < 2 || iteration > ROUNDS {
		return errors.New("invalid switch back iteration")
	}
	e.switchBackIteration = iteration
	return nil
}

func (e ExchangedHand) GetIssuer() *Player {
	return e.issuer
}

func (e *ExchangedHand) SetIssuer(issuer *Player) error {
	if issuer == nil {
		return errors.New("cannot set empty issuer")
	}
	e.issuer = issuer
	return nil
}

func (e ExchangedHand) GetCandidate() *Player {
	return e.issuer
}

func (e *ExchangedHand) SetCandidate(candidate *Player) error {
	if candidate == nil {
		return errors.New("cannot set empty candidate")
	}
	e.candidate = candidate
	return nil
}
