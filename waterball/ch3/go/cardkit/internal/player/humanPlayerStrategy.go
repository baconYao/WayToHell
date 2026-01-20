package player

import (
	"bufio"
	"cardkit/internal/card"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HumanStrategy[T card.Card] struct{}

func (s *HumanStrategy[T]) DecideName() string {
	fmt.Println("請設定該人類玩家名稱 (3 到 5 個字元):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func (s *HumanStrategy[T]) DecideCard(handCards []T) int {
	if len(handCards) == 0 {
		fmt.Println("無牌可出")
		return -1
	}
	fmt.Println("你的手牌如下:")
	for i, c := range handCards {
		fmt.Printf("%d. %s\n", i+1, c.ToString())
	}
	fmt.Print("請選擇出牌編號: ")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		index, err := strconv.Atoi(input)
		if err != nil || index < 1 || index > len(handCards) {
			fmt.Printf("無效輸入，請輸入 1 到 %d: ", len(handCards))
			continue
		}
		return index - 1
	}
}
