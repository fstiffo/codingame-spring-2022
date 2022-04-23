package main

import "math"

const (
	BoardLength  = 17630
	BoardWidth   = 9000
	BoardCenterX = BoardLength / 2
	BoardCenterY = BoardWidth / 2
	BackRadius   = 5000
)

var FrontRadius = distance(0, 0, BoardLength, BoardWidth) + 2500
var MiddleRadius = (FrontRadius + BackRadius) / 2

// Integer distance between two points
func distance(x1, y1, x2, y2 int) int {
	return int(math.Sqrt(float64(x1-x2)*float64(x1-x2) + float64(y1-y2)*float64(y1-y2)))
}

// Distance between two points in between two values
func DistanceIsBetween(x1, y1, x2, y2 int, d1, d2 int) bool {
	d := distance(x1, y1, x2, y2)
	return d >= d1 && d <= d2
}

// Position inside the board
func PositionInsideBoard(x, y int) bool {
	return x >= 0 && x < BoardLength && y >= 0 && y < BoardWidth
}

// Calculate monster's final position
func MonsterFinalPosition(monster Monster) (int, int) {
	return monster.x + monster.vx, monster.y + monster.vy
}

// Order monsters by distance to Base
func SortMonsters(monsters []Monster, baseX, baseY int) {
	for i := 0; i < len(monsters); i++ {
		for j := i + 1; j < len(monsters); j++ {
			if distance(monsters[i].x, monsters[i].y, baseX, baseY) > distance(monsters[j].x, monsters[j].y, baseX, baseY) {
				monsters[i], monsters[j] = monsters[j], monsters[i]
			}
		}
	}
}

// Sort monsters by distance to Base and threat level
func SortMonstersByThreat(monsters []Monster, baseX, baseY int) {
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
func NearestHero(monster Monster, heroes map[int]Common) Common {
	nearestHero := heroes[0]
	for _, hero := range heroes {
		if distance(monster.x, monster.y, hero.x, hero.y) < distance(monster.x, monster.y, nearestHero.x, nearestHero.y) {
			nearestHero = hero
		}
	}
	return nearestHero
}

// Find neareast monster to a hero
func NearestMonster(hero Common, s State) (Monster, bool) {
	nearestMonster := Monster{}
	nearestMonster.id = -1
	nearestMonster.x, nearestMonster.y = -BoardLength, -BoardWidth
	monsters := s.monsters
	for _, monster := range monsters {
		if distance(hero.x, hero.y, monster.x, monster.y) < distance(hero.x, hero.y, nearestMonster.x, nearestMonster.y) {
			nearestMonster = monster
		}
	}
	return nearestMonster, (nearestMonster.id != -1)
}

// Find first monster near my base
func NearBase(s State) (Monster, bool) {
	for _, monster := range s.monsters {
		if monster.nearBase {
			return monster, true
		}
	}
	return Monster{}, false
}

// Find first monster that is a threat for a base
func ThreatMonster(base int, s State) (Monster, bool) {
	for _, monster := range s.monsters {
		if monster.threatFor == base {
			return monster, true
		}
	}
	return Monster{}, false
}
