package main

type Base struct {
	x      int // The corner of the map representing your base
	y      int // The corner of the map representing your base
	health int // Your base's remaining health
	mana   int // Your base's remaining mana
}

// Create a new base
func NewBase(x, y, health, mana int) Base {
	return Base{x, y, health, mana}
}

type Common struct {
	id           int  // Unique identifier
	_type        int  // 0=monster, 1=your hero, 2=opponent hero
	x            int  // Position of this entity
	y            int  // Position of this entity
	shieldLife   int  // Ignore for this league; Count down until shield spell fades
	isControlled bool // Ignore for this league; True when this entity is under a control spell
}

// Create a new Common
func NewCommon(id, _type, x, y, shieldLife, isControlled int) Common {
	return Common{id, _type, x, y, shieldLife, isControlled == 1}
}

type Monster struct {
	Common
	health    int  // Remaining health of this monster
	vx        int  // Trajectory of this monster
	vy        int  // Trajectory of this monster
	nearBase  bool // true=monster with no target yet, false=monster targeting a base
	threatFor int  // Given this monster's trajectory, is it a threat to 1=your base, 2=your opponent's base, 0=neither
}

// Create a new Monster
func NewMonster(id, _type, x, y, shieldLife, isControlled, health, vx, vy, nearBase, threatFor int) Monster {
	return Monster{Common{id, _type, x, y, shieldLife, isControlled == 1}, health, vx, vy, nearBase == 1, threatFor}
}

type State struct {
	bases     [2]Base         // Your base and opponent's base
	monsters  map[int]Monster // All monsters on the board
	heroes    [3]Common       // Your heroes
	opponents [3]Common       // Opponent's heroes
	target    [3]int          // Your hero's target
	middleX   int             // The middle point beetween your base and the center of the board
	middleY   int
}

// Create a new State
func NewState(bases [2]Base, monsters map[int]Monster, heroes [3]Common, opponents [3]Common) State {
	var heroPtrs [3]*Common
	i := 0
	for _, v := range heroes {
		heroPtrs[i] = &v
		i++
	}
	middleX := (bases[0].x + BoardCenterX) / 2
	middleY := (bases[0].y + BoardCenterY) / 2
	return State{bases, monsters, heroes, opponents, [3]int{-1, -1, -1}, middleX, middleY}
}

// Update the state
func (s *State) Update(bases [2]Base, monsters map[int]Monster, heroes [3]Common, opponents [3]Common) {
	s.bases = bases
	s.monsters = monsters
	s.heroes = heroes
	s.opponents = opponents
}
