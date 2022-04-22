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

type Entity struct {
	id           int  // Unique identifier
	_type        int  // 0=monster, 1=your hero, 2=opponent hero
	x            int  // Position of this entity
	y            int  // Position of this entity
	shieldLife   int  // Ignore for this league; Count down until shield spell fades
	isControlled bool // Ignore for this league; True when this entity is under a control spell
	health       int  // Remaining health of this monster
	vx           int  // Trajectory of this monster
	vy           int  // Trajectory of this monster
	nearBase     bool // true=monster with no target yet, false=monster targeting a base
	threatFor    int  // Given this monster's trajectory, is it a threat to 1=your base, 2=your opponent's base, 0=neither
}

// Create a new entity
func NewEntity(id, _type, x, y, shieldLife, isControlled, health, vx, vy, nearBase, threatFor int) Entity {
	return Entity{id, _type, x, y, shieldLife, isControlled == 1, health, vx, vy, nearBase == 1, threatFor}
}
