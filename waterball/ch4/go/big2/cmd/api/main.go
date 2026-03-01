// Package main 為 Big2 遊戲的程式進入點。
// 執行方式: go run ./cmd/api 或 make run
package main

import (
	"fmt"
	"os"

	"big2/internal/game"
)

func main() {
	g, err := game.NewFromStdin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "初始化遊戲失敗: %v\n", err)
		os.Exit(1)
	}
	g.Run()
}
