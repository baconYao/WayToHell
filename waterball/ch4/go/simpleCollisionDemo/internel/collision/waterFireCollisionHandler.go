package collision

import (
	"fmt"
	"simpleCollisionDemo/internel/sprite"
)

type WaterFireHandler struct{ BaseCollisionHandler }

func NewWaterFireHandler() CollisionHandler {
	return &WaterFireHandler{}
}

func (h *WaterFireHandler) Handle(req *CollisionRequest) bool {
	match := func(r *CollisionRequest) bool {
		return (r.C1.GetID() == sprite.WaterID && r.C2.GetID() == sprite.FireID) || (r.C1.GetID() == sprite.FireID && r.C2.GetID() == sprite.WaterID)
	}
	action := func(r *CollisionRequest) {
		fmt.Println(">> 效果：水火相剋，兩者皆消失。")
		_ = r.World.RemoveSprite(r.FROM)
		_ = r.World.RemoveSprite(r.TO)
	}
	return h.BaseCollisionHandler.moveTemplate(req, match, action)
}
