package main

import "fmt"

// Hero commands
var HeroCommands = make(map[int]string)

// Find best command for each hero
func FindBestCommands(bases []Base, monsters []Entity, heroes, opponents map[int]Entity) {
	SortMonsters(monsters, bases[0].x, bases[0].y)
	hs := CopyMap(heroes)
	for i := 0; i < len(monsters); i++ {
		if len(hs) == 0 {
			break
		}
		x, y := MonsterFinalPosition(monsters[i])
		if PositionInsideBoard(x, y) {
			hero := NearestHero(monsters[i], hs)
			HeroCommands[hero.id] = fmt.Sprintf("MOVE %d %d", x, y)
			delete(hs, hero.id)

			// Add a second hero if the monster is strong and a threat
			if monsters[i].threatFor == 1 && monsters[i].health > 2 {
				hero := NearestHero(monsters[i], hs)
				HeroCommands[hero.id] = fmt.Sprintf("MOVE %d %d", x, y)
				delete(hs, hero.id)
			}
		}
	}
	for _, hero := range hs {
		HeroCommands[hero.id] = "WAIT"
	}

}
