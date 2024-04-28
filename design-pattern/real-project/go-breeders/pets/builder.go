package pets

import "errors"

// PetInterface defines the methods that we want our builder to have. These
// are used to set the fields in the Pet type, and to build the final product.
// Everything except the Build() function returns the type *Pet because we are
// going to implement the fluent interface
type PetInterface interface {
	SetSpecies(s string) *Pet
	SetBreed(s string) *Pet
	SetMinWeight(s int) *Pet
	SetMaxWeight(s int) *Pet
	SetWeight(s int) *Pet
	SetDescription(s string) *Pet
	SetLifeSpan(s int) *Pet
	SetGeographicOrigin(s string) *Pet
	SetColor(s string) *Pet
	SetAge(s int) *Pet
	SetAgeEstimated(s bool) *Pet
	Build() (*Pet, error)
}

func NewPetBuilder() *Pet {
	return &Pet{}
}

// SetSpecies sets the species for our pet, and returns a *Pet
func (p *Pet) SetSpecies(s string) *Pet {
	p.Species = s
	return p
}

// SetBreed sets the breed for our pet, and returns a *Pet
func (p *Pet) SetBreed(s string) *Pet {
	p.Breed = s
	return p
}

// SetMinWeight sets the minimum weight for our pet, and returns a *Pet
func (p *Pet) SetMinWeight(s int) *Pet {
	p.MinWeight = s
	return p
}

// SetMaxWeight sets the maximum weight for our pet, and returns a *Pet
func (p *Pet) SetMaxWeight(s int) *Pet {
	p.MaxWeight = s
	return p
}

// SetWeight sets the weight for our pet, and returns a *Pet
func (p *Pet) SetWeight(s int) *Pet {
	p.Weight = s
	return p
}

// SetDescription sets the description for our pet, and returns a *Pet
func (p *Pet) SetDescription(s string) *Pet {
	p.Description = s
	return p
}

// SetLifeSpan sets the lifespan for our pet, and returns a *Pet
func (p *Pet) SetLifeSpan(s int) *Pet {
	p.LifeSpan = s
	return p
}

// SetGeographicOrigin sets the geographic origin for our pet, and returns a *Pet
func (p *Pet) SetGeographicOrigin(s string) *Pet {
	p.GeographicOrigin = s
	return p
}

// SetColor sets the color for our pet, and returns a *Pet
func (p *Pet) SetColor(s string) *Pet {
	p.Color = s
	return p
}

// SetAge sets the age for our pet, and returns a *Pet
func (p *Pet) SetAge(s int) *Pet {
	p.Age = s
	return p
}

// SetAgeEstimated sets the estimated age for our pet, and returns a *Pet
func (p *Pet) SetAgeEstimated(s bool) *Pet {
	p.AgeEstimated = s
	return p
}

// Build uses tha various "Set" functions above to build a pet, using the
// fluent interface. The inclusion of this function makes this an example
// of the Builder pattern
func (p *Pet) Build() (*Pet, error) {
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("minimum weight must be less than maximum weight")
	}

	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2

	return p, nil
}
