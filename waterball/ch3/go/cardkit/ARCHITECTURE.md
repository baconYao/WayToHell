# 牌類遊戲框架 - 專案架構規劃

## 目錄結構

```
cardkit/
├── cmd/                    # 應用程式入口
│   ├── poker.go           # 撲克牌比大小遊戲入口
│   └── uno.go             # UNO 遊戲入口
│
├── internal/              # 內部包（不對外暴露）
│   │
│   ├── game/              # 遊戲框架核心（樣板方法模式）
│   │   ├── game.go       # Game 介面和樣板方法實作
│   │   ├── deck.go       # Deck 介面和基礎實作
│   │   └── hand.go       # Hand（手牌）介面和實作
│   │
│   ├── player/            # 玩家相關
│   │   ├── player.go     # Player 介面
│   │   ├── human.go      # 真實玩家實作（CLI 輸入）
│   │   └── ai.go         # AI 玩家實作（隨機選擇）
│   │
│   ├── card/              # 牌相關
│   │   ├── card.go       # Card 介面
│   │   ├── poker_card.go # 撲克牌實作（Rank + Suit）
│   │   └── uno_card.go   # UNO 牌實作（Color + Number）
│   │
│   ├── poker/             # 撲克牌遊戲具體實作
│   │   ├── poker_game.go # 撲克牌遊戲實作（繼承/實作 Game）
│   │   └── poker_deck.go # 撲克牌牌堆實作（52張牌）
│   │
│   └── uno/               # UNO 遊戲具體實作
│       ├── uno_game.go   # UNO 遊戲實作（繼承/實作 Game）
│       └── uno_deck.go   # UNO 牌堆實作（40張牌）
│
├── go.mod                 # Go 模組定義
├── go.sum                 # Go 依賴校驗
└── README.md              # 專案說明文檔
```

## 核心設計思路

### 1. 樣板方法模式（Template Method Pattern）

**`internal/game/game.go`** 定義遊戲流程樣板：

- `Initialize()` - 初始化遊戲（取名、洗牌）
- `DealCards()` - 抽牌階段（樣板方法）
- `PlayRounds()` - 遊戲回合循環（樣板方法）
- `DetermineWinner()` - 判定勝者（樣板方法）
- `Run()` - 主流程（樣板方法，呼叫上述方法）

具體遊戲（`poker_game.go`, `uno_game.go`）實作：

- `CreateDeck()` - 創建特定類型的牌堆
- `CompareCards()` - 比較牌的規則（撲克牌 vs UNO）
- `CanPlayCard()` - 判斷能否出牌（UNO 需要）
- `CheckWinCondition()` - 檢查勝利條件

### 2. 玩家抽象

**`internal/player/player.go`** 定義 Player 介面：

- `NameHimself()` - 取名
- `DrawCard()` - 抽牌
- `ShowCard()` / `PlayCard()` - 出牌
- `GetHand()` - 獲取手牌

實作：

- `human.go` - 透過 CLI 輸入進行選擇
- `ai.go` - 隨機選擇

### 3. 牌抽象

**`internal/card/card.go`** 定義 Card 介面：

- `Compare()` - 比較兩張牌
- `String()` - 顯示牌的內容

實作：

- `poker_card.go` - Rank (2-A) + Suit (Club, Diamond, Heart, Spade)
- `uno_card.go` - Color (BLUE, RED, YELLOW, GREEN) + Number (0-9)

### 4. 牌堆抽象

**`internal/game/deck.go`** 定義 Deck 介面：

- `Shuffle()` - 洗牌
- `Draw()` - 抽一張牌
- `IsEmpty()` - 是否為空
- `Reset()` - 重置牌堆（UNO 需要）

實作：

- `poker_deck.go` - 52 張撲克牌
- `uno_deck.go` - 40 張 UNO 牌

## 關鍵抽象點

### 共同流程（樣板方法）

1. 初始化：玩家取名 + 洗牌
2. 抽牌階段：輪流抽牌直到手牌數量達標
3. 遊戲循環：多回合出牌
4. 勝負判定：根據遊戲規則判定

### 差異點（需要子類實作）

1. **牌堆創建**：52 張撲克牌 vs 40 張 UNO 牌
2. **抽牌數量**：13 張 vs 5 張
3. **比牌規則**：撲克牌（Rank > Suit）vs UNO（顏色或數字匹配）
4. **出牌規則**：撲克牌（任意出）vs UNO（必須匹配）
5. **勝利條件**：撲克牌（13 回合後最高分）vs UNO（最快出完手牌）
6. **回合數**：固定 13 回合 vs 直到有人出完

## 設計原則

1. **開閉原則**：對擴展開放，對修改關閉
2. **單一職責**：每個模組只負責一個功能
3. **依賴倒置**：依賴抽象介面而非具體實作
4. **樣板方法**：將共同流程提取到父類，差異點由子類實作

## 下一步實作順序建議

1. 定義核心介面（Card, Player, Deck, Game）
2. 實作基礎類型（poker_card, uno_card）
3. 實作玩家類型（human, ai）
4. 實作牌堆（poker_deck, uno_deck）
5. 實作遊戲框架（game.go 樣板方法）
6. 實作具體遊戲（poker_game, uno_game）
7. 實作入口程式（cmd/poker.go, cmd/uno.go）
