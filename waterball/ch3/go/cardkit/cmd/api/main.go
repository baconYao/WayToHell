package main

import (
	"flag"
	"os"

	"cardkit/internal/game"
	"cardkit/pkg/logger"
)

func main() {
	// Define command-line flag for log level, default is INFO
	logLevel := flag.String("log_level", "INFO", "Log level: DEBUG, INFO, ERROR")
	cardGameType := flag.String("card_game_type", "showdown", "Card game type: showdown, uno")
	flag.Parse()

	// Configure logger
	logger.ConfigureLogger(*logLevel, os.Stdout)

	// Create card game
	var gameStrategy game.CardGameStrategy
	switch *cardGameType {
	case game.SHOWDOWN_GAME:
		logger.GetLogger().Info("Creating Showdown game...")
		gameStrategy = game.NewShowdownGame()
	case game.UNO_GAME:
		logger.GetLogger().Info("Creating Uno game...")
		gameStrategy = game.NewUnoGame()
	default:
		logger.GetLogger().Error("Invalid card game type: %s", *cardGameType)
		os.Exit(1)
	}
	cardGame := game.NewCardGame(gameStrategy)
	cardGame.Start()
}
