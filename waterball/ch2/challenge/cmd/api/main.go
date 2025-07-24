package main

import (
	"flag"
	"fmt"
	"matchmaking-system/internal/matchmaking"
	"matchmaking-system/pkg/logger"
)

func main() {
	fmt.Println("go!!")

	// Define command-line flag for log level, default is INFO
	logLevel := flag.String("log_level", "INFO", "Log level: DEBUG, INFO, ERROR")
	flag.Parse()

	// Create Logger with the specified log level
	logger.ConfigureLogger(*logLevel, nil)
	log := logger.GetLogger()

	// matchmaking.NewIndividual()

	matchmakingSystem, err := matchmaking.NewMatchmakingSystem(nil, matchmaking.NewDistanceBased())
	if err != nil {
		log.Error("error creating Matchmaking System: %v\n", err)
		return
	}
	if err = matchmakingSystem.Start(); err != nil {
		log.Error("%s", err)
	}
}
