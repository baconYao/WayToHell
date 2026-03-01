// Package game 負責遊戲流程：初始化（玩家、洗牌發牌、責任鏈）、尋找梅花 3 玩家、
// 回合循環、判定勝負，以及從標準輸入讀取並驅動整個流程。
package game

import (
	"io"
	"os"
)

// Game 代表一局大老二遊戲，負責初始化、回合循環與勝負判定。
// 依 OOD 與循序圖：內含 players, deck, patternChain, winner, isFirstRound，並協調 Round。
type Game struct {
	// TODO: 依 OOD 補上 players, deck, patternChain, winner, isFirstRound
	input io.Reader
}

// NewFromStdin 從標準輸入建立一局遊戲（讀取洗好的牌堆、四位玩家名稱、後續動作）。
func NewFromStdin() (*Game, error) {
	return &Game{input: os.Stdin}, nil
}

// Run 執行遊戲主循環：發牌、每回合輪流出牌直到有人出完手牌並宣告贏家。
func (g *Game) Run() {
	_ = g.input
	// TODO: 依 Readme 與 Sequential 實作完整流程
}
