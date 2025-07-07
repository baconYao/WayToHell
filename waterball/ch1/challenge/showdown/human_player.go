package showdown

import (
	"bufio"
	"fmt"
	"os"
)

type Human struct {
	BasePlayer
}

func NewHumanPlayer() *Human {
	return &Human{BasePlayer: BasePlayer{
		name:          nil,
		point:         0,
		cards:         make([]Card, 0),
		privilege:     true,
		exchangedHand: nil,
	}}
}

// NameHimSelf assigns the name for human player by CLI
func (h *Human) NameHimSelf() error {
	fmt.Println("Please give a name for this human player")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	err := h.setName([]byte(input))
	if err != nil {
		return err
	}

	return nil
}
