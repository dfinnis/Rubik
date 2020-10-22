package rubik

import (
	"fmt"
	"time"
	"os"
	"strconv"
)

// cepo struct describes a cube state. Corner, Edge, Permutation, Orientation
type cepo struct {
	cP 		[8]int8		// cornerPermutation	(0-7)
	cO 		[8]int8		// cornerOrientation	(0-2)	0 = good, 1 = twisted clockwise, 2 = twisted anti-clockwise
	eP 		[12]int8	// edgePermutation		(0-11)
	eO 		[12]int8	// edgeOrientation		(0-1)	0 = good, 1 = bad
	move	string		// last move
	move2	string		// move before last
}

// parseArg parses arguments, returns mix and flags
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

// initCube returns a new solved cube
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

// RunRubik is the main and only exposed function
func RunRubik() {
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
