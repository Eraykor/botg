package main

import (
	"fmt"
	"io"
)

type Item struct {
	Name              string
	Cost              int
	Damage            int
	Health, MaxHealth int
	Mana, MaxMana     int
	MoveSpeed         int
	ManaRegeneration  int
	IsPotion          bool
}

func NewItem(r io.Reader) *Item {
	item := &Item{}

	fmt.Fscanln(r,
		&item.Name,
		&item.Cost,
		&item.Damage,
		&item.Health,
		&item.MaxHealth,
		&item.Mana,
		&item.MaxMana,
		&item.MoveSpeed,
		&item.ManaRegeneration,
		&item.IsPotion,
	)

	return item
}

func (item *Item) String() string {
	s := fmt.Sprintf("Name: %s, Cost: %d", item.Name, item.Cost)
	if item.Damage > 0 {
		s += fmt.Sprintf(", %d", item.Damage)
	}
	if item.Health > 0 {
		s += fmt.Sprintf(", %d", item.Health)
	}
	if item.MaxHealth > 0 {
		s += fmt.Sprintf(", %d", item.MaxHealth)
	}
	if item.Mana > 0 {
		s += fmt.Sprintf(", %d", item.Mana)
	}
	if item.MaxMana > 0 {
		s += fmt.Sprintf(", %d", item.MaxMana)
	}
	if item.MoveSpeed > 0 {
		s += fmt.Sprintf(", %d", item.MoveSpeed)
	}
	if item.ManaRegeneration > 0 {
		s += fmt.Sprintf(", %d", item.ManaRegeneration)
	}
	if item.IsPotion {
		s += ", IsPotion"
	}

	return s
}
