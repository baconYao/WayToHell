package hero

import (
	"fmt"
	"hero-one-to-one/pet"
)

// Hero represents a character with level, experience, health points, and an optional pet.
type Hero struct {
	level    int
	totalExp int
	hp       int
	pet      *pet.Pet
}

// NewHero creates a new Hero with default values.
func NewHero() *Hero {
	return &Hero{
		level:    1,
		totalExp: 0,
		hp:       100,
		pet:      nil,
	}
}

// GainExp adds experience points and updates the level using the provided LevelSheet.
func (h *Hero) GainExp(exp int, levelSheet LevelSheet) error {
	if exp < 0 {
		return fmt.Errorf("error: 無法獲得負數的經驗值 %d", exp)
	}
	currentLevel := h.getLevel()
	h.setTotalExp(h.getTotalExp() + exp)
	h.setLevel(levelSheet.QueryLevel(h.getTotalExp()))
	fmt.Printf("英雄目前等級 %d，獲得 %d EXP，最新總共經驗值為 %d，最新等級為 %d。\n", currentLevel, exp, h.getTotalExp(), h.getLevel())
	return nil
}

// GetPet returns the hero's pet, which may be nil.
func (h *Hero) GetPet() *pet.Pet {
	return h.pet
}

// SetPet sets the hero's pet and updates the pet's owner.
func (h *Hero) SetPet(pet *pet.Pet) {
	if h.pet != nil {
		h.pet.SetOwner(nil)
	}
	h.pet = pet
	if pet != nil {
		pet.SetOwner(h)
	}
}

// RemovePet clears the hero's pet and its owner reference.
func (h *Hero) RemovePet() {
	if h.pet != nil {
		h.pet.SetOwner(nil)
	}
	h.pet = nil
}

// setTotalExp sets the total experience points, ensuring it's non-negative.
func (h *Hero) setTotalExp(totalExp int) error {
	if totalExp < 0 {
		return fmt.Errorf("TotalExp must be greater than or equal to 0, got %d", totalExp)
	}
	h.totalExp = totalExp
	return nil
}

// setLevel sets the level, ensuring it's non-negative.
func (h *Hero) setLevel(level int) error {
	if level < 0 {
		return fmt.Errorf("level must be greater than or equal to 0, got %d", level)
	}
	h.level = level
	return nil
}

// SetHp sets the hp, ensuring it's non-negative.
func (h *Hero) SetHp(hp int) error {
	if hp < 0 {
		return fmt.Errorf("HP must be greater than or equal to 0, got %d", hp)
	}
	h.hp = hp
	fmt.Println("英雄血量更新至", hp)
	return nil
}

// getTotalExp returns the total experience points.
func (h *Hero) getTotalExp() int {
	return h.totalExp
}

// getLevel returns the hero's current level.
func (h *Hero) getLevel() int {
	return h.level
}

// GetHp returns the hero's health points.
func (h *Hero) GetHp() int {
	return h.hp
}
