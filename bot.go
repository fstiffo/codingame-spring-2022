package main

import "fmt"

// Hero's position indexes
const BACK = 0
const MIDDLE = 1
const FRONT = 2

// Hero commands
var HeroCommands [3]string

// Find best command for each heroe
func FindBestCommandsGreen0(bases []Base, monsters []Monster, heroes, opponents map[int]Common) {
	SortMonsters(monsters, bases[0].x, bases[0].y)
	hs := CopyMap(heroes)
	for i := 0; i < len(monsters); i++ {
		if len(hs) == 0 {
			break
		}
		x, y := MonsterFinalPosition(monsters[i])
		if PositionInsideBoard(x, y) {
			hero := NearestHero(monsters[i], hs)
			if monsters[i].threatFor == 1 && bases[0].mana > 9 {
				HeroCommands[hero.id] = "SPELL WIND " + fmt.Sprintf("%d %d", bases[1].x, bases[1].y)
			} else {
				HeroCommands[hero.id] = fmt.Sprintf("MOVE %d %d", x, y)
			}
			delete(hs, hero.id)
		}
	}
	for _, hero := range hs {
		HeroCommands[hero.id] = "MOVE " + fmt.Sprintf("%d %d", bases[1].x, bases[1].y)
	}

}

// Find best command for each hero new version
func (s *State) FindBestCommandsBronze() {
	s.FindBestCommandsForBack()
	s.FindBestCommandsForMiddle()
	s.FindBestCommandsForFront()
}

// Find best command the hero in back
func (s *State) FindBestCommandsForBack() {
	command := "WAIT"

	// If there is a monster near my base, first try to push it away from base otherwise attack it
	monster, ok := s.NearBase()
	if ok {
		if s.bases[0].mana > 9 &&
			s.HeroIsNearer(s.heroes[BACK], monster) &&
			DistanceIsBetween(s.heroes[BACK].x, s.heroes[BACK].y, monster.x, monster.y, DamageRange, WindRange) {
			command = fmt.Sprintf("SPELL WIND %d %d", s.bases[1].x, s.bases[1].y)
			goto done
		} else {
			command = fmt.Sprintf("MOVE %d %d", monster.x, monster.y)
			s.target[BACK] = monster.id
			goto done
		}
	}

	// If hero is too far from the base, move to it
	if dist(s.heroes[BACK].x, s.heroes[BACK].y, s.bases[0].x, s.bases[0].y) > BackRadius {
		command = fmt.Sprintf("MOVE %d %d", s.bases[0].x, s.bases[0].y)
		goto done
	}

	// If hero have a target, go to it first
	monster, ok = s.monsters[s.target[BACK]]
	if ok {
		command = fmt.Sprintf("MOVE %d %d", monster.x, monster.y)
		goto done
	}
	// Else target the nearest monster
	monster, ok = s.NearestMonster(s.heroes[BACK])
	if ok {
		if DistanceIsBetween(s.heroes[BACK].x, s.heroes[BACK].y, monster.x, monster.y, 0, BackRadius) {
			command = fmt.Sprintf("MOVE %d %d", monster.x, monster.y)
			s.target[BACK] = monster.id
		}
	}

done:
	Trace("Back target:", s.target[BACK])
	HeroCommands[BACK] = command + " B"
}

// Find best command for the hero in the middle
func (s *State) FindBestCommandsForMiddle() {
	// By default, move to the middle point
	command := fmt.Sprintf("MOVE %d %d", s.middleX, s.middleY)

	// If hero have a target, go to it first
	monster, ok := s.monsters[s.target[MIDDLE]]
	if ok {
		command = fmt.Sprintf("MOVE %d %d", monster.x, monster.y)
		goto done
	}

	// If hero is too far from the middle, move to it
	if dist(s.heroes[MIDDLE].x, s.heroes[MIDDLE].y, s.middleX, s.middleY) > MiddleRadius {
		command = fmt.Sprintf("MOVE %d %d", s.middleX, s.middleY)
		goto done
	}

	// If there is a monster targeting my base, attack it
	monster, ok = s.ThreatMonster(1)
	if ok {
		command = fmt.Sprintf("MOVE %d %d", monster.x, monster.y)
		s.target[MIDDLE] = monster.id
		goto done
	}
	// Else target the nearest monster
	monster, ok = s.NearestMonster(s.heroes[MIDDLE])
	if ok {
		command = fmt.Sprintf("MOVE %d %d", monster.x, monster.y)
		s.target[MIDDLE] = monster.id
	}

done:
	HeroCommands[MIDDLE] = command + " M"
}

// Find best command for the hero in front
func (s *State) FindBestCommandsForFront() {
	// By default, move to the center of the board
	command := fmt.Sprintf("MOVE %d %d", BoardCenterX, BoardCenterY)

	// If hero have a target, go to it first
	monster, ok := s.monsters[s.target[FRONT]]
	if ok {
		command = fmt.Sprintf("MOVE %d %d", monster.x, monster.y)
		goto done
	}

	// If hero is too far from the front, move to it
	if dist(s.heroes[FRONT].x, s.heroes[FRONT].y, BoardCenterX, BoardCenterY) > FrontRadius {
		command = fmt.Sprintf("MOVE %d %d", BoardCenterX, BoardCenterY)
		goto done
	}

	// If there is a monster targeting the opponent's base, push it otherwise leave it alone
	monster, ok = s.ThreatMonster(2)
	if ok && s.bases[0].mana > 9 &&
		s.HeroIsNearer(s.heroes[FRONT], monster) &&
		DistanceIsBetween(s.heroes[FRONT].x, s.heroes[FRONT].y, monster.x, monster.y, DamageRange, WindRange) {
		command = fmt.Sprintf("SPELL WIND %d %d", s.bases[1].x, s.bases[1].y)
		goto done
	}
	// Else target the nearest monster
	monster, ok = s.NearestMonster(s.heroes[FRONT])
	if ok {
		command = fmt.Sprintf("MOVE %d %d", monster.x, monster.y)
		s.target[FRONT] = monster.id
	}

done:
	HeroCommands[FRONT] = command + " F"
}
