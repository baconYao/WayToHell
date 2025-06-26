package main

import (
	"fmt"
	"hero-one-to-one/fruit"
	"hero-one-to-one/hero"
	"hero-one-to-one/pet"
)

func main() {
	h := hero.NewHero()
	pet := pet.NewPet("Cat")
	h.SetPet(pet)

	fmt.Printf("Hero 目前血量: %d\n", h.GetHp())
	fmt.Printf("Hero's 寵物名稱: %s\n", h.GetPet().GetName())

	for i := 0; i < 5; i++ {
		pet.EatFruit(fruit.NewFruit())
	}

	h.RemovePet()

	for i := 0; i < 5; i++ {
		pet.EatFruit(fruit.NewFruit())
	}

}
