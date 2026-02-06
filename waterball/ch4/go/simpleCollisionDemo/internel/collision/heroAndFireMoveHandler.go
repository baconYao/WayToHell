package collision

import (
	"fmt"
	"simpleCollisionDemo/internel/sprite"
)

type HeroFireHandler struct{ BaseCollisionHandler }

func NewHeroFireHandler() CollisionHandler {
	return &HeroFireHandler{}
}

func (h *HeroFireHandler) Handle(req *CollisionRequest) bool {
	match := func(r *CollisionRequest) bool {
		return (r.C1.GetID() == sprite.HeroID && r.C2.GetID() == sprite.FireID) || (r.C1.GetID() == sprite.FireID && r.C2.GetID() == sprite.HeroID)
	}
	action := func(r *CollisionRequest) {
		fmt.Println(">> 效果：英雄與火相遇，英雄生命值-10，火被移除")
		if r.C1.GetID() == sprite.HeroID {
			hero := r.C1.(*sprite.HeroSprite)
			hero.SetHP(hero.GetHP() - 10)
			_ = r.World.RemoveSprite(r.TO)
			if hero.GetHP() <= 0 {
				fmt.Println("英雄死亡")
				_ = r.World.RemoveSprite(r.FROM)
				return
			}
			fmt.Println("英雄移動到火的位置")
			r.World.MoveSprite(r.FROM, r.TO)
		} else {
			hero := r.C2.(*sprite.HeroSprite)
			hero.SetHP(hero.GetHP() + 10)
			_ = r.World.RemoveSprite(r.FROM)
			if hero.GetHP() <= 0 {
				fmt.Println("英雄死亡")
				_ = r.World.RemoveSprite(r.TO)
			}
		}
	}
	return h.BaseCollisionHandler.moveTemplate(req, match, action)
}
