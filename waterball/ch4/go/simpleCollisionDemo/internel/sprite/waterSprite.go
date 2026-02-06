package sprite

const WaterID byte = 'W'

type WaterSprite struct {
	Sprite
}

func NewWaterSprite() *WaterSprite {
	return &WaterSprite{Sprite: *NewSprite(WaterID)}
}

func (w *WaterSprite) GetID() byte {
	return w.Sprite.GetID()
}
