package hero

import (
	"fmt"
)

// Hero represents a character with level, experience, health points, and an optional pet.
type Hero struct {
	Name     string
	Level    int
	TotalExp int
	Hp       int
}

// NewHero creates a new Hero with default values.
func NewHero(name string) *Hero {
	return &Hero{
		Name:     name,
		Level:    1,
		TotalExp: 0,
		Hp:       100,
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

// setTotalExp sets the total experience points, ensuring it's non-negative.
func (h *Hero) setTotalExp(totalExp int) error {
	if totalExp < 0 {
		return fmt.Errorf("TotalExp must be greater than or equal to 0, got %d", totalExp)
	}
	h.TotalExp = totalExp
	return nil
}

// setLevel sets the level, ensuring it's non-negative.
func (h *Hero) setLevel(level int) error {
	if level < 0 {
		return fmt.Errorf("level must be greater than or equal to 0, got %d", level)
	}
	h.Level = level
	return nil
}

// SetHp sets the hp, ensuring it's non-negative.
func (h *Hero) SetHp(hp int) error {
	if hp < 0 {
		return fmt.Errorf("HP must be greater than or equal to 0, got %d", hp)
	}
	h.Hp = hp
	fmt.Println("英雄血量更新至", hp)
	return nil
}

// GetName returns the name of hero
func (h Hero) GetName() string {
	return h.Name
}

// getTotalExp returns the total experience points
func (h Hero) getTotalExp() int {
	return h.TotalExp
}

// getLevel returns the hero's current level
func (h Hero) getLevel() int {
	return h.Level
}

// GetHp returns the hero's health points.
func (h Hero) GetHp() int {
	return h.Hp
}
