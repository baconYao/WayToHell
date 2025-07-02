package game

import (
	"errors"
	"fmt"
)

type Game struct {
	player1 Player
	player2 Player
}

func NewGame(p1, p2 Player) (*Game, error) {
	if p1.GetNumber() == p2.GetNumber() {
		return nil, errors.New("player numbers must be unique")
	}
	return &Game{player1: p1, player2: p2}, nil
}

// counterMap 定義勝負規則 (Key 贏 value)
var counterMap = map[Decision]Decision{
	Scissors: Paper,
	Paper:    Stone,
	Stone:    Scissors,
}

// Start starts a game
func (g *Game) Start() {
	p1Decision := g.player1.Decide()
	p2Decision := g.player2.Decide()

	fmt.Printf("玩家 #%d 出了 %s\n", g.player1.GetNumber(), p1Decision)
	fmt.Printf("玩家 #%d 出了 %s\n", g.player2.GetNumber(), p2Decision)

	if p1Decision == p2Decision {
		fmt.Println("平手！")
	} else if counterMap[p1Decision] == p2Decision {
		fmt.Printf("玩家 #%d 獲勝！\n", g.player1.GetNumber())
	} else {
		fmt.Printf("玩家 #%d 獲勝！\n", g.player2.GetNumber())
	}
}
