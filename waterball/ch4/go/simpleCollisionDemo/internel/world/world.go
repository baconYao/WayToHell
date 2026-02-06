package world

import (
	"errors"
	"fmt"
	"math/rand"
	"simpleCollisionDemo/internel/collision"
	"simpleCollisionDemo/internel/sprite"
)

const (
	WorldSize   = 30
	SpriteCount = 10
)

type World struct {
	grid             []sprite.SpriteInterface
	collisionHandler collision.CollisionHandler
}

// NewWorld creates a world with the given move handler chain (can be nil for tests).
func NewWorld(collisionHandler collision.CollisionHandler) *World {
	return &World{
		grid:             make([]sprite.SpriteInterface, WorldSize),
		collisionHandler: collisionHandler,
	}
}

// InitializeWithRandomSprites places 10 sprites (x Hero, y Water, z Fire, x+y+z=10)
// at 10 distinct random positions in [0, WorldSize-1].
func (w *World) InitializeWithRandomSprites() {
	x := rand.Intn(SpriteCount + 1)
	y := rand.Intn(SpriteCount - x + 1)
	z := SpriteCount - x - y

	sprites := make([]sprite.SpriteInterface, 0, SpriteCount)
	for i := 0; i < x; i++ {
		sprites = append(sprites, sprite.NewHeroSprite())
	}
	for i := 0; i < y; i++ {
		sprites = append(sprites, sprite.NewWaterSprite())
	}
	for i := 0; i < z; i++ {
		sprites = append(sprites, sprite.NewFireSprite())
	}

	rand.Shuffle(len(sprites), func(i, j int) {
		sprites[i], sprites[j] = sprites[j], sprites[i]
	})

	positions := w.distinctRandomIndices(WorldSize, SpriteCount)
	for i, pos := range positions {
		if err := w.addSpriteToGrid(pos, sprites[i]); err != nil {
			panic(fmt.Sprintf("addSprite at %d: %v", pos, err))
		}
	}
}

func (w *World) distinctRandomIndices(max, n int) []int {
	if n > max {
		panic("n must be <= max")
	}
	idx := make([]int, max)
	for i := 0; i < max; i++ {
		idx[i] = i
	}
	rand.Shuffle(max, func(i, j int) {
		idx[i], idx[j] = idx[j], idx[i]
	})
	return idx[:n]
}

// Handle delegates to the collision handler chain.
func (w *World) HandleMove(from, to int) {
	// 1. 基本檢查：邊界與來源是否存在
	if from < 0 || from >= WorldSize || to < 0 || to >= WorldSize {
		fmt.Println(">> 錯誤：位置超出世界邊界。")
		return
	}

	movingSprite := w.GetSprite(from)
	if movingSprite == nil {
		fmt.Println(">> 錯誤：起點位置沒有生命體。")
		return
	}

	// 2. 碰撞偵測
	targetSprite := w.GetSprite(to)
	if targetSprite != nil {
		// 觸發碰撞處理程序
		req := &collision.CollisionRequest{
			C1: w.grid[from], C2: w.grid[to],
			World: w, FROM: from, TO: to,
		}
		if !w.collisionHandler.Handle(req) { // 委派給碰撞處理器鏈
			fmt.Println(">> 無法處理的碰撞，移動拒絕。")
		}
	} else {
		// 沒有阻礙，直接移動
		w.MoveSprite(from, to)
		fmt.Printf(">> 移動成功：%c 從 %d 移到了 %d\n", movingSprite.GetID(), from, to)
	}
}

func (w *World) addSpriteToGrid(index int, s sprite.SpriteInterface) error {
	if index < 0 || index >= len(w.grid) {
		return errors.New("index out of range")
	}
	if w.grid[index] != nil {
		return errors.New("index already has a sprite")
	}
	w.grid[index] = s
	return nil
}

func (w *World) RemoveSprite(index int) error {
	if index < 0 || index >= len(w.grid) {
		return errors.New("index out of range")
	}
	if w.grid[index] != nil {
		w.grid[index] = nil
	}
	return nil
}

func (w *World) GetSprite(index int) sprite.SpriteInterface {
	if index < 0 || index >= len(w.grid) {
		return nil
	}
	return w.grid[index]
}

func (w *World) Getgrid() []sprite.SpriteInterface {
	return w.grid
}

func (w *World) MoveSprite(from int, to int) {
	w.grid[to] = w.grid[from]
	w.grid[from] = nil
}

func (w *World) ShowWorld() {
	fmt.Println("World: [")
	for i, s := range w.grid {
		if s == nil {
			fmt.Printf("(%d: -),", i)
			continue
		}
		fmt.Printf("(%d: %c),", i, s.GetID())
	}
	fmt.Println("\n]")
}
