package matchmaking

import (
	"fmt"
	"matchmaking-system/pkg/logger"
)

// MatchmakingSystem manages Individuals and uses MatchType for matching
type MatchmakingSystem struct {
	individuals []*Individual
	matchType   MatchType
	logger      *logger.Logger
}

// NewMatchmakingSystem creates a new MatchmakingSystem instance
func NewMatchmakingSystem(individuals []*Individual, matchType MatchType) (*MatchmakingSystem, error) {
	if matchType == nil {
		return nil, fmt.Errorf("matchType is not set")
	}

	return &MatchmakingSystem{
		individuals: individuals,
		matchType:   matchType,
		logger:      logger.GetLogger(),
	}, nil
}

func (m MatchmakingSystem) Start() error {
	m.logger.Info("Start the Matchmaking System")
	if len(m.individuals) < 2 {
		m.logger.Error("There are currently '%d' participants, but at least two are required.", len(m.individuals))
		return fmt.Errorf("not enough participants to form a match")
	}

	for _, i := range m.individuals {
		matchee, err := m.match(i)
		if err != nil {
			return err
		}
		m.logger.Info("Matcher ID '%d' (Reverse: %v) and Matchee ID '%d'", i.getID(), i.isUsingReverse(), matchee.getID())
	}
	return nil
}

func (m *MatchmakingSystem) match(matcher *Individual) (*Individual, error) {
	m.logger.Debug("Matching ID: '%d'", matcher.getID())
	return m.matchType.Match(matcher, m.individuals)
}

func (m MatchmakingSystem) getIndividuals() []*Individual {
	return m.individuals
}

func (m *MatchmakingSystem) setIndividuals(individuals []*Individual) {
	m.individuals = individuals
}

func (m MatchmakingSystem) getMatchType() MatchType {
	return m.matchType
}

func (m *MatchmakingSystem) SetMatchType(matchType MatchType) {
	m.matchType = matchType
}
