package rubik

import (
	"os/exec"
	"fmt"//
	"github.com/go-vgo/robotgo"
)
func formatNotation(spin string) string {
	return spin
}

func runGraphic(mix string, solution string) {
	exec.Command("open", "http://iamthecu.be/").Run()
	// // Cmd + Option + J. // -jsconsole
	robotgo.Sleep(1)
	robotgo.KeyTap("enter")
	robotgo.Sleep(5)

	// orientate cube to default. white up, green front
	robotgo.KeyTap("up")	// comment for demo
	robotgo.KeyTap("left")	// comment for demo
	
	// robotgo.TypeStr("rdRDrdRDrdRDrdRDrdRDrdRD")
	// robotgo.TypeStr("rdRD")

	mixFormated := formatNotation(mix)
	fmt.Printf("\nmix = %v\n", mixFormated)//
	robotgo.TypeStr(mixFormated)

	robotgo.KeyTap("right")
	robotgo.KeyTap("right")
	robotgo.KeyTap("right")
	robotgo.KeyTap("right")

	robotgo.Sleep(2)
	solutionFormated := formatNotation(solution)
	fmt.Printf("\nsolution = %v\n", solutionFormated)//
	robotgo.TypeStr(solutionFormated)

	robotgo.KeyTap("right")
	robotgo.KeyTap("right")
	robotgo.KeyTap("right")	
	robotgo.KeyTap("right")
	robotgo.KeyTap("up")
	robotgo.KeyTap("up")
	robotgo.KeyTap("up")
	robotgo.KeyTap("up")
}