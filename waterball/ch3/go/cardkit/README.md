# 牌類遊戲框架 (Card Game Framework)

基於樣板方法模式（Template Method Pattern）實作的牌類遊戲框架，支援多種牌類遊戲的快速開發。

## 專案結構

```
cardkit/
├── cmd/                    # 應用程式入口
│   ├── poker.go           # 撲克牌比大小遊戲入口
│   ├── uno.go             # UNO 遊戲入口
│   └── showdown.go        # (保留原有文件)
│
├── internal/              # 內部包（不對外暴露）
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
│   │   ├── poker_game.go # 撲克牌遊戲實作
│   │   └── poker_deck.go # 撲克牌牌堆實作（52張牌）
│   │
│   └── uno/               # UNO 遊戲具體實作
│       ├── uno_game.go   # UNO 遊戲實作
│       └── uno_deck.go   # UNO 牌堆實作（40張牌）
│
├── ARCHITECTURE.md        # 詳細架構設計文檔
└── README.md              # 本文件
```

## 設計模式

### 樣板方法模式（Template Method Pattern）

核心思想：將遊戲的共同流程抽象到 `internal/game/game.go` 中，具體遊戲的差異點由子類實作。

**共同流程（樣板方法）：**

1. `Initialize()` - 初始化遊戲（玩家取名、洗牌）
2. `DealCards()` - 抽牌階段
3. `PlayRounds()` - 遊戲回合循環
4. `DetermineWinner()` - 判定勝者
5. `Run()` - 主流程

**差異點（子類實作）：**

- 牌堆創建（52 張撲克牌 vs 40 張 UNO 牌）
- 抽牌數量（13 張 vs 5 張）
- 比牌規則（撲克牌 vs UNO）
- 出牌規則（任意出 vs 必須匹配）
- 勝利條件（最高分 vs 最快出完）

## 支援的遊戲

### 1. 簡易撲克牌比大小遊戲

- 4 位玩家（Human/AI）
- 每人 13 張手牌
- 13 回合比大小
- 勝利條件：最高分

### 2. 簡易 UNO

- 4 位玩家（Human/AI）
- 每人 5 張手牌
- 出牌必須匹配顏色或數字
- 勝利條件：最快出完手牌

## 開發狀態

⚠️ **當前狀態：架構規劃完成，尚未開始實作**

所有文件目前僅包含包聲明和 TODO 註釋，等待具體實作。

## 下一步

參考 `ARCHITECTURE.md` 了解詳細的架構設計和實作順序建議。
