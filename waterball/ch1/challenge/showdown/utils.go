package showdown

import (
	"fmt"
	"showdown/logger"
	"strings"
)

// PrintCards prints the cards in the format "Suit-Rank" using the custom logger
func PrintCardsHelper(cards []Card) {
	log := logger.GetLogger()
	var output strings.Builder
	for i, card := range cards {
		if card.IsValid() {
			output.WriteString(fmt.Sprintf("%s-%s", card.GetSuit().String(), card.GetRank().String()))
		} else {
			output.WriteString("Invalid card")
		}
		if i < len(cards)-1 {
			output.WriteString(", ")
		}
	}
	log.Debug("%s", output.String())
}
