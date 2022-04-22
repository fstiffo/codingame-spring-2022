package main

import (
	"fmt"
	"os"
)

func main() {
	// baseX: The corner of the map representing your base
	var baseX, baseY int
	fmt.Scan(&baseX, &baseY)

	// heroesPerPlayer: Always 3
	var heroesPerPlayer int
	fmt.Scan(&heroesPerPlayer)

	for {
		bases := make([]Base, 2) // Bases of each player, in the order your base, then opponent's base
		for i := 0; i < 2; i++ {
			var health, mana int
			fmt.Scan(&health, &mana)
			bases[i] = Base{0, 0, health, mana}
		}
		bases[0].x = baseX
		bases[0].y = baseY
		bases[1].x = BoardLength - baseX
		bases[1].y = BoardWidth - baseY

		// entityCount: Amount of heros and monsters you can see
		var entityCount int
		fmt.Scan(&entityCount)

		monsters := make([]Entity, entityCount-2*heroesPerPlayer)
		heroes := make(map[int]Entity)
		opponents := make(map[int]Entity)

		for monstersCount, i := 0, 0; i < entityCount; i++ {
			var id, _type, x, y, shieldLife, isControlled, health, vx, vy, nearBase, threatFor int
			fmt.Scan(&id, &_type, &x, &y, &shieldLife, &isControlled, &health, &vx, &vy, &nearBase, &threatFor)
			fmt.Fprintln(os.Stderr, _type, id, x, y, vx, vy)
			switch _type {
			case 0:
				monsters[monstersCount] = NewEntity(id, _type, x, y, shieldLife, isControlled, health, vx, vy, nearBase, threatFor)
				monstersCount += 1
			case 1:
				heroes[id] = NewEntity(id, _type, x, y, shieldLife, isControlled, health, vx, vy, nearBase, threatFor)
			case 2:
				opponents[id] = NewEntity(id, _type, x, y, shieldLife, isControlled, health, vx, vy, nearBase, threatFor)
			}
		}

		FindBestCommands(bases, monsters, heroes, opponents)

		for i := 0; i < heroesPerPlayer; i++ {

			// In the first league: MOVE <x> <y> | WAIT; In later leagues: | SPELL <spellParams>;
			fmt.Println(HeroCommands[i])
		}
	}
}
