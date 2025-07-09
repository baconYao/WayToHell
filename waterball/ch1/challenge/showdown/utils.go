package showdown

import (
	"fmt"
	"strings"
)

// PrettyCardsHelper returns a string that shows cards in the format "Index: Suit-Rank"
func PrettyCardsHelper(cards []Card) string {
	// log := logger.GetLogger()
	var output strings.Builder
	for i, card := range cards {
		if card.IsValid() {
			output.WriteString(fmt.Sprintf("%d: %s-%s", i, card.GetSuit().String(), card.GetRank().String()))
		} else {
			output.WriteString("Invalid card")
		}
		if i < len(cards)-1 {
			output.WriteString(", ")
		}
	}
	// log.Debug("%s", output.String())
	return output.String()
}
