package main

import (
	"fmt"
	"io"
	"math"
)

type Entity struct {
	UnitID        int
	Team          int
	UnitType      string
	X             int
	Y             int
	AttackRange   int
	Health        int
	MaxHealth     int
	Shield        int
	AttackDamage  int
	MovementSpeed int
	StunDuration  int
	GoldValue     int

	// HERO
	Countdown1, Countdown2, Countdown3 int
	Mana                               int
	MaxMana                            int
	ManaRegeneration                   int
	HeroType                           string
	IsVisible                          int
	ItemsOwned                         int

	// TMP
	DistanceToHero  float64
	DistanceToTower float64
}

func NewEntity(r io.Reader) *Entity {
	entity := &Entity{}

	fmt.Fscan(r,
		&entity.UnitID,
		&entity.Team,
		&entity.UnitType,
		&entity.X,
		&entity.Y,
		&entity.AttackRange,
		&entity.Health,
		&entity.MaxHealth,
		&entity.Shield,
		&entity.AttackDamage,
		&entity.MovementSpeed,
		&entity.StunDuration,
		&entity.GoldValue,
		&entity.Countdown1,
		&entity.Countdown2,
		&entity.Countdown3,
		&entity.Mana,
		&entity.MaxMana,
		&entity.ManaRegeneration,
		&entity.HeroType,
		&entity.IsVisible,
		&entity.ItemsOwned,
	)

	return entity
}

func (entity *Entity) String() string {
	return fmt.Sprintf("ID: %d, Type: %s, Team: %d, X: %d, Y: %d", entity.UnitID, entity.UnitType, entity.Team, entity.X, entity.Y)
}

func (e *Entity) DistanceTo(other *Entity) float64 {
	return math.Abs(math.Sqrt(math.Pow((float64)(e.X-other.X), 2) + math.Pow((float64)(e.Y-other.Y), 2)))
}

func (e *Entity) Attack(unitID int) {
	fmt.Println("ATTACK", unitID)
}

func (e *Entity) Move(distanceX, distanceY float64) {
	if e.Team == 0 {
		fmt.Println("MOVE", (float64)(e.X)+distanceX, (float64)(e.Y)+distanceY)
	} else {
		fmt.Println("MOVE", (float64)(e.X)-distanceX, (float64)(e.Y)-distanceY)
	}
}

func (e *Entity) MoveAndAttack(distanceX, distanceY float64, unitID int) {
	if e.Team == 0 {
		fmt.Println("MOVE_ATTACK", (float64)(e.X)+distanceX, (float64)(e.Y)+distanceY, unitID)
	} else {
		fmt.Println("MOVE_ATTACK", (float64)(e.X)-distanceX, (float64)(e.Y)-distanceY, unitID)
	}
}
