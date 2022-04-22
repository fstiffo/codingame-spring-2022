package main

import "fmt"

// Hero commands
var HeroCommands = make(map[int]string)

// Find best command for each hero
func FindBestCommands(bases []Base, monsters []Entity, heroes, opponents map[int]Entity) {
	SortMonsters(monsters, bases[0].x, bases[0].y)
	for i := 0; i < len(monsters); i++ {
		if len(heroes) == 0 {
			break
		}
		x, y := MonsterFinalPosition(monsters[i])
		if PositionInsideBoard(x, y) {
			hero := NearestHero(monsters[i], heroes)
			HeroCommands[hero.id] = fmt.Sprintf("MOVE %d %d", x, y)
			delete(heroes, hero.id)
		}
	}
	for _, hero := range heroes {
		HeroCommands[hero.id] = "WAIT"
	}

}
