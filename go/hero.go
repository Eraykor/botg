package main

import (
	"fmt"
	"math"
)

type Hero struct {
	*Entity

	Inventory *Inventory

	FarestAlly   *Entity
	NearestEnemy *Entity

	LastHitableEnemy *Entity
	LastHitableAlly  *Entity
}

func NewHero(entity *Entity) *Hero {
	return &Hero{
		Entity: entity,
	}
}

func (h *Hero) PhaseDefense(gs *GameState, rs *RoundState) bool {
	for _, ally := range rs.Allies {
		ally.DistanceToHero = h.DistanceTo(ally)
		ally.DistanceToTower = rs.Tower.DistanceTo(ally)

		if h.FarestAlly == nil || ally.DistanceToTower >= h.FarestAlly.DistanceToTower {
			h.FarestAlly = ally
		}

		if ally.DistanceToHero < (float64)(h.AttackRange+h.MovementSpeed) && ally.Health <= h.AttackDamage {
			h.LastHitableAlly = ally
		}
	}

	for _, enemy := range rs.Enemies {
		enemy.DistanceToHero = h.DistanceTo(enemy)
		enemy.DistanceToTower = rs.Tower.DistanceTo(enemy)

		if h.NearestEnemy == nil || enemy.DistanceToTower < h.NearestEnemy.DistanceToTower {
			h.NearestEnemy = enemy
		}
	}

	if h.DistanceTo(rs.Tower) < 100 {
		return false
	}

	safeDistance := (float64)(h.MovementSpeed) / 2
	if h.FarestAlly.DistanceToTower < rs.Tower.DistanceTo(h.Entity) {
		if h.FarestAlly.DistanceToHero > (float64)(h.MovementSpeed) {
			h.Move((float64)(-h.MovementSpeed), 0)
		} else if h.FarestAlly.DistanceToHero+h.NearestEnemy.DistanceToHero+safeDistance < (float64)(h.AttackRange) {
			h.MoveAndAttack(-h.FarestAlly.DistanceToHero-safeDistance, 0, h.NearestEnemy.UnitID)
		} else {
			h.Move(-h.FarestAlly.DistanceToHero-safeDistance, 0)
		}
		return true
	}

	return false
}

func (h *Hero) PhasePurchase(gs *GameState, rs *RoundState) bool {
	var bestPotion *Item
	for _, potion := range gs.Potions {
		if potion.Health > 0 && h.MaxHealth-h.Health >= potion.Health && rs.Gold >= potion.Cost {
			bestPotion = potion
		}
	}

	if bestPotion != nil {
		fmt.Println("BUY", bestPotion.Name)
		rs.Gold -= bestPotion.Cost
		return true
	}

	var bestItem *Item
	for _, item := range gs.Items {
		if item.Damage > 0 && rs.Gold > item.Cost {
			if bestItem == nil || bestItem.Damage < item.Damage {
				bestItem = item
			}
		}
	}

	if bestItem != nil {
		if h.ItemsOwned < 3 {
			h.Inventory.Buy(bestItem)
			rs.Gold -= bestItem.Cost
			return true
		} else {
			if h.Inventory.WeakestItem != nil && bestItem.Damage > h.Inventory.WeakestItem.Damage {
				h.Inventory.Sell(h.Inventory.WeakestItem)
				return true
			}
		}
	}

	return false
}

func (h *Hero) PhaseSkill(gs *GameState, rs *RoundState) bool {
	switch h.HeroType {
	case "IRONMAN":
		for _, enemy := range rs.Enemies {
			if enemy.DistanceTo(h.Entity) < 250 && h.Mana > 50 && h.Countdown3 == 0 {
				fmt.Println("BURNING", enemy.X, enemy.Y)
				return true
			}
		}
	case "DOCTOR_STRANGE":
		for _, enemy := range rs.Enemies {
			if enemy.UnitType == "HERO" {
				if enemy.DistanceTo(h.Entity) < 300 && h.Mana >= 40 && h.Countdown3 == 0 &&
					h.DistanceTo(rs.Tower) < (float64)(rs.Tower.AttackRange-200) {
					fmt.Println("PULL", enemy.UnitID)
					return true
				}
			}
		}
	}
	return false
}

func (h *Hero) PhaseAttack(gs *GameState, rs *RoundState) bool {
	for _, enemy := range rs.Enemies {
		if enemy.DistanceToHero < (float64)(h.AttackRange+h.MovementSpeed)-math.Floor(h.FarestAlly.DistanceToHero) {
			if enemy.UnitType == "HERO" && h.NearestEnemy.DistanceToTower < h.FarestAlly.DistanceToTower {
				if h.NearestEnemy == nil || enemy.Health < h.NearestEnemy.Health {
					h.NearestEnemy = enemy
				}
				break
			}
		} else if enemy.Health <= h.AttackDamage && (h.LastHitableEnemy == nil || enemy.Health < h.LastHitableEnemy.Health) {
			h.LastHitableEnemy = enemy
		}
	}

	target := h.NearestEnemy

	if h.LastHitableEnemy != nil {
		target = h.LastHitableEnemy
	} else if h.LastHitableAlly != nil {
		target = h.LastHitableAlly
	}

	var didAttack bool

	if (float64)(h.AttackRange) >= target.DistanceToHero {
		if (float64)(h.AttackRange)*0.9 < target.DistanceToHero {
			h.Attack(target.UnitID)
		} else {
			h.MoveAndAttack(-((float64)(h.AttackRange)*0.9 - target.DistanceToHero), 0, target.UnitID)
		}
		didAttack = true
	} else {
		if (float64)(h.AttackRange+h.MovementSpeed) > target.DistanceToHero {
			h.MoveAndAttack(target.DistanceToHero-(float64)(h.AttackRange)+10, 0, target.UnitID)
			didAttack = true
		} else {
			h.Move((float64)(h.MovementSpeed), 0)
		}
	}

	if didAttack {
		if target.Team == gs.MyTeam {
			for idx, ally := range rs.Allies {
				if ally == target {
					rs.Allies = append(rs.Allies[:idx], rs.Allies[idx+1:]...)
					break
				}
			}
		} else {
			for idx, enemy := range rs.Enemies {
				if enemy == target {
					rs.Enemies = append(rs.Enemies[:idx], rs.Enemies[idx+1:]...)
					break
				}
			}
		}
	}

	return true
}
