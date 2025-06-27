package guild

import (
	"fmt"
	"hero-many-to-many/hero"
)

type Guild struct {
	Name    string
	Members map[hero.Hero]struct{} // 使用 map[Hero]struct{} 作為 set
}

func NewGuild(name string, members map[hero.Hero]struct{}) (*Guild, error) {
	if name == "" {
		return nil, fmt.Errorf("guild name cannot be empty")
	}

	if len(members) > 10 || len(members) < 1 {
		return nil, fmt.Errorf("num of the members must be within 1-10")
	}

	// 初始化 Guild 的 Members map
	g := &Guild{
		Name:    name,
		Members: make(map[hero.Hero]struct{}, len(members)), // 初始化 map，預分配空間
	}

	// 將傳入的 members 複製到新的 map 中
	// 避免直接使用外部傳入的 map，防止外部修改影響 Guild 內部的狀態
	for member := range members {
		g.Members[member] = struct{}{}
	}

	return g, nil
}

func (g *Guild) GetMembers() []hero.Hero {
	members := make([]hero.Hero, 0, len(g.Members))
	for member := range g.Members {
		members = append(members, member)
	}
	return members
}

func (g *Guild) Join(member hero.Hero) error {
	if len(g.Members) == 10 {
		return fmt.Errorf("cannot join the guild since it's already had 10 members")
	}

	if _, exists := g.Members[member]; exists {
		return fmt.Errorf("cannot join the guild twice")
	}

	g.Members[member] = struct{}{}
	return nil
}

func (g *Guild) Leave(member hero.Hero) error {
	if _, exists := g.Members[member]; !exists {
		return fmt.Errorf("only the member of the guild can leave")
	}

	if len(g.Members) == 1 {
		return fmt.Errorf("the only member of the guild cannot leave")
	}

	delete(g.Members, member)
	return nil
}

func (g Guild) GetName() string {
	return g.Name
}

func (g *Guild) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("cannot assign the empty name for a guild")
	}
	g.Name = name
	return nil
}
