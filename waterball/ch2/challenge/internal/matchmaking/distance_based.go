package matchmaking

import (
	"matchmaking-system/pkg/logger"
	"math"
)

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

func (d *DistanceBased) Match(matcher *Individual, matchees []*Individual) (*Individual, error) {
	d.logger.Debug("Starting DistanceBased match")
	var candidate *Individual
	matcherCoord := matcher.getCoord()

	if !matcher.isUsingReverse() {
		d.logger.Debug("Macther is not using reverse strategy, matching the closest one...")

		closestDist := int(math.Inf(1))
		for _, other := range matchees {
			if other.getID() == matcher.getID() {
				continue
			}
			dist := calculateDistance(matcherCoord, other.getCoord())
			d.logger.Debug("Distance between ID %d and ID %d: %d", matcher.getID(), other.getID(), dist)
			if dist <= closestDist {
				if candidate == nil || candidate.getID() > other.getID() {
					candidate = other
				}
			}
		}
	} else {
		d.logger.Info("Macther is using reverse strategy, matching the farthest one...")
		closestDist := -1
		for _, other := range matchees {
			if other.getID() == matcher.getID() {
				continue
			}
			dist := calculateDistance(matcherCoord, other.getCoord())
			d.logger.Debug("Distance between ID %d and ID %d: %d", matcher.getID(), other.getID(), dist)
			if dist >= closestDist {
				if candidate == nil || candidate.getID() > other.getID() {
					candidate = other
				}
			}
		}
	}
	return candidate, nil
}

func calculateDistance(a, b Coord) int {
	x := a.getX() - b.getX()
	y := a.getY() - b.getY()
	return int(math.Sqrt(float64(x*x + y*y)))
}
