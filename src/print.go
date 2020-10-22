package rubik

import (
	"fmt"
	"os"
	"time"
)

const Reset		= "\x1B[0m"
const Bright	= "\x1B[1m"
const Green		= "\x1B[32m"

// printSolution prints final output
func printSolution(solution string, elapsed time.Duration, cube *cepo) {
	if isSolved(cube) == false {
		fmt.Printf("\nError: Solution Incorrect :(\n")
	} else {
		fmt.Printf("%v\nSolution Correct, cube solved! :)\n%v", Green, Reset)//
	}
	fmt.Printf("\nHalf Turn Metric: %v\n", halfTurnMetric(solution))
	fmt.Printf("\n%vSolution:\n%v%v\n\n", Bright, Reset, solution)
	fmt.Printf("Solve time:\n%v\n\n", elapsed)
}

func errorExit(message string) {
	fmt.Printf("Error: %s\n", message)
	printUsage()
}

func printUsage() {
	fmt.Printf("\nUsage:\tgo build; ./Rubik \"mix\" [-r [length]] [-v] [-g] [-h]\n\n")
	fmt.Printf("    mix should be valid sequence string e.g.\n")
	fmt.Printf("    \"U U' U2 D D' D2 R R' R2 L L' L2 F F' F2 B B' B2\"\n")
	fmt.Printf("    or mix \"filepath\" e.g. \"mix/superflip.txt\" reads a file\n")
	fmt.Printf("    or mix \"-r [len]\" or \"--random [len]\" mixes randomly\n\n")
	fmt.Printf("    [-v] (--visualizer) show visual of mix and solution\n")
	fmt.Printf("    [-g] (--group) show solution breakdown by subgroup\n")
	fmt.Printf("    [-h] (--help) show usage\n\n")
	os.Exit(1)
}

// dumpCube prints cube state
func dumpCube(cube *cepo) {
	fmt.Printf("\n+-------------+-----------------+-------------------------------------+")
	fmt.Printf("\n|             | Corner          |  Edge                               |\n")
	fmt.Printf("|      Number | ")
	for i := range cube.cP {
		fmt.Printf("%v ", i)
	}
	fmt.Printf("| ")
	for i := range cube.eP {
		fmt.Printf("%2v ", i)
	}
	fmt.Printf("|")
	fmt.Printf("\n+-------------+-----------------+-------------------------------------+\n")
	fmt.Printf("| Permutation | ")
	for _, cP := range cube.cP {
		fmt.Printf("%v ", cP)
	}
	fmt.Printf("| ")
	for _, eP := range cube.eP {
		fmt.Printf("%2v ", eP)
	}
	fmt.Printf("|")

	fmt.Printf("\n| Orientation | ")
	for _, cO := range cube.cO {
		fmt.Printf("%v ", cO)
	}
	fmt.Printf("| ")
	for _, eO := range cube.eO {
		fmt.Printf("%2v ", eO)
	}
	fmt.Printf("|")
	fmt.Printf("\n+-------------+-----------------+-------------------------------------+\n")
}
