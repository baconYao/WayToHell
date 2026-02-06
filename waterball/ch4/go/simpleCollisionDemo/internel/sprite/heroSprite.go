package sprite

const HeroID byte = 'H'

type HeroSprite struct {
	Sprite
	HP int
}

func NewHeroSprite() *HeroSprite {
	return &HeroSprite{Sprite: *NewSprite(HeroID), HP: 30}
}

func (h *HeroSprite) GetHP() int {
	return h.HP
}

func (h *HeroSprite) SetHP(hp int) {
	h.HP = hp
}

func (h *HeroSprite) GetID() byte {
	return h.Sprite.GetID()
}
