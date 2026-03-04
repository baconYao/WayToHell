// Package round 定義回合 Round：頂牌 topPlay、PASS 計數、當前玩家、首回合規則（梅花 3）、回合是否結束、下一家。
package round

import (
	"big2/internal/card"
	"big2/internal/play"
)

const NumPlayers = 4

// Round 一回合：頂牌、連續 PASS 數、當前玩家索引、是否為遊戲首回合。
type Round struct {
	topPlay            play.Play
	passCount          int
	currentPlayerIndex int
	isFirstRoundOfGame bool
}

// New 建立新回合。isFirstRound 表示是否為整局遊戲的第一回合（首手必須含梅花 3）。
func New(isFirstRound bool) *Round {
	return &Round{
		passCount:          0,
		isFirstRoundOfGame: isFirstRound,
	}
}

// Start 設定此回合由 firstPlayerIndex 先出牌。
func (r *Round) Start(firstPlayerIndex int) {
	r.currentPlayerIndex = firstPlayerIndex
	r.topPlay = nil
	r.passCount = 0
}

// GetTopPlay 回傳目前檯面上的頂牌（最後一次打出的牌型）；若尚無則為 nil。
func (r *Round) GetTopPlay() play.Play {
	return r.topPlay
}

// CurrentPlayerIndex 回傳當前要出牌的玩家索引 0..3。
func (r *Round) CurrentPlayerIndex() int {
	return r.currentPlayerIndex
}

// CheckFirstMoveRule 檢查首手規則：若為遊戲首回合且檯面尚無頂牌，選牌必須包含梅花 3。
func (r *Round) CheckFirstMoveRule(cards []card.Card) bool {
	if !r.isFirstRoundOfGame || r.topPlay != nil {
		return true
	}
	club3 := card.MustParseCard("C[3]")
	for _, c := range cards {
		if c.Compare(club3) == 0 {
			return true
		}
	}
	return false
}

// IsRoundEnded 是否已結束（連續三人 PASS）。
func (r *Round) IsRoundEnded() bool {
	return r.passCount >= 3
}

// NextPlayer 輪到下一家。
func (r *Round) NextPlayer() {
	r.currentPlayerIndex = (r.currentPlayerIndex + 1) % NumPlayers
}

// AcceptPlay 接受一次出牌並更新頂牌與 PASS 計數。
func (r *Round) AcceptPlay(p play.Play) {
	if _, isPass := p.(*play.PassPlay); isPass {
		r.passCount++
		return
	}
	r.topPlay = p
	r.passCount = 0
}

// RoundWinnerIndex 回合結束時，回傳本回合頂牌玩家索引（即下一回合先出牌者）。
func (r *Round) RoundWinnerIndex() int {
	if r.topPlay == nil {
		return r.currentPlayerIndex
	}
	return r.topPlay.GetPlayerIndex()
}
