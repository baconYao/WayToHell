package sprite

// SpriteInterface is implemented by all sprite types so the world and move handlers can store any sprite.
type SpriteInterface interface {
	GetID() byte
}

type Sprite struct {
	ID byte
}

func NewSprite(id byte) *Sprite {
	return &Sprite{ID: id}
}

func (s *Sprite) GetID() byte {
	return s.ID
}
