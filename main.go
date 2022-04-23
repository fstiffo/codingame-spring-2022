package main

import (
	"fmt"
	"os"
)

func main() {
	firstRound := true
	var state State

	// baseX: The corner of the map representing your base
	var baseX, baseY int
	fmt.Scan(&baseX, &baseY)

	// heroesPerPlayer: Always 3
	var heroesPerPlayer int
	fmt.Scan(&heroesPerPlayer)

	for {
		var bases [2]Base // Bases of each player, in the order your base, then opponent's base
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

		monsters := make(map[int]Monster)
		var heroes [3]Common
		var opponents [3]Common

		for heroCount, opponentCount, i := 0, 0, 0; i < entityCount; i++ {
			var id, _type, x, y, shieldLife, isControlled, health, vx, vy, nearBase, threatFor int
			fmt.Scan(&id, &_type, &x, &y, &shieldLife, &isControlled, &health, &vx, &vy, &nearBase, &threatFor)
			fmt.Fprintln(os.Stderr, _type, id, x, y, vx, vy)
			switch _type {
			case 0:
				monsters[id] = NewMonster(id, _type, x, y, shieldLife, isControlled, health, vx, vy, nearBase, threatFor)
			case 1:
				heroes[heroCount] = NewCommon(id, _type, x, y, shieldLife, isControlled)
				heroCount++
			case 2:
				opponents[opponentCount] = NewCommon(id, _type, x, y, shieldLife, isControlled)
				opponentCount++
			}
		}
		if firstRound {
			state = NewState(bases, monsters, heroes, opponents)
		} else {
			state.Update(bases, monsters, heroes, opponents)
		}

		FindBestCommands2(state)

		for i := range heroes {

			// In the first league: MOVE <x> <y> | WAIT; In later leagues: | SPELL <spellParams>;
			fmt.Println(HeroCommands[i])
		}
	}
}
