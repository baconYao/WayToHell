# Card Game Framework

## How to play?

```go
// Play Showdown Game
$ make run
// Play Uno Game
$ make run CARD_GAME_STRATEGY=uno

// Debug
$ make dev
$ make dev CARD_GAME_STRATEGY=uno
```

## Requirements Specification

### **Template Method — Let’s Build a Card Game Framework!**

You are required to implement the following two simple card games.  
Your goal is to **minimize duplicated code as much as possible** across the implementations.

### **Game 1: Simple High-Card Poker Game**

1. This game supports **four players**. Each player can be either a *Human Player* or an *AI Player*.
    1. Player implementations:
        1. **Human Player**: makes choices via a **Command Line Interface (CLI)**.
        2. **AI Player**: makes choices randomly.
    2. In the following requirements, **P1, P2, P3, and P4** represent the **first, second, third, and fourth players**, respectively.

2. The game contains a **Deck**.
    1. The deck initially contains **52 Cards**.
    2. Each card has a **Rank** and a **Suit**.

3. The game proceeds according to the following steps (a–d):
    1. **Game Initialization**
        1. P1–P4 take turns to **name themselves**.
        2. The deck is **shuffled**.
    2. **Card Drawing Phase**  
       Starting from P1, players P1–P4 take turns drawing cards from the deck until **each player has 13 cards in hand**.
    3. After all cards are drawn, the game continues for **13 rounds**. In each round:
        1. P1–P4 take turns to **play (show) one card**  
           (players do not know what cards others have played at this stage).
        2. The cards played by P1–P4 are revealed.
        3. The played cards are **compared to determine the winner**, and **1 point** is awarded to the winning player.
    4. After 13 rounds, all players will have played all their cards and the game ends.  
       The player with the **highest score** is declared the winner, and the **winner’s name is displayed**.

4. **Card Comparison Rules**
    1. Cards are first compared by **rank**; the higher rank wins.  
       If ranks are equal, compare **suits**; the higher suit wins.
    2. Rank order from lowest to highest:  
       `2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K, A`
    3. Suit order from lowest to highest:  
       `Club, Diamond, Heart, Spade`

---

### **Game 2: Simple UNO**

1. This game supports **four players**. Each player can be either a *Human Player* or an *AI Player*.
    1. Player implementations:
        1. **Human Player**: makes choices via a **Command Line Interface (CLI)**.
        2. **AI Player**: makes choices randomly.
    2. In the following requirements, **P1, P2, P3, and P4** represent the **first, second, third, and fourth players**, respectively.
    3. The game contains a **Deck**.
        1. Each card has a **Color** and a **Number**.
        2. The deck initially contains **40 Cards**:  
           4 colors (`BLUE`, `RED`, `YELLOW`, `GREEN`) × numbers `0–9`.
    4. The game proceeds according to the following steps (i–iii):
        1. **Game Initialization**
            1. P1–P4 take turns to **name themselves**.
            2. The deck is **shuffled**.
        2. **Card Drawing Phase**  
           Starting from P1, players P1–P4 take turns drawing cards until **each player has 5 cards in hand**.
        3. **Game Execution**
            1. Reveal the first card from the deck and place it on the table.
            2. Starting from P1, players take turns in the order:  
               `P1 → P2 → P3 → P4 → P1 → ...`
            3. A played card must match the **color** or **number** of the latest card on the table.  
               The played card becomes the new latest card on the table.
            4. The first player to **play all cards in their hand** wins the game.
            5. If a player has no playable cards, the player must **draw one card from the deck**.  
               If the deck is empty, all table cards except the latest one are returned to the deck and **reshuffled**.


## Requirements Specification (Chinese Version)

### **樣板方法——寫一個牌類遊戲框架吧！**

你要實作以下兩款簡單的牌類遊戲，而你的任務是將程式中***重複程式碼的部分減至越少越好***。

### **第一款：簡易撲克牌比大小遊戲**

