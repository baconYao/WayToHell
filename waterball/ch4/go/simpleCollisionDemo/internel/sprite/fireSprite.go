package sprite

const FireID byte = 'F'

type FireSprite struct {
	Sprite
}

func NewFireSprite() *FireSprite {
	return &FireSprite{Sprite: *NewSprite(FireID)}
}

func (f *FireSprite) GetID() byte {
	return f.Sprite.GetID()
}
