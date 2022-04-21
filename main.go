package main

import (
	"fmt"
	"os"
)

type Base struct {
	health int
	mana   int
}

type Entity struct {
	id           int // id: Unique identifier
	_type        int // type: 0=monster, 1=your hero, 2=opponent hero
	x            int // x: Position of this entity
	y            int // u: Position of this entity
	shieldLife   int // shieldLife: Ignore for this league; Count down until shield spell fades
	isControlled int // isControlled: Ignore for this league; Equals 1 when this entity is under a control spell
	health       int // health: Remaining health of this monster
	vx           int // vx: Trajectory of this monster
	vy           int // vy: Trajectory of this monster
	nearBase     int // nearBase: 0=monster with no target yet, 1=monster targeting a base
	threatFor    int // threatFor: Given this monster's trajectory, is it a threat to 1=your base, 2=your opponent's base, 0=neither
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func manhattanDist(x1, y1, x2, y2 int) int {
	return absDiffInt(x1, x2) + absDiffInt(y1, y2)
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// baseX: The corner of the map representing your base
	var baseX, baseY int
	fmt.Scan(&baseX, &baseY)

	// heroesPerPlayer: Always 3
	var heroesPerPlayer int
	fmt.Scan(&heroesPerPlayer)

	for {
		for i := 0; i < 2; i++ {
			// health: Your base health
			// mana: Ignore in the first league; Spend ten mana to cast a spell
			var health, mana int
			fmt.Scan(&health, &mana)
		}
		// entityCount: Amount of heros and monsters you can see
		var entityCount int
		fmt.Scan(&entityCount)

		monsterEntities := make([]Entity, entityCount-6)
		myHeroes := make([]Entity, 3)
		opponentHeroes := make([]Entity, 3)
		var monstersCount, myHeroesCount, opponentHeroesCount int

		for i := 0; i < entityCount; i++ {
			// id: Unique identifier
			// type: 0=monster, 1=your hero, 2=opponent hero
			// x: Position of this entity
			// shieldLife: Ignore for this league; Count down until shield spell fades
			// isControlled: Ignore for this league; Equals 1 when this entity is under a control spell
			// health: Remaining health of this monster
			// vx: Trajectory of this monster
			// nearBase: 0=monster with no target yet, 1=monster targeting a base
			// threatFor: Given this monster's trajectory, is it a threat to 1=your base, 2=your opponent's base, 0=neither
			var id, _type, x, y, shieldLife, isControlled, health, vx, vy, nearBase, threatFor int
			fmt.Scan(&id, &_type, &x, &y, &shieldLife, &isControlled, &health, &vx, &vy, &nearBase, &threatFor)
			fmt.Fprintln(os.Stderr, _type, id, x, y, vx, vy, manhattanDist(0, 0, vx, vy))
			switch _type {
			case 0:
				monsterEntities[monstersCount] =
					Entity{id, _type, x, y, shieldLife, isControlled, health, vx, vy, nearBase, threatFor}
				monstersCount += 1
			case 1:
				myHeroes[myHeroesCount] =
					Entity{id, _type, x, y, shieldLife, isControlled, health, vx, vy, nearBase, threatFor}
				myHeroesCount += 1
			case 2:
				opponentHeroes[opponentHeroesCount] =
					Entity{id, _type, x, y, shieldLife, isControlled, health, vx, vy, nearBase, threatFor}
				opponentHeroesCount += 1
			}

		}
		for i := 0; i < heroesPerPlayer; i++ {

			// In the first league: MOVE <x> <y> | WAIT; In later leagues: | SPELL <spellParams>;
			fmt.Println("WAIT")
		}
	}
}
