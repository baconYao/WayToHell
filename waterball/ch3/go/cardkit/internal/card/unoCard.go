package card

import "fmt"

type Color int

const (
	Blue Color = iota
	Red
	Yellow
	Green
)

type UnoCard struct {
	Color  Color
	Number int
}

func (c UnoCard) ToString() string {
	colorName := map[Color]string{
		Blue:   "Blue",
		Red:    "Red",
		Yellow: "Yellow",
		Green:  "Green",
	}
	return fmt.Sprintf("<%s %d>", colorName[c.Color], c.Number)
}
