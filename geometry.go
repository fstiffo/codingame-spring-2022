package main

import (
	"math"
)

const (
	BoardLength                              = 17630
	BoardWidth                               = 9000
	BoardCenterX                             = BoardLength / 2
	BoardCenterY                             = BoardWidth / 2
	BackRadius                               = 5000
	DamageRange                              = 800
	WindRange                                = 1280
	SpellRange                               = 2200
	TglCenterStartX, TglCenterStartY         = 953, 953 // l*sqrt(3)/4 where l = SpellRange
	_30DegInRadians                  float64 = math.Pi / 6
)

var FrontRadius = dist(0, 0, BoardLength, BoardWidth) + 2500
var MiddleRadius = (FrontRadius + BackRadius) / 2
var DiagX, DiagY = norm(BoardLength, BoardWidth)

// Normalize a 2D vector
func norm(x, y int) (float64, float64) {
	length := dist(0, 0, x, y)
	return float64(x) / length, float64(y) / length

}

// Integer distance between two points
func dist(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64(x1-x2)*float64(x1-x2) + float64(y1-y2)*float64(y1-y2))
}

// Check if the distance between two points is between the two values d1 and d2
func DistanceIsBetween(x1, y1, x2, y2 int, d1, d2 int) bool {
	d := dist(x1, y1, x2, y2)
	return d >= float64(d1) && d <= float64(d2)
}

// Rotate a 2D vector by an angle with y axis
func Rot(x, y float64, angle float64) (float64, float64) {
	return x*math.Cos(angle) - y*math.Sin(angle), x*math.Sin(angle) + y*math.Cos(angle)
}

// Multiply a 2D vector by a scalar
func Mul(x, y float64, scalar float64) (float64, float64) {
	return x * scalar, y * scalar
}

// Add 2D vectors
func Add(x1, y1, x2, y2 float64) (float64, float64) {
	return x1 + x2, y1 + y2
}

// Position inside the board
func PositionInsideBoard(x, y int) bool {
	return x >= 0 && x < BoardLength && y >= 0 && y < BoardWidth
}

// Calculate monster's final position
func MonsterFinalPosition(monster Monster) (int, int) {
	return monster.x + monster.vx, monster.y + monster.vy
}

// Sort monsters by distance from base
func SortMonsters(monsters []Monster, baseX, baseY int) {
	for i := 0; i < len(monsters); i++ {
		for j := i + 1; j < len(monsters); j++ {
			if dist(monsters[i].x, monsters[i].y, baseX, baseY) > dist(monsters[j].x, monsters[j].y, baseX, baseY) {
				monsters[i], monsters[j] = monsters[j], monsters[i]
			}
		}
	}
}

// Return a slice of monsters ids sorted ascending by distance from a point
func MonstersSortedByDistance(monsters map[int]Monster, x, y int) []int {
	// Convert map to slice of ids
	ids := []int{}
	for id := range monsters {
		ids = append(ids, id)
	}
	// Sort ids ascending by distance of the monster from point (x, y)
	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			if dist(monsters[ids[i]].x, monsters[ids[i]].y, x, y) > dist(monsters[ids[j]].x, monsters[ids[j]].y, x, y) {
				ids[i], ids[j] = ids[j], ids[i]
			}
		}
	}
	// Debug
	for i, id := range ids {
		Trace("Sorted:", i, id, monsters[id].x, monsters[id].y)
	}
	return ids
}

// Sort monsters by distance to Base and threat level
func SortMonstersByThreat(monsters []Monster, baseX, baseY int) {
	for i := 0; i < len(monsters); i++ {
		for j := i + 1; j < len(monsters); j++ {
			if dist(monsters[i].x, monsters[i].y, baseX, baseY) > dist(monsters[j].x, monsters[j].y, baseX, baseY) {
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
		if dist(monster.x, monster.y, hero.x, hero.y) < dist(monster.x, monster.y, nearestHero.x, nearestHero.y) {
			nearestHero = hero
		}
	}
	return nearestHero
}

// Find neareast monster to a hero
func (s *State) NearestMonster(hero Common) (Monster, bool) {
	nearestMonster := Monster{}
	nearestMonster.id = -1
	nearestMonster.x, nearestMonster.y = -BoardLength, -BoardWidth
	monsters := s.monsters
	for _, monster := range monsters {
		if dist(hero.x, hero.y, monster.x, monster.y) < dist(hero.x, hero.y, nearestMonster.x, nearestMonster.y) {
			nearestMonster = monster
		}
	}
	return nearestMonster, (nearestMonster.id != -1)
}

// Find first monster near my base
func (s *State) NearBase() (Monster, bool) {
	for id := range s.sorted {
		if s.monsters[id].nearBase && s.monsters[id].threatFor == 1 {
			return s.monsters[id], true
		}
	}
	return Monster{}, false
}

// Find first monster that:
// 0 = will never reach the base
// 1 = will eventually reach my the base
// 2 = will eventually reach opponent's the base
func (s *State) ThreatMonster(base int) (Monster, bool) {
	for _, monster := range s.monsters {
		if monster.threatFor == base {
			return monster, true
		}
	}
	return Monster{}, false
}

// Check if the hero is nearer to his base than the monster
func (s *State) HeroIsNearer(hero Common, monster Monster) bool {
	return dist(hero.x, hero.y, s.bases[0].x, s.bases[0].y) < dist(monster.x, monster.y, s.bases[0].x, s.bases[0].y)
}
