package main

import (
	"fmt"
	"io"
)

type RoundState struct {
	Gold      int
	EnemyGold int
	Type      int

	Allies   []*Entity
	Enemies  []*Entity
	Neutrals []*Entity
	Heroes   []*Hero
	Tower    *Entity
}

func NewRoundState(r io.Reader, gs *GameState) *RoundState {
	rs := &RoundState{}

	fmt.Fscan(r,
		&rs.Gold,
		&rs.EnemyGold,
		&rs.Type,
	)

	var entityCount int
	fmt.Fscan(r, &entityCount)

	for i := 0; i < entityCount; i++ {
		entity := NewEntity(r)

		if entity.Team == gs.MyTeam {
			if entity.UnitType == "HERO" {
				rs.Heroes = append(rs.Heroes, NewHero(entity))
			} else {
				rs.Allies = append(rs.Allies, entity)
			}
			if entity.UnitType == "TOWER" {
				rs.Tower = entity
			}
		} else if entity.Team == -1 {
			rs.Neutrals = append(rs.Neutrals, entity)
		} else {
			rs.Enemies = append(rs.Enemies, entity)
		}
	}

	return rs
}
