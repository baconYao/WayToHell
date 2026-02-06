package collision

import (
	"fmt"
	"simpleCollisionDemo/internel/sprite"
)

type WaterWaterHandler struct{ BaseCollisionHandler }

func NewWaterWaterHandler() CollisionHandler {
	return &WaterWaterHandler{}
}

func (h *WaterWaterHandler) Handle(req *CollisionRequest) bool {
	match := func(r *CollisionRequest) bool {
		return (r.C1.GetID() == sprite.WaterID && r.C2.GetID() == sprite.WaterID)
	}
	action := func(r *CollisionRequest) {
		fmt.Println(">> 效果：水水相遇，沒事發生。")
	}
	return h.BaseCollisionHandler.moveTemplate(req, match, action)
}
