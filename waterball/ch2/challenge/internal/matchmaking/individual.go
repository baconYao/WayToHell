package matchmaking

import (
	"fmt"
	"strings"
)

// Gender defines the gender type
type Gender string

const (
	Male   Gender = "MALE"
	Female Gender = "FEMALE"
)

type Coord struct {
	X int
	Y int
}

func (c Coord) getX() int {
	return c.X
}

func (c Coord) getY() int {
	return c.X
}

// Individual represents a participant in the matchmaking system
type Individual struct {
	ID      int      // Positive integer, unique
	Gender  Gender   // MALE or FEMALE
	Age     int      // Must be >= 18
	Intro   string   // 0-200 characters
	Habits  []string // List of habits
	Coord   Coord
	Reverse bool
}

func NewIndividual(id int, gender Gender, age int, intro string, habits string, x, y int, reverse bool) (*Individual, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be positive")
	}
	if gender != Male && gender != Female {
		return nil, fmt.Errorf("invalid gender")
	}
	if age < 18 {
		return nil, fmt.Errorf("age must be at least 18")
	}
	if len(intro) > 200 {
		return nil, fmt.Errorf("intro must be 200 characters or less")
	}
	habitList := strings.Split(habits, ",")
	for i, h := range habitList {
		habitList[i] = strings.TrimSpace(h)
		if len(habitList[i]) < 1 || len(habitList[i]) > 10 {
			return nil, fmt.Errorf("each habit must be 1-10 characters")
		}
	}

	return &Individual{
		ID:      id,
		Gender:  gender,
		Age:     age,
		Intro:   intro,
		Habits:  habitList,
		Coord:   Coord{X: x, Y: y},
		Reverse: reverse,
	}, nil
}

func (i Individual) getID() int {
	return i.ID
}

func (i Individual) getGender() Gender {
	return i.Gender
}

func (i Individual) getAge() int {
	return i.Age
}

func (i Individual) getIntro() string {
	return i.Intro
}

func (i Individual) getHabits() []string {
	return i.Habits
}

func (i Individual) getCoord() Coord {
	return i.Coord
}

func (i Individual) isUsingReverse() bool {
	return i.Reverse
}
