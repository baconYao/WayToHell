# Showdown Pocker Game

## Requirements Specification

You have received a simplified requirements document for a card game called “Showdown” (high-card comparison game). Based on the description in this document, you are required to analyze it sentence by sentence, draw UML sequence diagrams and class diagrams, build the domain model of Showdown, and once you have about 80% understanding of the domain model, implement it using object-oriented programming (OOP) in the programming language you are most familiar with.

### Players

This game supports four players. A player can be either a Human Player or an AI Player.

Player implementations:

- Human Player: makes choices via a Command Line Interface (CLI).

- AI Player: makes choices randomly.

In the following requirements, P1, P2, P3, and P4 represent the first, second, third, and fourth players, respectively.

### Deck and Cards

The game contains a Deck.

The deck initially contains 52 Cards.

Each card has a Rank and a Suit.

### Game Flow

The game proceeds according to the following steps (1–4):

- Game Initialization

- P1–P4 take turns to name themselves.

- The deck is shuffled.

- Card Drawing Phase
    - Starting from P1, players P1–P4 take turns drawing cards from the deck until each player has 13 cards in hand.

- Playing Rounds
    - After all cards are drawn, the game continues for 13 rounds. In each round, the following steps are executed in order:

- P1–P4 take turns, and for each player:

    - Decide whether to use the “Exchange Hands” privilege (see Requirement 5).

    - Play (show) one card (at this step, players do not know what cards the others have played).

    - Reveal the cards played by P1–P4.

    - Compare the played cards to determine the winner, and award 1 point to the winning player.

- Game End
    - After 13 rounds, all players will have played all their cards and the game ends.
    - The player with the highest score is declared the winner, and the winner’s name is displayed.

### Card Comparison Rules

First compare the rank of the cards. The card with the higher rank wins.
If the ranks are equal, compare the suit; the card with the higher suit wins.

Rank order from lowest to highest:
2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K, A

Suit order from lowest to highest:
Club, Diamond, Heart, Spade

### Exchange Hands Privilege

In each round, before playing a card, a player may choose whether to use the “Exchange Hands” privilege.
Each player may use this privilege only once per game. If the player chooses to use it, the following steps are executed in order:

- The player selects another player (excluding themselves) to exchange hands with.

- The two players swap their hands.

- After three rounds, the two players’ hands are swapped back.

- If, after exchanging hands, a player finds that they have no cards to play, that player may skip playing a card and will not participate in the card comparison for that round.

## Requirements Specification (Chinese Version)

你收到了一份 **撲克牌比大小遊戲 (Showdown)** 的簡易需求文件， 你需要基於這份文件的描述，一句一句地分析，並用 UML 繪製循序圖、類別圖，建出 Showdown 的領域模型，對領域模型有個 80% 的掌握後，接著就透過將模型用你最熟悉的程式語言將它 OOP 實作吧！

1. 這款遊戲能支援四位玩家：玩家可以為*真實玩家 (Human Player)* 也能為*電腦玩家 (AI Player)*。
    1. 玩家的實作：
        1. 真實玩家：使用指令介面輸入 (Command Line Interface) 來做選擇。
        2. 電腦玩家：隨機做選擇。
    2. 在以下需求中：我們用 *P1, P2, P3, P4 來表示第一、二、三和第四順位的玩家。*
2. 遊戲中有一副*牌堆 (Deck)*。
    1. 牌堆中一開始存有 52 張*牌 (Card)*。
    2. 每張牌都會擁有*階級 (Rank)* 及*花色 (Suit)*。
3. 遊戲依照以下 1~4 的流程進行：
    1. 遊戲開始時，依序執行以下：
        1. 請 P1~P4 為自己取名 (Name himself)。
        2. 牌堆會進行洗牌 (Shuffle)。
    2. **抽牌階段**：由 P1 開始，P1~P4 輪流從牌堆中抽牌 (Draw Card)，直到所有人都擁有手牌 (Hand) 13 張牌為止。
    3. 抽完牌後，在接下來的 13 回合中，每一回合依序執行以下：
        1. P1~P4 輪流 (Takes a turn) 依序執行以下：
            1. 決定要不要使用 **「交換手牌 (Exchange Hands)」** 特權，參見需求 5。
            2. 出 (Show) 一張牌（此步驟彼此皆無法知曉彼此出的牌）。
        2. 顯示 P1~P4 各出的牌的內容。
        3. 將 P1~P4 出的牌進行*比大小決勝負*，將最勝者的分數(Point)加一。
    4. 13 回合後，P1~P4 皆已出完全部的牌，遊戲結束。取得最多分數的玩家為勝者，將勝者的名稱顯示出來。
4. **牌與牌之間的比大小決勝規則**：
    1. 先比較牌的階級，此時階級較大者勝，如果階級相同則比較花色，此時花色較大者勝。
    2. 階級由小到大依序為：2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K, A
    3. 花色由小到大依序為：梅花 (Club)、菱形 (Diamond)、愛心 (Heart)、黑桃 (Spade)
5. 每輪中，玩家在出牌前都能選擇要不要使用 *「交換手牌 (Exchange Hands)」* 特權（此特權每場遊戲各玩家皆**只能使用一次**），如果要的話，依序執行以下：
    1. 玩家選擇要與哪位玩家（自己以外）交換手牌。
    2. 選擇後，雙方的手牌交換。
    3. **三回合後，雙方的手牌會交換回來。**
    4. 如果有玩家在換牌之後，發現沒有牌出了，此時該玩家可以不必出牌，並在此輪無法參與比大小決勝。