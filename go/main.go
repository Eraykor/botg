package main

import (
	"fmt"
	"os"
)

type Phase struct {
	Name string // DEFENSE, PURCHASE, SKILL, ATTACK
	Do   func(*GameState, *RoundState) bool
}

func main() {
	gs := NewGameState(os.Stdin)

	heroPicks := []string{"IRONMAN", "HULK"}
	pick := 0

	heroesInventories := make(map[string]*Inventory)
	for _, heroPick := range heroPicks {
		heroesInventories[heroPick] = NewInventory()
	}

	for {
		rs := NewRoundState(os.Stdin, gs)

		if rs.Type < 0 {
			fmt.Println(heroPicks[pick])
			pick++
			continue
		}

		for _, hero := range rs.Heroes {
			hero.Inventory = heroesInventories[hero.HeroType]

			phases := []*Phase{
				&Phase{"DEFENSE", hero.PhaseDefense},
				&Phase{"PURCHASE", hero.PhasePurchase},
				&Phase{"SKILL", hero.PhaseSkill},
				&Phase{"ATTACK", hero.PhaseAttack},
			}

			for _, phase := range phases {
				if phase.Do(gs, rs) {
					fmt.Fprintln(os.Stderr, "HERO", hero.HeroType, phase.Name)
					continue
				}
			}
		}
	}
}
