package main

import (
	"fmt"
	"simpleCollisionDemo/internel/collision"
	"simpleCollisionDemo/internel/world"
	"simpleCollisionDemo/pkg/cli"
)

func main() {
	// Chain of Responsibility: head → ... → tail (nil)
	wf := collision.NewWaterFireHandler()
	ww := collision.NewWaterWaterHandler()
	ff := collision.NewFireFireHandler()
	hf := collision.NewHeroFireHandler()
	hw := collision.NewHeroWaterHandler()
	hh := collision.NewHeroHeroHandler()

	wf.SetNext(ww)
	ww.SetNext(ff)
	ff.SetNext(hf)
	hf.SetNext(hw)
	hw.SetNext(hh)
	hh.SetNext(nil)

	w := world.NewWorld(wf)
	w.InitializeWithRandomSprites()

	fmt.Printf("World initialized with 10 sprites (world size %d)\n", world.WorldSize)
	maxIndex := world.WorldSize - 1

	for {
		fmt.Println("--------------------------------")
		w.ShowWorld()
		from, to, err := cli.ReadMove(maxIndex)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		w.HandleMove(from, to)
	}
}
