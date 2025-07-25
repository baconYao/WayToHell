package matchmaking

import (
	"matchmaking-system/pkg/logger"
	"math"
)

// HabitBased implements the habit-based matching strategy
type HabitBased struct {
	logger *logger.Logger
}

// NewHabitBased creates a new HabitBased instance
func NewHabitBased() *HabitBased {
	return &HabitBased{
		logger: logger.GetLogger(),
	}
}

func (h *HabitBased) Match(matcher *Individual, matchees []*Individual) (*Individual, error) {
	h.logger.Debug("HabitBased Match...")
	var candidate *Individual

	if !matcher.isUsingReverse() {
		h.logger.Debug("Macther is not using reverse strategy, matching ...")
		bestCommon := -1
		for _, other := range matchees {
			if other.getID() == matcher.getID() {
				continue
			}
			common := h.countCommonHabits(matcher, other)
			if bestCommon == common {
				if candidate == nil || other.getID() < candidate.getID() {
					candidate = other
					bestCommon = common
				}
			} else if bestCommon < common {
				candidate = other
				bestCommon = common
			}
		}
	} else {
		h.logger.Debug("Macther is using reverse strategy, matching ...")
		worstCommon := int(math.Inf(1))
		for _, other := range matchees {
			if other.getID() == matcher.getID() {
				continue
			}
			common := h.countCommonHabits(matcher, other)
			if worstCommon == common {
				if candidate == nil || other.getID() < candidate.getID() {
					candidate = other
					worstCommon = common
				}
			} else if worstCommon > common {
				candidate = other
				worstCommon = common
			}
		}
	}

	return candidate, nil
}

// countCommonHabits counts the number of common habits between two Individuals
func (h HabitBased) countCommonHabits(matcher, matchee *Individual) int {
	habbits := make([]string, 0)
	count := 0
	for _, h1 := range matcher.getHabits() {
		for _, h2 := range matchee.getHabits() {
			if h1 == h2 {
				count++
				habbits = append(habbits, h1)
			}
		}
	}
	h.logger.Debug("Common habits %v with ID: %d", habbits, matchee.getID())
	return count
}
