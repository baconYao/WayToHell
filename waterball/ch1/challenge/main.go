package main

import (
	"flag"
	"showdown/logger"
	"showdown/showdown"
)

func main() {
	// Define command-line flag for log level, default is INFO
	logLevel := flag.String("log_level", "INFO", "Log level: DEBUG, INFO, ERROR")
	flag.Parse()

	// Create Logger with the specified log level
	log := logger.NewLogger(*logLevel)

	// Create four players
	log.Debug("Creating four players...\n")
	players := []showdown.Player{showdown.NewHumanPlayer(), showdown.NewAIPlayer(), showdown.NewAIPlayer(), showdown.NewAIPlayer()}
	game, err := showdown.NewShowdown(players)
	if err != nil {
		log.Error("error creating Showdown: %v\n", err)
		return
	}
	if err := game.Start(); err != nil {
		log.Error("error: %v\n", err)
		return
	}

}
