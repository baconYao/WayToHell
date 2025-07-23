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
