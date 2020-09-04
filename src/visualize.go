package rubik

import (
	"os/exec"
	"github.com/go-vgo/robotgo"
	"strings"
)

// formatNotation replaces 42 spin notation with iamthecu.be notation
func formatNotation(spin string) string {
	spin = strings.ReplaceAll(spin, "U'", "u")
	spin = strings.ReplaceAll(spin, "U2", "U U")
	spin = strings.ReplaceAll(spin, "D'", "d")
	spin = strings.ReplaceAll(spin, "D2", "D D")
	spin = strings.ReplaceAll(spin, "R'", "r")
	spin = strings.ReplaceAll(spin, "R2", "R R")
	spin = strings.ReplaceAll(spin, "L'", "l")
	spin = strings.ReplaceAll(spin, "L2", "L L")
	spin = strings.ReplaceAll(spin, "F'", "f")
	spin = strings.ReplaceAll(spin, "F2", "F F")
	spin = strings.ReplaceAll(spin, "B'", "b")
	spin = strings.ReplaceAll(spin, "B2", "B B")
	return spin
}

// spinCube spins 360 "up" "down" "right" or "left"
func spinCube(direction string) {
	for i := 0; i < 4; i++ {
		robotgo.KeyTap(direction)
	}
}

func runGraphic(mix string, solution string, visualizer bool) {
	if visualizer {
		exec.Command("open", "http://iamthecu.be/").Run()
		// // Cmd + Option + J. // -jsconsole
		robotgo.Sleep(1)
		robotgo.KeyTap("enter")
		robotgo.Sleep(5)
		robotgo.TypeStr(formatNotation(mix))
		spinCube("right")
		robotgo.TypeStr(formatNotation(solution))
		spinCube("right")
		spinCube("up")
	}
}