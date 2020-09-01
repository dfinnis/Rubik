package rubik

import (
	"fmt"
	"time"
	"os"
	"io/ioutil"
	"strconv"
	"strings"
)

// corner permutation cP[0] is face U top left cubie
// edge permutation   eP[0] is face U left cubie

type cepo struct {
	cP 		[8]int8		// cornerPermutation	(0-7)
	cO 		[8]int8		// cornerOrientation	(0-2)	0 = good, 1 = twisted clockwise, 2 = twisted anti-clockwise
	eP 		[12]int8	// edgePermutation		(0-11)
	eO 		[12]int8	// edgeOrientation		(0-1)	0 = good, 1 = bad // bool?
	move	string		// last move
	move2	string		// move before last
}

const Reset		= "\x1B[0m"
// const White		= "\x1B[0m"					// 0 U
// const Orange	= "\x1B[38;2;255;165;0m"	// 1 L
const Green		= "\x1B[32m"				// 2 F
const Red		= "\x1B[31m"				// 3 R
// const Blue		= "\x1B[34m"				// 4 B
// const Yellow	= "\x1B[33m"				// 5 D



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


func parseArg() (string, bool, int, bool) {
	args := os.Args[1:]
	var random int = -1
	if len(args) == 0 {
		errorExit("not enough arguments, no mix given")
	} else if len(args) > 4 {
		errorExit("too many arguments")
	}
	mix := args[0]
	if mix == "-h" || mix == "--help" {
		printUsage()
	}
	visualizer := false
	group := false
	if len(args) > 1 {
		for i := 1; i < len(args); i++ {
			if args[i] == "-h" || args[i] == "--help" {
				printUsage()
			} else if args[i] == "-v" || args[i] == "--visualizer" {
				visualizer = true
			} else if args[i] == "-g" || args[i] == "--group" {
				group = true
			} else {
				length, err := strconv.Atoi(args[1])
				if (mix == "-r" || mix == "--random") && i == 1 && err == nil && length > 0 && length < 100 {
					random = length
				} else {
					fmt.Printf("Error bad argument: %v\n", args[i])
					printUsage()
				}
			}
		}
	}
	return mix, visualizer, random, group
}

func makeMix(mix string, length int) string {
	if mix == "-r" || mix == "--random" {
		return randomMix(length)
	} else if len(strings.Fields(mix)) == 1 {
		allMoves := listMoves(initCube(), 0)
		for _, move := range allMoves {
			if strings.Fields(mix)[0] == move {
				return mix
			}
		}
	}
	file, err := ioutil.ReadFile(mix)
	if err != nil {
		errorExit("failed to read mix file")
	}
	if len(file) > 200 {
		errorExit("file too long")
	}
	filepath := mix
	mix = string(file)
	fmt.Printf("\nMix read from filepath \"%v\":\n%v\n\n", filepath, mix)
	return mix
}

func initCube() *cepo {
	cepo := &cepo{}
	for i := range cepo.cP {
		cepo.cP[i] = int8(i)
	}
	for i := range cepo.eP {
		cepo.eP[i] = int8(i)
	}
	return cepo
}

// func dumpCube(cube *cepo) {
// 	fmt.Printf("\n\n#### -- CUBE -- ####\n")
// 	for i, corner := range cube.cP {
// 		fmt.Printf("Corner Permutation %v:\t%v\n", i, corner)//
// 	}
// 	fmt.Println()//
// 	for i, corner := range cube.cO {
// 		fmt.Printf("Corner Orientation %v:\t%v\n", i, corner)//
// 	}
// 	fmt.Println()//
// 	for i, edge := range cube.eP {
// 		fmt.Printf("Edge Permutation %v:\t%v\n", i, edge)//
// 	}
// 	fmt.Println()//
// 	for i, edge := range cube.eO {
// 		fmt.Printf("Edge Orientation %v:\t%v\n", i, edge)//
// 	}
// }

func isSolved(cube *cepo) bool {
	for i := range cube.cP {
		if cube.cP[i] != int8(i) {
			return false
		}
	}
	for i := range cube.cO {
		if cube.cO[i] != 0 {
			return false
		}
	}
	for i := range cube.eP {
		if cube.eP[i] != int8(i) {
			return false
		}
	}
	for i := range cube.eO {
		if cube.eO[i] != 0 {
			return false
		}
	}
	return true
}

func printSolution(solution string, elapsed time.Duration, cube *cepo) {
	if isSolved(cube) == false {
		fmt.Printf("%v\nError: Solution Incorrect :(%v\n", Red, Reset)
	} else {
		fmt.Printf("%v\nSolution Correct, cube solved! :)\n%v", Green, Reset)//
	}
	fmt.Printf("\nHalf Turn Metric: %v\n", halfTurnMetric(solution))
	fmt.Printf("\n%vSolution:\n%v%v\n\n", "\x1B[1m", "\x1B[0m", solution)
	fmt.Printf("Solve time:\n%v\n\n", elapsed)
}

func RunRubik2() {
	mix, visualizer, length, group := parseArg()
	mix = makeMix(mix, length)
	tables := makeTables()
	cube := initCube()
	spin(mix, cube)
	start := time.Now()
	solution := solve(cube, tables, group)
	elapsed := time.Since(start)
	printSolution(solution, elapsed, cube)
	runGraphic(mix, solution, visualizer)
}