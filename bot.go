package main

import (
	"fmt"
	"os"
)

// Hero commands
var HeroCommands [3]string

// Find best commands for each hero
func (state *State) FindBestCommands() {
	if state.turn == 0 {
		if state.bottomRight { // My base at bottom-right corner
			DiagX, DiagY = -1*DiagX, -1*DiagY
		}
		fmt.Fprintln(os.Stderr, "BottomRight:", state.bottomRight)
	}
	if state.turn < 3 {
		HeroCommands[0] = "WAIT 0"
		x0, y0 := float64(state.heroes[0].x), float64(state.heroes[0].y)
		x, y := Add(x0, y0, float64(SpellRange), 0.0)
		if state.bottomRight {
			HeroCommands[1] = fmt.Sprintf("MOVE %d %d 1", BoardLength-1-int(x), BoardLength-1-int(y))
		} else {
			HeroCommands[1] = fmt.Sprintf("MOVE %d %d 1", int(x), int(y))
		}
		x, y = Add(x0, y0, 0.0, float64(SpellRange))
		if state.bottomRight {
			HeroCommands[2] = fmt.Sprintf("MOVE %d %d 1", BoardLength-1-int(x), BoardLength-1-int(y))
		} else {
			HeroCommands[2] = fmt.Sprintf("MOVE %d %d 1", int(x), int(y))
		}
	} else {
		HeroCommands[0] = "WAIT 0"
		HeroCommands[1] = "WAIT 1"
		HeroCommands[2] = "WAIT 2"
	}
}