1. 這款遊戲能支援四位玩家：玩家可以為*真實玩家 (Human Player)* 也能為*電腦玩家 (AI Player)*。
    1. 玩家的實作：
        1. 真實玩家：使用指令介面輸入 (Command Line Interface) 來做選擇。
        2. 電腦玩家：隨機做選擇。
    2. 在以下需求中：我們用 *P1, P2, P3, P4 來表示第一、二、三和第四順位的玩家。*
2. 遊戲中有一副*牌堆 (Deck)*。
    1. 牌堆中一開始存有 52 張*牌 (Card)*。
    2. 每張牌都會擁有*階級 (Rank)* 及*花色 (Suit)*。
3. 遊戲依照以下 a~d 的流程進行：
    1. 遊戲開始時，依序執行以下：
        1. 請 P1~P4 為自己取名 (Name himself)。
        2. 牌堆會進行洗牌 (Shuffle)。
    2. **抽牌階段：** 由 P1 開始，P1~P4 輪流從牌堆中抽牌 (Draw Card)，直到所有人都擁有手牌 (Hand) 13 張牌為止。
    3. 抽完牌後，在接下來的 13 回合中，每一回合依序執行以下：
        1. P1~P4 輪流 (Takes a turn) 出 (Show) 一張牌（此步驟彼此皆無法知曉彼此出的牌）。
        2. 顯示 P1~P4 各出的牌的內容。
        3. 將 P1~P4 出的牌進行*比大小決勝負*，將最勝者的分數(Point)加一。
    4. 13 回合後，P1~P4 皆已出完全部的牌，遊戲結束。取得最多分數的玩家為勝者，將勝者的名稱顯示出來。
4. **牌與牌之間的比大小決勝規則**：
    1. 先比較牌的階級，此時階級較大者勝，如果階級相同則比較花色，此時花色較大者勝。
    2. 階級由小到大依序為：2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K, A
    3. 花色由小到大依序為：梅花 (Club)、菱形 (Diamond)、愛心 (Heart)、黑桃 (Spade)

### **第二款：簡易 UNO**

1. 這款遊戲能支援四位玩家：玩家可以為*真實玩家 (Human Player)* 也能為*電腦玩家 (AI Player)*。
    1. 玩家的實作：
        1. 真實玩家：使用指令介面輸入 (Command Line Interface) 來做選擇。
        2. 電腦玩家：隨機做選擇。
    2. 在以下需求中：我們用 *P1, P2, P3, P4 來表示第一、二、三和第四順位的玩家。*
    3. 遊戲中有一副*牌堆 (Deck)*。
        1. 每張牌都會擁有顏色 *(Color)* 及數字 *(Number)*。
        2. 牌堆中一開始存有 40 張*牌 (Card)：4 種顏色 (BLUE, RED, YELLOW, GREEN) x 10 個數字 (0~9)。*
    4. 遊戲依照以下 i~iii 的流程進行：
        1. 遊戲開始時，依序執行以下：
            1. 請 P1~P4 為自己取名 (Name himself)。
            2. 牌堆會進行洗牌 (Shuffle)。
        2. **抽牌階段：** 由 P1 開始，P1~P4 輪流從牌堆中抽牌 (Draw Card)，直到所有人都擁有手牌 (Hand) 5 張牌為止。
        3. **遊戲執行以下流程：**
            1. 從牌堆中翻出第一張牌到檯面上。
            2. 由 P1 開始，出牌順序為 P1 → P2 → P3 → P4 → P1 以此類推。
            3. 玩家出的牌必須與檯面上最新的牌的顏色一樣，或是數字一樣。出完的牌就會成為檯面上最新的牌。
            4. 最快出完手中牌的人為遊戲的贏家。
            5. 如果玩家沒有任何可出的牌，玩家就必須從牌堆中抽一張牌，如果此時牌堆空了，則會先把檯面上除了最新的牌以外的牌放回牌堆中進行洗牌。