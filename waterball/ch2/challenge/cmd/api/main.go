package main

import (
	"flag"
	"matchmaking-system/internal/matchmaking"
	"matchmaking-system/pkg/logger"
)

func main() {
	// Define command-line flag for log level, default is INFO
	logLevel := flag.String("log_level", "INFO", "Log level: DEBUG, INFO, ERROR")
	flag.Parse()

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
			Habits:  []string{"cooking", "reading"},
			Coord:   matchmaking.Coord{X: 11, Y: 21},
			Reverse: false,
		},
		{
			ID:      3,
			Gender:  matchmaking.Male,
			Age:     30,
			Intro:   "Tech enthusiast",
			Habits:  []string{"gaming", "reading"},
			Coord:   matchmaking.Coord{X: 15, Y: 25},
			Reverse: false,
		},
	}

	individuals := []*matchmaking.Individual{}
	for i := range data {
		individuals = append(individuals, &data[i])
	}

	matchmakingSystem, err := matchmaking.NewMatchmakingSystem(individuals, matchmaking.NewDistanceBased())
	if err != nil {
		log.Error("error creating Matchmaking System: %v\n", err)
		return
	}
	if err = matchmakingSystem.Start(); err != nil {
		log.Error("%s", err)
	}
}
