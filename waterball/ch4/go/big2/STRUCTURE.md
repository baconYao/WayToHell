# Big2 專案結構與類別說明

程式進入點：**`cmd/api/main.go`**

```
big2/
├── cmd/
│   └── api/
│       └── main.go              # 進入點：NewGame() 並 Run()
├── internal/
│   ├── card/                    # 撲克牌：Card, Rank, Suit，比較與解析
│   ├── deck/                    # 牌堆：建立、發牌、是否為空
│   ├── handCards/               # 手牌：增刪牌、排序、查詢
│   ├── cardPatternHandler/      # 牌型責任鏈：Single, Pair, Straight, FullHouse
│   ├── play/                    # 出牌：Play 介面，NormalPlay, PassPlay
│   ├── player/                  # 玩家：手牌、出牌決策、輸入
│   ├── round/                   # 回合：頂牌、PASS 數、首手規則、下一家
│   ├── game/                    # 遊戲流程：初始化、回合循環、勝負
│   └── utils/                   # 工具：ReadLineFromScanner 等
├── materials/
│   ├── OOD.md
│   ├── Sequential.md
│   └── STRUCTURE.md
├── tests/                       # 測試輸入/預期輸出
├── go.mod
└── Makefile
```

---

# 類別職責、屬性與操作

依據 `materials/OOD.md` 與現有實作整理。

---

## 1. Game

**職責**：負責整局遊戲的生命週期——從標準輸入讀取牌堆與玩家名稱、建立牌堆與四位玩家、發牌、建立牌型責任鏈、找出持有梅花 3 的玩家作為首出者，並驅動回合循環直到有人出完手牌後宣告贏家。

**屬性**：`deck`（牌堆）、`players`（四位玩家）、`patternChain`（牌型責任鏈鏈頭）、`winner`（贏家，nil 表示尚未結束）、`isFirstRound`（是否為遊戲首回合）、`starterIndex`（本回合先出牌者索引）、`scanner`（共用輸入來源）。

**操作**：`NewGame()`（建立並初始化一局）、`Run()`（執行主循環）、`findPlayerWithClub3()`（回傳持有梅花 3 的玩家索引）。

---

## 2. Deck

**職責**：代表一副牌堆。從規格字串建立牌堆、從頂部發牌、回報是否已空。牌序為 [底部 … 頂部]，發牌從頂部取牌。

**屬性**：`cards`（牌列表）。

**操作**：`NewFromShuffledCards(line)`、`Empty()`、`Deal()`（回傳 Card 與是否成功）。

---

## 3. Round

**職責**：代表「一回合」的完整生命週期與狀態，是遊戲循環中「輪流出牌直到三人 PASS」的單位。Round 不負責建立自己，由 Game 在每回合開始時 `New` 並 `Start`；其職責是：(1) 追蹤檯面頂牌（`topPlay`）與連續 PASS 次數（`passCount`），(2) 決定當前該誰出牌（`currentPlayerIndex`）並在每次出牌後切換下一家（`NextPlayer`），(3) 檢查首手規則（遊戲首回合且檯面尚無牌時，選牌必須含梅花 3），(4) 判斷回合是否結束（連續三人 PASS 即 `IsRoundEnded`），(5) 在回合結束時提供本回合贏家索引（`RoundWinnerIndex`），供 Game 作為下一回合的 `starterIndex`。Round 不持有 Player 或 HandCards，只透過 Play 的 `GetPlayerIndex()` 取得贏家索引，與 Game 的協定是「由 Game 依 `CurrentPlayerIndex()` 取玩家、呼叫 `AcceptPlay(playResult)` 更新狀態、再呼叫 `NextPlayer()`」。

**屬性**：
- `topPlay`：目前檯面上的頂牌（最後一次打出的 Play；nil 表示尚無人出牌）。
- `passCount`：本回合連續 PASS 的次數（0～3）；達 3 表示回合結束。
- `currentPlayerIndex`：當前應出牌的玩家索引（0..3）。
- `isFirstRoundOfGame`：是否為整局遊戲的第一回合（用於首手必須含梅花 3 的規則）。

**操作**：
- `New(isFirstRound)`：建立新回合（尚未指定先出牌者）。
- `Start(firstPlayerIndex)`：設定本回合先出牌者並重置 `topPlay`、`passCount`、`currentPlayerIndex`。
- `GetTopPlay()`：回傳目前頂牌（供 Player 比大小與判斷可否 PASS）。
- `CheckFirstMoveRule(cards)`：若為遊戲首回合且檯面尚無牌，檢查選牌是否含梅花 3。
- `IsRoundEnded()`：是否已結束（`passCount >= 3`）。
- `AcceptPlay(play)`：接受一次出牌——若為 PassPlay 則 `passCount++`，否則將 `topPlay` 設為該 Play 並將 `passCount` 歸零。
- `NextPlayer()`：將 `currentPlayerIndex` 切換至下一家（(index+1)%4）。
- `RoundWinnerIndex()`：回傳本回合贏家索引（有頂牌則為頂牌出牌者，否則為當前玩家）。

---

## 4. Player

**職責**：代表一位玩家。持有手牌與共用輸入來源；在輪到自己時透過 `Play(chain, round)` 讀取輸入、顯示手牌、解析出牌或 PASS，並以責任鏈驗證牌型、以 Round 檢查首手規則與頂牌，最後回傳合法的 Play（NormalPlay 或 PassPlay）。不負責回合流程，只負責「這一手」的決策與回傳。

