# Big2 專案目錄結構

程式進入點：**`cmd/api/main.go`**

```
big2/
├── cmd/
│   └── api/
│       └── main.go          # 進入點：從 Stdin 建立遊戲並 Run()
├── internal/                # 僅供本模組使用的套件
│   ├── card/card.go         # 撲克牌：Card, Rank, Suit，比較邏輯
│   ├── deck/deck.go         # 牌堆：洗牌、發牌、輸入解析
│   ├── hand/hand.go         # 手牌 HandCards：增刪牌、是否為空
│   ├── player/player.go     # 玩家：名稱、手牌、出牌/PASS、與 pattern 鏈互動
│   ├── round/round.go       # 回合：topPlay、passCount、首手規則、下一家
│   ├── play/play.go         # 出牌型別：Play, NormalPlay, PassPlay，強弱比較
│   ├── pattern/pattern.go   # 牌型責任鏈（OCP）：Single, Pair, Straight, FullHouse
│   └── game/game.go         # 遊戲流程：初始化、回合循環、勝負判定
├── materials/
│   ├── OOD.md
│   └── Sequential.md
├── go.mod
├── Readme.md
└── STRUCTURE.md             # 本說明
```

## 對應 OOD / 循序圖

| 目錄       | 對應概念 |
|------------|----------|
| `card`     | Card, Rank, Suit |
| `deck`     | Deck |
| `hand`     | HandCards |
| `player`   | Player |
| `round`    | Round |
| `play`     | Play, NormalPlay, PassPlay |
| `pattern`  | CardPatternHandler 與 Single, Pair, Straight, FullHouse |
| `game`     | Game（協調上述元件並驅動流程） |

## 執行

```bash
go run ./cmd/api
# 或
make run
```
