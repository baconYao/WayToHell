package collision

import (
	"simpleCollisionDemo/internel/sprite"
)

// WorldMap 介面：定義碰撞處理器需要的功能，避免循環引用
type WorldMap interface {
	RemoveSprite(index int) error
	MoveSprite(from, to int)
}

type CollisionRequest struct {
	C1, C2   sprite.SpriteInterface
	World    WorldMap
	FROM, TO int // 座標
}

type CollisionHandler interface {
	Handle(req *CollisionRequest) bool
	SetNext(CollisionHandler)
}

type BaseCollisionHandler struct {
	next CollisionHandler
}

func (b *BaseCollisionHandler) SetNext(next CollisionHandler) { b.next = next }

// moveTemplate 定義了固定的模板骨架
func (b *BaseCollisionHandler) moveTemplate(req *CollisionRequest,
	match func(*CollisionRequest) bool,
	collisionHandler func(*CollisionRequest)) bool {

	if match(req) {
		collisionHandler(req)
		return true // 處理完成
	}

	if b.next != nil {
		return b.next.Handle(req)
	}
	return false // 鏈條結尾，無人能處理
}
