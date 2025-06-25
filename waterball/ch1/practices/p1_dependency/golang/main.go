package main

import (
	"fmt"
	"hero-dependency/hero"
)

func main() {
	myHero := hero.NewHero()
	levelSheet := hero.NewLevelSheet()

	// Test gaining experience
	exps := []int{0, 100, 900, -200}
	for _, exp := range exps {
		err := myHero.GainExp(exp, *levelSheet)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
