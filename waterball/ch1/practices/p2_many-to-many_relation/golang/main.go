package main

import (
	"fmt"
	"hero-many-to-many/guild"
	"hero-many-to-many/hero"
)

func main() {
	a := hero.NewHero("here-a")
	b := hero.NewHero("here-b")
	c := hero.NewHero("here-c")

	// 創建初始 Hero map
	members := map[hero.Hero]struct{}{
		*a: {},
	}

	// 創建公會
	g, err := guild.NewGuild("BaconYao", members)
	if err != nil {
		fmt.Println("Error creating guild:", err)
		return
	}

	fmt.Println("初始成員數量:", len(g.GetMembers()))

	// 英雄 b 加入公會
	if err := g.Join(*b); err != nil {
		fmt.Println("Error joining guild:", err)
	}
	fmt.Println("B 加入, 成員數量:", len(g.GetMembers()))

	// 英雄 c 加入公會
	if err := g.Join(*c); err != nil {
		fmt.Println("Error joining guild:", err)
	}
	fmt.Println("C 加入, 成員數量:", len(g.GetMembers()))

	// 英雄 A "再次" 加入公會
	if err := g.Join(*a); err != nil {
		fmt.Println("Error:", err)
	}

	// 英雄 A 離開公會
	if err := g.Leave(*a); err != nil {
		fmt.Println("Error leaving guild:", err)
	}
	fmt.Println("A 離開公會, 成員數量:", len(g.GetMembers()))

	// 英雄 B 離開公會
	if err := g.Leave(*b); err != nil {
		fmt.Println("Error leaving guild:", err)
	}
	fmt.Println("B 離開公會, 成員數量:", len(g.GetMembers()))

	// 英雄 B "再次" 離開公會
	if err := g.Leave(*b); err != nil {
		fmt.Println("Error leaving guild:", err)
	}

	// 英雄 C 離開公會
	if err := g.Leave(*c); err != nil {
		fmt.Println("Error leaving guild:", err)
	}
}
