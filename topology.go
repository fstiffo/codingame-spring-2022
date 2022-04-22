package main

import "math"

const BoardLength = 17630
const BoardWidth = 9000

// Integer distance between two points
func distance(x1, y1, x2, y2 int) int {
	return int(math.Sqrt(float64(x1-x2)*float64(x1-x2) + float64(y1-y2)*float64(y1-y2)))
}

// Position inside the board
func PositionInsideBoard(x, y int) bool {
	return x >= 0 && x < BoardLength && y >= 0 && y < BoardWidth
}

// Calculate monster's final position
func MonsterFinalPosition(monster Entity) (int, int) {
	return monster.x + monster.vx, monster.y + monster.vy
}

// Order monsters by distance to Base
func SortMonsters(monsters []Entity, baseX, baseY int) {
	for i := 0; i < len(monsters); i++ {
		for j := i + 1; j < len(monsters); j++ {
			if distance(monsters[i].x, monsters[i].y, baseX, baseY) > distance(monsters[j].x, monsters[j].y, baseX, baseY) {
				monsters[i], monsters[j] = monsters[j], monsters[i]
			}
		}
	}
}

// Sort monsters by distance to Base and threat level
func SortMonstersByThreat(monsters []Entity, baseX, baseY int) {
	for i := 0; i < len(monsters); i++ {
		for j := i + 1; j < len(monsters); j++ {
			if distance(monsters[i].x, monsters[i].y, baseX, baseY) > distance(monsters[j].x, monsters[j].y, baseX, baseY) {
				monsters[i], monsters[j] = monsters[j], monsters[i]
			} else if monsters[i].threatFor > monsters[j].threatFor {
				monsters[i], monsters[j] = monsters[j], monsters[i]
			}
		}
	}
}

// Find neareast hero to a monster
func NearestHero(monster Entity, heroes map[int]Entity) Entity {
	nearestHero := heroes[0]
	for _, hero := range heroes {
		if distance(monster.x, monster.y, hero.x, hero.y) < distance(monster.x, monster.y, nearestHero.x, nearestHero.y) {
			nearestHero = hero
		}
	}
	return nearestHero
}
