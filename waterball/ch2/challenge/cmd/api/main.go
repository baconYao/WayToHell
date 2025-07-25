package main

import (
	"flag"
	"fmt"
	"matchmaking-system/internal/matchmaking"
	"matchmaking-system/pkg/logger"
	"os"
)

func main() {
	// Define command-line flag for log level, default is INFO
	logLevel := flag.String("log_level", "INFO", "Log level: DEBUG, INFO, ERROR")
	matchStrategy := flag.String("match_strategy", "distance", "Match Strategy: distance, habit")
	flag.Parse()

	validStrategies := map[string]bool{
		"distance": true,
		"habit":    true,
	}
	if !validStrategies[*matchStrategy] {
		fmt.Fprintf(os.Stderr, "Error: Invalid match_strategy '%s'. Must be one of: distance, habit\n", *matchStrategy)
		os.Exit(1)
	}

	// Create Logger with the specified log level
	logger.ConfigureLogger(*logLevel, nil)
	log := logger.GetLogger()

	// Generate dummy data
	data := []matchmaking.Individual{
		{
			ID:      1,
			Gender:  matchmaking.Male,
			Age:     25,
			Intro:   "Love sports",
			Habits:  []string{"basketball", "cooking", "gaming"},
			Coord:   matchmaking.Coord{X: 10, Y: 20},
			Reverse: false,
		},
		{
			ID:      2,
			Gender:  matchmaking.Female,
			Age:     22,
			Intro:   "Enjoy reading",
			Habits:  []string{"cooking", "reading", "painting"},
			Coord:   matchmaking.Coord{X: 11, Y: 21},
			Reverse: false,
		},
		{
			ID:      3,
			Gender:  matchmaking.Male,
			Age:     30,
			Intro:   "Tech enthusiast",
			Habits:  []string{"gaming", "reading", "coding", "painting"},
			Coord:   matchmaking.Coord{X: 15, Y: 25},
			Reverse: false,
		},
		{
			ID:      4,
			Gender:  matchmaking.Male,
			Age:     19,
			Intro:   "Music Plyer",
			Habits:  []string{"music", "piano", "guitar", "drum"},
			Coord:   matchmaking.Coord{X: 7, Y: 3},
			Reverse: true, // Prefer Reverse
		},
		{
			ID:      5,
			Gender:  matchmaking.Female,
			Age:     22,
			Intro:   "Band vocalist",
			Habits:  []string{"music", "guitar", "sing"},
			Coord:   matchmaking.Coord{X: 8, Y: 2},
			Reverse: false,
		},
	}

	individuals := []*matchmaking.Individual{}
	for i := range data {
		individuals = append(individuals, &data[i])
	}

	matchmakingSystem, err := matchmaking.NewMatchmakingSystem(individuals, matchmaking.NewDistanceBased())
	if *matchStrategy == "habit" {
		// Dynamic adjust match strategy by setter
		matchmakingSystem.SetMatchType(matchmaking.NewHabitBased())
	}
	if err != nil {
		log.Error("error creating Matchmaking System: %v\n", err)
		return
	}
	if err = matchmakingSystem.Start(); err != nil {
		log.Error("%s", err)
	}
}
