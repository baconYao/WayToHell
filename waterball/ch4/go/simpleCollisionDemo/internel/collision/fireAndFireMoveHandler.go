package collision

import (
	"fmt"
	"simpleCollisionDemo/internel/sprite"
)

type FireFireHandler struct{ BaseCollisionHandler }

func NewFireFireHandler() CollisionHandler {
	return &FireFireHandler{}
}

func (h *FireFireHandler) Handle(req *CollisionRequest) bool {
	match := func(r *CollisionRequest) bool {
		return (r.C1.GetID() == sprite.FireID && r.C2.GetID() == sprite.FireID)
	}
	action := func(r *CollisionRequest) {
		fmt.Println(">> 效果：火火相遇，沒事發生。")
	}
	return h.BaseCollisionHandler.moveTemplate(req, match, action)
}
