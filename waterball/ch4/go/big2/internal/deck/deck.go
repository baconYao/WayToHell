// Package deck 定義牌堆 Deck：從輸入解析洗好的牌堆、發牌，以及是否為空。
// 規格：最左邊為牌堆底部、最右邊為牌堆頂部，發牌時從頂部取牌。
package deck

import "big2/internal/card"

// Deck 一副牌堆。牌序為 [底部 ... 頂部]，發牌從頂部（最後一張）取。
type Deck struct {
	cards []card.Card
}

// NewFromShuffledCards 從規格格式的一行字串建立牌堆（空白分隔，例如 "S[8] S[9] C[3] ..."）。
// 輸入順序為左＝底部、右＝頂部，發牌時會從頂部開始取。
func NewFromShuffledCards(line string) (*Deck, error) {
	cards, err := card.ParseCards(line)
	if err != nil {
		return nil, err
	}
	return &Deck{cards: cards}, nil
}

// Empty 回傳牌堆是否已空（無牌可發）。
func (d *Deck) Empty() bool {
	return len(d.cards) == 0
}

// Deal 從牌堆頂部發一張牌。若牌堆已空則回傳零值與 false。
func (d *Deck) Deal() (card.Card, bool) {
	if d.Empty() {
		return card.Card{}, false
	}
	last := len(d.cards) - 1
	c := d.cards[last]
	d.cards = d.cards[:last]
	return c, true
}
