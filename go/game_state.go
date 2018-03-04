package main

import (
	"fmt"
	"io"
)

type GameState struct {
	MyTeam  int
	Items   []*Item
	Potions []*Item
}

func NewGameState(r io.Reader) *GameState {
	gs := &GameState{}

	fmt.Fscan(r, &gs.MyTeam)

	var bushAndSpawnCount int
	fmt.Fscan(r, &bushAndSpawnCount)

	for i := 0; i < bushAndSpawnCount; i++ {
		var entityType string
		var x, y, radius int

		fmt.Fscanln(r, &entityType, &x, &y, &radius)
	}

	var itemCount int
	fmt.Fscan(r, &itemCount)

	for i := 0; i < itemCount; i++ {
		item := NewItem(r)

		if item.IsPotion {
			gs.Potions = append(gs.Potions, item)
		} else {
			gs.Items = append(gs.Items, item)
		}
	}

	return gs
}
