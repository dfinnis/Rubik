package rubik

import (
	"fmt"
	"os"
	"math/rand"
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

func initRubik() *rubik {
	r = &rubik{}			//	0000 0000 0000 0000 0000 0000 0000 0000
	// r.cube[0] = 0x44444444	//													rm!!!!!!!!!!!!!!!!!	
	r.cube[1] = 0x11111111	//	286331153	//	0001 0001 0001 0001 0001 0001 0001 0001
	r.cube[2] = 0x22222222	//	572662306	//	0010 0010 0010 0010 0010 0010 0010 0010
	r.cube[3] = 0x33333333	//	858993459	//	0011 0011 0011 0011 0011 0011 0011 0011
	r.cube[4] = 0x44444444	//	1145324612	//	0100 0100 0100 0100 0100 0100 0100 0100
	r.cube[5] = 0x55555555	//	1431655765	//	0101 0101 0101 0101 0101 0101 0101 0101
	return r
}

func parseArg() string {
	args := os.Args[1:]
	if len(args) == 0 {
		errorExit("not enough arguments, no mix given")
	} else if len(args) > 1 {
		errorExit("too many arguments")
	}
	mix := args[0]
	// random := false
	// if len(args) == 2 {
	// 	if args[1] != "-r" {
	// 		errorExit("make a proper usage func")//////!!!!!!!
	// 	}
	// 	// fmt.Printf("args[1]: %v\n\n", args[1])//
	// 	random = true
	// }
	// fmt.Println(reflect.TypeOf(moveList))
	return mix
}

// randomMix returns a random 20 to 24 spin long mix
func randomMix() string {
	var mix string
	spin := []string{
		"U",
		"U'",
		"U2",
		"D",
		"D'",
		"D2",
		"R",
		"R'",
		"R2",
		"L",
		"L'",
		"L2",
		"F",
		"F'",
		"F2",
		"B",
		"B'",
		"B2",
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(4) + 20
	for i := 0; i <= n; i++ {
		mix += 	spin[rand.Intn(len(spin))]
		if i != n {
			mix += " "
		}
	}
	return mix
}

func printSolution(solution string) {
	fmt.Printf("\nSolution: %v\n\n", solution)
}

func RunRubik() {
	mix := parseArg()
	if mix == "-r" || mix == "--random" {
		mix = randomMix()
	}
	fmt.Printf("mix: %s\n", mix)//
	r := initRubik()
	dumpCube(&r.cube)////
	spin(mix, &r.cube)
	// dumpCube(&r.cube)////
	// solution := solve(&r.cube)
	// solution := solveHuman(&r.cube)
	solution := "F U"
	printSolution(solution)
	runGraphic(mix, solution)
}

// a & 196	query a value for its set bits
// &=		selectively clearing bits of an integer value to zero
// |=		set arbitrary bits for a given integer value