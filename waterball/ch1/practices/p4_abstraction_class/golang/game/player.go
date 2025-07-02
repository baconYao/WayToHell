package game

type Player interface {
	Decide() Decision
	GetNumber() int
}

type player struct {
	number int
}

func (p player) GetNumber() int {
	return p.number
}

func (p player) Decide() Decision {
	panic("Decide method must be implemented by concrete types")
}
