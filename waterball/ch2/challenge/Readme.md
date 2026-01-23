# Matchmaking System

## How to play?

```go
// Run Distance Based Startegy
$ make run
// Run Habit Based Startegy
$ make run MATCH_STRATEGY=habit

// Debug
$ make dev
$ make dev MATCH_STRATEGY=habit
```

## Requirements Specification



### You are given a large collection of Individual data. Each record contains the following fields:

- ID
    - A positive integer (> 0).
    - Each individual has a unique ID.

- Gender
    - MALE represents male.
    - FEMALE represents female.

- Age

    - A positive integer.
    - Must be at least 18 years old.

- Introduction (Intro)
    - Text with a length of approximately 0–200 characters.

- Habits
    - An individual may have multiple habits.
    - Each habit is represented by a string of length 1–10 characters.
    - Habits are separated by commas.
    - Example: Basketball, Cooking, Gaming

- Coordinates (Coord)
    - Represented as (x, y), indicating the individual’s position on the x-axis and y-axis.

### Matchmaking System

You are required to implement a Matchmaking System.
The system matches each user with the most suitable individual.

The system initially provides two matchmaking strategies, and more strategies will be added in the future:

- Distance-Based Strategy
    - Match with the individual who is closest in distance.
    - If distances are equal, choose the individual with the smaller ID.
    - Given the current user’s coordinates (x, y) and another individual’s coordinates (x′, y′), the distance is calculated as: (*y*−*y*′)2＋(*x*−*x*′)2
- Habit-Based Strategy
    - Match with the individual who has the largest number of shared habits.
    - If the number of shared habits is equal, choose the individual with the smaller ID.
- Reverse Strategies
    - Some users prefer to meet more diverse people. Therefore, for each strategy, the system must also provide a reverse implementation:
    - The reverse of Distance-Based matches the individual who is farthest away.
    - The reverse of Habit-Based matches the individual with the smallest number of shared habits.

### Design Requirements (Non-Functional Requirements)

Since the client will continuously develop new matchmaking strategies, the system must allow the Client (the system user in code, typically the main method) to swap matchmaking strategies externally.

### Advanced Challenge

Can you implement Requirement (3) — Reverse Strategies — while satisfying the following constraints?

#### No Direct Coupling to “Reverse”

The following classes must not directly couple to the concept of “Reverse”:

- Matchmaking System class
- Distance-Based strategy
- Habit-Based strategy

Additionally:

- You must not use boolean fields (e.g., reverse: boolean)
- You must not use boolean parameters (e.g., match(..., reverse: boolean))

The following are forbidden implementations:

- `new MatchmakingSystem(new DistanceBasedMatchmakingStrategy(), reverse = true)` ❌
- `matchMakingSystem.match(individual, reverse = true)` ❌
- `distanceBasedStrategy.reverse()` ❌
- `new DistanceBasedStrategy(reverse = true)` ❌
- `matchMakingSystem.reverse()` ❌

#### No Combinatorial Explosion

You must not introduce four separate classes representing:

- Distance-Based (normal)
- Distance-Based (reverse)
- Habit-Based (normal)
- Habit-Based (reverse)

If your design contains four matchmaking strategy classes, it is considered a combinatorial explosion and is not allowed.


## Requirements Specification (Chinese Version)

你正在開發一個交友配對系統 (Matchmaking System)。

1. 你手邊有一大群對象 (Individual) 資料，每一筆資料皆記載著以下欄位：
    1. 編號 (ID)：正整數 (>0)，每位對象的編號都不同。
    2. 性別 (Gender)：`MALE` 表示男生、`FEMALE` 表示女生
    3. 年紀 (Age)：正整數；至少 18 歲。
    4. 自我介紹 (Intro)：長度約 0~200 的文字。
    5. 興趣 (Habits)：可以為多個興趣，每個興趣以 1~10 長度的文字表示，每個興趣之間以一個逗號隔開。例如：`打籃球, 煮菜, 玩遊戲`
    6. 座標 (Coord)：以 `(x, y)` 表示該用戶所在的 x 軸和 y 軸的位置。
2. 你要撰寫一個配對系統 (Matchmaking System)，**系統會幫每位用戶配對最適合他的用戶，** 而你的系統主要提供了兩種配對策略，並且未來會持續新增新的策略：
    1. **距離先決 (Distance-Based)：** 配對與自己距離最近的對象（距離相同則選擇編號較小的那位）。
        1. 假設自己的座標為 (*x*,*y*) 而對象的座標為(*x*′,*y*′)，則距離公式為：(*y*−*y*′)2＋(*x*−*x*′)2
    2. 興趣**先決 (Habit-Based)：** 配對與自己興趣擁有最大交集量的對象（興趣交集量相同則選擇編號較小的那位）。
3. 不過，也有用戶喜歡認識更多元的人，因此你的系統還要為你每個策略提供「**反向 (reverse)**」的實作版本，像是把距離先決的策略改成是配對與自己距離最遠的對象，而把興趣先決改成是配對與自己興趣擁有**最小交集量**的對象。

### **設計需求（非功能性需求）**

由於客戶會不斷研發新的配對策略，我們希望系統能夠允許 Client (程式中系統的使用方，通常會是你的 **Main method**）能在外部抽換系統的配對策略。

### **進階挑戰題**

你想得到如何遵守以下條件來來實踐需求 (3) —— 反向策略嗎？

1. 「配對系統類別/距離先決策略/興趣先決」類別不能直接耦合「反向 (Reverse)」的概念，也不能使用 boolean 欄位 (e.g., `reverse: boolean`) 或是參數 (e.g., `match(…, …, reverse: boolean)`) 來區分正反向。
    1. 以下這些都是被禁止的錯誤示範：
        1. `new MatchmakingSystem(new DistanceBasedMatchmakingStrategy(), Reverse=True)` (X)
        2. `matchMakingSystem.match(individual, Reverse=True)` (X)
        3. `distanceBasedStrategy.reverse()` (X)
        4. `new DistanceBasedStrategy(Reverse=True)` (X)
        5. `matchMakingSystem.reverse()` (X)
2. 不組合爆炸：意即你不能讓程式中出現攸關於「正向距離先決、反向距離先決、正向興趣先決、反向興趣先決」概念的四個類別，只要你有四個配對策略的類別，就是組合爆炸了，是不被允許的。