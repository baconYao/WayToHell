package pet

import (
	"fmt"
	"hero-one-to-one/fruit"
	"strings"
)

// Owner defines the interface for a pet's owner.
type Owner interface {
	SetHp(hp int) error
	GetHp() int
}

// Pet represents a hero's companion.
type Pet struct {
	name  string
	owner Owner
}

// NewPet creates a new Pet with the given name.
func NewPet(name string) *Pet {
	return &Pet{name: name, owner: nil}
}

// EatFruit helps owner gain 10 HP each time.
func (p *Pet) EatFruit(fruit fruit.Fruit) {
	fmt.Println("寵物吃水果...")
	if p.owner != nil {
		p.owner.SetHp(p.owner.GetHp() + 10)
	}
}

// GetOwner returns the pet's owner, which may be nil.
func (p *Pet) GetOwner() Owner {
	return p.owner
}

// SetOwner sets the pet's owner.
func (p *Pet) SetOwner(owner Owner) {
	p.owner = owner
}

// GetName returns the pet's name.
func (p *Pet) GetName() string {
	return p.name
}

// setName sets the pet's name, ensuring it's not empty.
func (p *Pet) setName(name string) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("pet's name must not be empty, got %q", name)
	}
	p.name = name
	return nil
}
