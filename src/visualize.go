package rubik

import (
	"os/exec"
	"fmt"//
	"github.com/go-vgo/robotgo"
	"strings"
)
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

func runGraphic(mix string, solution string) {
	exec.Command("open", "http://iamthecu.be/").Run()
	// // Cmd + Option + J. // -jsconsole
	robotgo.Sleep(1)
	robotgo.KeyTap("enter")
	robotgo.Sleep(5)

	// // orientate cube to default. white up, green front
	// robotgo.KeyTap("up")	// comment for demo
	// robotgo.KeyTap("left")	// comment for demo

	mixFormated := formatNotation(mix)
	fmt.Printf("\nmixFormated = %v\n", mixFormated)//
	robotgo.TypeStr(mixFormated)

	robotgo.KeyTap("right")
	robotgo.KeyTap("right")
	robotgo.KeyTap("right")
	robotgo.KeyTap("right")

	solution = "U U' U2 U2 D D' D2 D2 R R' R2 R2 L L' L2 L2 F F' F2 F2 B B' B2 B2" // test
	solutionFormated := formatNotation(solution)
	fmt.Printf("\nsolution = %v\n", solutionFormated)//
	robotgo.TypeStr(solutionFormated)

	// robotgo.KeyTap("right")
	// robotgo.KeyTap("right")
	// robotgo.KeyTap("right")	
	// robotgo.KeyTap("right")
	// robotgo.KeyTap("up")
	// robotgo.KeyTap("up")
	// robotgo.KeyTap("up")
	// robotgo.KeyTap("up")
}