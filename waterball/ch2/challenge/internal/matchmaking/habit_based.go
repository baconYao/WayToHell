package matchmaking

import "matchmaking-system/pkg/logger"

// HabitBased implements the habit-based matching strategy
type HabitBased struct {
	logger *logger.Logger
}

// NewHabitBased creates a new HabitBased instance
func NewHabitBased() *DistanceBased {
	return &DistanceBased{
		logger: logger.GetLogger(),
	}
}

func (h *HabitBased) Match(matcher *Individual, matchees []*Individual) (*Individual, error) {
	h.logger.Debug("HabitBased Match...")
	return nil, nil
}
