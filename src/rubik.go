package rubik

import (
	"fmt"
	"os"
	"time"
)

// rubik struct contains all information about current rubik state
type rubik struct {
	cube	[6]uint32
}

// rubik var r contains all information about current rubik state
var r *rubik

func errorExit(message string) {
	fmt.Printf("Error: %s\n", message)
	os.Exit(1)
}

func printUsage() {
	fmt.Printf("\nUsage:\tgo build; ./Rubik \"mix\" [-v] [-h]\n\n")
	fmt.Printf("    mix should be valid sequence string e.g.\n")
	fmt.Printf("    \"U U' U2 D D' D2 R R' R2 L L' L2 F F' F2 B B' B2\"\n")
	fmt.Printf("    or mix \"$(< mix/superflip.txt)\" reads a file\n")
	fmt.Printf("    or mix \"-r\" or \"--random\" mixes randomly\n\n")
	fmt.Printf("    [-v] (--visualizer) show visual of mix and solution\n")
	fmt.Printf("    [-h] (--help) show usage\n\n")
	os.Exit(1)
}

func initRubik() *rubik {
	r = &rubik{}			//	0			//	0000 0000 0000 0000 0000 0000 0000 0000
	r.cube[1] = 0x11111111	//	286331153	//	0001 0001 0001 0001 0001 0001 0001 0001
	r.cube[2] = 0x22222222	//	572662306	//	0010 0010 0010 0010 0010 0010 0010 0010
	r.cube[3] = 0x33333333	//	858993459	//	0011 0011 0011 0011 0011 0011 0011 0011
	r.cube[4] = 0x44444444	//	1145324612	//	0100 0100 0100 0100 0100 0100 0100 0100
	r.cube[5] = 0x55555555	//	1431655765	//	0101 0101 0101 0101 0101 0101 0101 0101
	return r
}

func parseArg() (string, bool) {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Error: not enough arguments, no mix given\n")
		printUsage()
	} else if len(args) > 3 {
		fmt.Printf("Error: too many arguments\n")
		printUsage()
	}
	mix := args[0]
	if mix == "-h" || mix == "--help" {
		printUsage()
	}
	visualizer := false
	// debug := false
	// binary := false
	if len(args) > 1 {
		for i := 1; i < len(args); i++ {
			if args[i] == "-v" || args[i] == "--visualizer" {
				visualizer = true
			// } else if args[i] == "-d" || args[i] == "--debug" {
			// 	debug = true
			// } else if args[i] == "-b" || args[i] == "--binary" {
			// 	debug = true
			// 	binary = true
			} else {
				fmt.Printf("Error: bad argument\n")
				printUsage()
			}
		}
	}
	return mix, visualizer
}

func isSolved(cube *[6]uint32) bool {
	if cube[0]&0x77777777 == 0 &&
	cube[1]&0x77777777 == 0x11111111 &&
	cube[2]&0x77777777 == 0x22222222 && 
	cube[3]&0x77777777 == 0x33333333 && 
	cube[4]&0x77777777 == 0x44444444 && 
	cube[5]&0x77777777 == 0x55555555 {
		return true
	}
	return false
}

func printSolution(solution string, elapsed time.Duration, cube *[6]uint32) {
	spin(solution, cube)
	if isSolved(cube) == false {
		fmt.Printf("\nError: Solution incorrect :(\n")
	}
	fmt.Printf("\n%vSolution:\n%v%v\n\n", "\x1B[1m", "\x1B[0m", solution)
	fmt.Printf("Solve time:\n%v\n\n", elapsed)
}

func solvePlaceHolder() string { /////rm!!!!!!
	// solution := "F U"
	solution := randomMix()
	// time.Sleep(100000000)//
	return solution
}

func RunRubik() {
	mix, visualizer := parseArg()
	if mix == "-r" || mix == "--random" {
		mix = randomMix()
	}
	r := initRubik()
	spin(mix, &r.cube)
	// dumpCube(&r.cube)////
	start := time.Now()
	// solution := solve(&r.cube)
	solution := solvePlaceHolder()//rm!!!
	elapsed := time.Since(start)
	printSolution(solution, elapsed, &r.cube)
	runGraphic(mix, solution, visualizer)
}

// a & 196	query a value for its set bits
// &=		selectively clearing bits of an integer value to zero
// |=		set arbitrary bits for a given integer value


// implementation of kociemba
// how to find len(to G1)