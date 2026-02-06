package collision

import (
	"fmt"
	"simpleCollisionDemo/internel/sprite"
)

type HeroWaterHandler struct{ BaseCollisionHandler }

func NewHeroWaterHandler() CollisionHandler {
	return &HeroWaterHandler{}
}

func (h *HeroWaterHandler) Handle(req *CollisionRequest) bool {
	match := func(r *CollisionRequest) bool {
		return (r.C1.GetID() == sprite.HeroID && r.C2.GetID() == sprite.WaterID) || (r.C1.GetID() == sprite.WaterID && r.C2.GetID() == sprite.HeroID)
	}
	action := func(r *CollisionRequest) {
		fmt.Println(">> 效果：英雄與水相遇，英雄生命值增加10，水被移除")
		if r.C1.GetID() == sprite.HeroID {
			hero := r.C1.(*sprite.HeroSprite)
			hero.SetHP(hero.GetHP() + 10)
			_ = r.World.RemoveSprite(r.TO)
			fmt.Println("英雄移動到水的位置")
			r.World.MoveSprite(r.FROM, r.TO)
		} else {
			hero := r.C2.(*sprite.HeroSprite)
			hero.SetHP(hero.GetHP() + 10)
			_ = r.World.RemoveSprite(r.FROM)
		}
	}
	return h.BaseCollisionHandler.moveTemplate(req, match, action)
}
