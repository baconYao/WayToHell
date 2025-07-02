package main

import (
	"abstraction-class/game"
	"fmt"
)

func main() {
	g, err := game.NewGame(game.NewHuman(1), game.NewAI(2))
	if err != nil {
		fmt.Println(err)
	}
	g.Start()
}
