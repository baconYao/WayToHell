package matchmaking

import "matchmaking-system/pkg/logger"

// DistanceBased implements the distance-based matching strategy
type DistanceBased struct {
	logger *logger.Logger
}

// NewDistanceBased creates a new DistanceBased instance
func NewDistanceBased() *DistanceBased {
	return &DistanceBased{
		logger: logger.GetLogger(),
	}
}
