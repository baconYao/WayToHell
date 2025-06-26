package hero

import (
	"fmt"
	"strings"
)

type Pet struct {
	name  string
	owner *Hero
}

func NewPet(name string) *Pet {
	return &Pet{name: name, owner: nil}
}

// eatFruit helps owner gain 10 HP each time
func (p Pet) EatFruit(fruit Fruit) {
	fmt.Println("寵物吃水果...")
	if p.owner != nil {
		p.owner.setHp(p.owner.GetHp() + 10)
	}
}

func (p Pet) getOwner() *Hero {
	return p.owner
}

func (p *Pet) setOwner(owner *Hero) {
	p.owner = owner
}

func (p Pet) GetName() string {
	return p.name
}

func (p *Pet) setName(name string) error {
	if strings.TrimSpace(name) != "" {
		return fmt.Errorf("pet's name must be string, got %s", name)
	}
	return nil
}