**屬性**：`Index`、`Name`、`Hand`（HandCards）、`scanner`（共用輸入）。

**操作**：`New`、`SetHand`、`SetScanner`、`AddHandCard`、`Play(chain, round)`、`ShowCards`（與內部 `showHandOnly`）。

---

## 5. HandCards

**職責**：代表一位玩家的手牌集合。持牌並維持由小到大排序；提供加入、移除、是否為空、取得牌列表與依索引取牌，供 Player 顯示與出牌使用。

**屬性**：`cards`（牌列表）。

**操作**：`New`、`AddCard`、`RemoveCards`、`IsEmpty`、`GetCards`、`CardAt(i)`、`Len`。

---

## 6. Card

**職責**：代表單張撲克牌。封裝數字（Rank）與花色（Suit），提供與另一張牌的比較（先比數字再比花色），以及字串表示與解析。

**屬性**：`Rank`、`Suit`。

**操作**：`Compare(other)`、`String()`、`ParseCard`/`ParseCards`、`MustParseCard` 等。

---

## 7. Suit / Rank

**職責**：花色與數字的列舉。Suit：Club, Diamond, Heart, Spade。Rank：3, 4, …, 10, J, Q, K, A, 2。

---

## 8. Play（介面）

**職責**：定義「一次出牌」的契約。可能是實際出牌（NormalPlay）或放棄出牌（PassPlay）。提供出牌者索引、與頂牌比強弱、以及取得出牌內容（PassPlay 為空）。

**操作**：`GetPlayerIndex()`、`IsStrongerThan(other)`、`GetCards()`。

---

## 9. NormalPlay

**職責**：代表一次實際出牌。持有驗證通過的 CardPatternHandler（context）、複製的牌與比較用牌；透過 handler 取得牌型名稱與型別比較，透過 compareCard 與頂牌比大小。複製牌與比較牌是為了避免責任鏈被下一位玩家覆寫。

**屬性**：`handler`、`PlayerIndex`、`Cards`、`CompareCard`。

**操作**：`GetPlayerIndex`、`GetCards`、`PatternName()`（委派 handler.Name）、`Handler()`、`IsStrongerThan(other)`（先 handler.IsSameType 再 compareCard 比大小）、`NewNormalPlay(playerIndex, handler)`。

---

## 10. PassPlay

**職責**：代表放棄出牌。不包含牌；`IsStrongerThan` 一律回傳 false。

**屬性**：`PlayerIndex`。

**操作**：`GetPlayerIndex`、`GetCards()`（nil）、`IsStrongerThan`（false）。

---

## 11. CardPatternHandler（抽象類別，Chain of Responsibility）

**職責**：牌型責任鏈的抽象處理者，實作「Chain of Responsibility」模式。鏈上每個節點負責一種牌型（單張、對子、順子、葫蘆等）；呼叫端只與鏈頭互動，將選牌傳入 `Validate(cards)`，由鏈頭依序嘗試，若某節點符合規則則持有該牌並回傳自己（Handler），否則轉交給 `next`，若整條鏈都不符合則回傳 nil。同一條鏈在整局遊戲中由所有玩家共用，因此 NormalPlay 只保留「驗證通過當下」的 handler 參考與牌／比較牌的複本，避免下一位玩家驗證時覆寫節點內的 cards。責任鏈的組裝（Single → Pair → Straight → FullHouse）由 `BuildChain()` 完成，Game 只持有鏈頭並傳給 Player。擴充新牌型時只需新增一個 Handler 並接上鏈，符合 OCP。

**屬性**：
- `next`：鏈中下一個處理者（可為 nil）。
- `cards`：本節點驗證通過時持有的牌（會被下一次 Validate 覆寫，故 NormalPlay 需複製）。

**操作**：
- `SetNext(handler)`：設定下一節點。
- `Validate(cards)`：若本節點規則符合則設定 `cards` 並回傳自己，否則若 `next != nil` 則回傳 `next.Validate(cards)`，否則回傳 nil。
- `Name()`：牌型名稱（如「單張」「對子」）。
- `GetComparisonCard()`：本牌型的比較用牌（用於與頂牌比大小）。
- `IsSameType(other)`：與另一 Handler 是否為同一牌型（同型才能比大小）。
- `Cards()`：回傳目前持有的牌副本。

**具體類別**：Single（單張）、Pair（對子）、Straight（順子）、FullHouse（葫蘆），各自實作上述操作並在 `Validate` 內實作規則，不符合時轉交 next。

---

## 12. Single / Pair / Straight / FullHouse

**職責**：CardPatternHandler 的具體實作。Single：一張牌。Pair：兩張同數字。Straight：五張連續數字（含大老二特殊順）。FullHouse：三張同數字 + 兩張同數字。各節點在 `Validate` 中檢查張數與規則，符合則設定 `cards` 並回傳自己，否則委派給 `next`。

---

## 對應 OOD / 循序圖

| 目錄 / 套件       | 對應概念 |
|------------------|----------|
| `card`           | Card, Rank, Suit |
| `deck`           | Deck |
| `handCards`      | HandCards |
| `cardPatternHandler` | CardPatternHandler（抽象）、Single, Pair, Straight, FullHouse |
| `play`           | Play 介面，NormalPlay, PassPlay |
| `player`         | Player |
| `round`          | Round |
| `game`           | Game |
| `utils`          | 工具函式 |

## 執行

```bash
go run ./cmd/api
# 或
make run
make run NAME=normal-no-error-play1   # 指定測試案例
```
