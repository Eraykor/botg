package main

import "fmt"

//import "os"

/**
 * Made with love by AntiSquid, Illedan and Wildum.
 * You can help children learn to code while you participate by donating to CoderDojo.
 **/

func main() {
	var myTeam int
	fmt.Scan(&myTeam)

	// bushAndSpawnPointCount: usefrul from wood1, represents the number of bushes and the number of places where neutral units can spawn
	var bushAndSpawnPointCount int
	fmt.Scan(&bushAndSpawnPointCount)

	for i := 0; i < bushAndSpawnPointCount; i++ {
		// entityType: BUSH, from wood1 it can also be SPAWN
		var entityType string
		var x, y, radius int
		fmt.Scan(&entityType, &x, &y, &radius)
	}
	// itemCount: useful from wood2
	var itemCount int
	fmt.Scan(&itemCount)

	for i := 0; i < itemCount; i++ {
		// itemName: contains keywords such as BRONZE, SILVER and BLADE, BOOTS connected by "_" to help you sort easier
		// itemCost: BRONZE items have lowest cost, the most expensive items are LEGENDARY
		// damage: keyword BLADE is present if the most important item stat is damage
		// moveSpeed: keyword BOOTS is present if the most important item stat is moveSpeed
		// isPotion: 0 if it's not instantly consumed
		var itemName string
		var itemCost, damage, health, maxHealth, mana, maxMana, moveSpeed, manaRegeneration, isPotion int
		fmt.Scan(&itemName, &itemCost, &damage, &health, &maxHealth, &mana, &maxMana, &moveSpeed, &manaRegeneration, &isPotion)
	}
	for {
		var gold int
		fmt.Scan(&gold)

		var enemyGold int
		fmt.Scan(&enemyGold)

		// roundType: a positive value will show the number of heroes that await a command
		var roundType int
		fmt.Scan(&roundType)

		var entityCount int
		fmt.Scan(&entityCount)

		for i := 0; i < entityCount; i++ {
			// unitType: UNIT, HERO, TOWER, can also be GROOT from wood1
			// shield: useful in bronze
			// stunDuration: useful in bronze
			// countDown1: all countDown and mana variables are useful starting in bronze
			// heroType: DEADPOOL, VALKYRIE, DOCTOR_STRANGE, HULK, IRONMAN
			// isVisible: 0 if it isn't
			// itemsOwned: useful from wood1
			var unitId, team int
			var unitType string
			var x, y, attackRange, health, maxHealth, shield, attackDamage, movementSpeed, stunDuration, goldValue, countDown1, countDown2, countDown3, mana, maxMana, manaRegeneration int
			var heroType string
			var isVisible, itemsOwned int
			fmt.Scan(&unitId, &team, &unitType, &x, &y, &attackRange, &health, &maxHealth, &shield, &attackDamage, &movementSpeed, &stunDuration, &goldValue, &countDown1, &countDown2, &countDown3, &mana, &maxMana, &manaRegeneration, &heroType, &isVisible, &itemsOwned)
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		// If roundType has a negative value then you need to output a Hero name, such as "DEADPOOL" or "VALKYRIE".
		// Else you need to output roundType number of any valid action, such as "WAIT" or "ATTACK unitId"
		fmt.Println("WAIT")
		fmt.Println("WAIT")
	}
}
