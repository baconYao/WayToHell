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
		return nil, fmt.Errorf("Match type is not set")
	}

	return &MatchmakingSystem{
		individuals: individuals,
		matchType:   matchType,
		logger:      logger.GetLogger(),
	}, nil
}

func (m MatchmakingSystem) Start() {
	// TODO
}

func (m *MatchmakingSystem) Match(matcher *Individual) (*Individual, error) {
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

func (m *MatchmakingSystem) setMatchType(matchType MatchType) {
	m.matchType = matchType
}
