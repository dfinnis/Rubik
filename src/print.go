package rubik

import (
	"fmt"
	"os"
	"time"
)

const Reset		= "\x1B[0m"
const Bright	= "\x1B[1m"
const Green		= "\x1B[32m"
const Yellow	= "\x1B[33m"

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

// dumpCube prints cube state, & highlights features solved in given subgroup (-1 = no group)
func dumpCube(cube *cepo, group int8) {
	fmt.Printf("\n")
	fmt.Printf("+-------------+-----------------+---------------------------------------+\n")
	fmt.Printf("|             | Corner          |  Edge                                 |\n")
	fmt.Printf("|      Number | 0 1 2 3 4 5 6 7 |  0  1  2  3   4  5  6  7   8  9 10 11 |\n")
	fmt.Printf("+-------------+-----------------+-------------+------------+------------+\n")
	fmt.Printf("| Permutation | ")
	// Corner Permutation
	if group == 2 {
		fmt.Printf("%v", Bright)
	}
	if group >= 2 {
		fmt.Printf("%v", Green)
	}
	for _, cP := range cube.cP {
		fmt.Printf("%v ", cP)
	}
	fmt.Printf("%v| ", Reset)

	// Edge Permutation
	for i, eP := range cube.eP {
		if i == 4 || i == 8 {
			fmt.Printf(" ")
		}
		if group == 1 && i == 8 {
			fmt.Printf("%v", Bright)
		}
		if group == 1 && i == 8 {
			fmt.Printf("%v", Yellow)
		}
		if group == 2 {
			if i == 0 {
				fmt.Printf("%v", Bright)
				fmt.Printf("%v", Yellow)
			}
			if i == 8 {
				fmt.Printf("%v", Reset)
				fmt.Printf("%v", Yellow)
			}
		}
		if group == 3 && i == 0 {
			fmt.Printf("%v", Bright)
			fmt.Printf("%v", Green)
		}
		fmt.Printf("%2v ", eP)
	}
	fmt.Printf("%v|", Reset)

	fmt.Printf("\n| Orientation | ")
	// Corner Orientation
	if group == 1 {
		fmt.Printf("%v", Bright)
	}
	if group >= 1 {
		fmt.Printf("%v", Green)
	}
	for _, cO := range cube.cO {
		fmt.Printf("%v ", cO)
	}
	fmt.Printf("%v| ", Reset)

	// Edge Orientation
	if group == 0 {
		fmt.Printf("%v", Bright)
	}
	if group >= 0 {
		fmt.Printf("%v", Green)
	}
	for i, eO := range cube.eO {
		if i == 4 || i == 8 {
			fmt.Printf(" ")
		}
		fmt.Printf("%2v ", eO)
	}
	fmt.Printf("%v|", Reset)
	fmt.Printf("\n+-------------+-----------------+---------------------------------------+\n")
}
