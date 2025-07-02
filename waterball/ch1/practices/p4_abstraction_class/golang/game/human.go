package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Human struct {
	player
}

func NewHuman(number int) Player {
	return &Human{player: player{number: number}}
}

func (h Human) Decide() Decision {
	fmt.Println("請出拳 (1) 剪刀 (2) 石頭 (3) 布：")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	num, err := strconv.Atoi(input)
	if err != nil || num < 1 || num > 3 {
		fmt.Println("只能輸入範圍 1~3 的數字，請再輸入一次。")
		return h.Decide() // 遞迴呼叫直到輸入有效
	}
	switch num {
	case 1:
		return Scissors
	case 2:
		return Stone
	case 3:
		return Paper
	default:
		fmt.Println("只能輸入範圍 1~3 的數字，請再輸入一次。")
		return h.Decide()
	}
}
