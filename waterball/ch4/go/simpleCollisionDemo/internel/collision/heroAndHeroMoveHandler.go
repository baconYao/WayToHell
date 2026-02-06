package collision

import (
	"fmt"
	"simpleCollisionDemo/internel/sprite"
)

type HeroHeroHandler struct{ BaseCollisionHandler }

func NewHeroHeroHandler() CollisionHandler {
	return &HeroHeroHandler{}
}

func (h *HeroHeroHandler) Handle(req *CollisionRequest) bool {
	match := func(r *CollisionRequest) bool {
		return (r.C1.GetID() == sprite.HeroID && r.C2.GetID() == sprite.HeroID)
	}
	action := func(r *CollisionRequest) {
		fmt.Println(">> 效果：英雄與英雄相遇，沒事發生")
	}
	return h.BaseCollisionHandler.moveTemplate(req, match, action)
}
