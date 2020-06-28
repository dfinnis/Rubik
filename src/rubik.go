package rubik

import (
	"fmt"
	"os"
	"math/rand"
	"math/bits"
	"time"
	// "strconv"//
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
	r = &rubik{}
	r.cube[1] = 286331153//		0001
	r.cube[2] = 572662306//		0010
	r.cube[3] = 858993459//		0011
	r.cube[4] = 1145324612//	0100
	r.cube[5] = 1431655765//	0101
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
	// fmt.Println(reflect.TypeOf(moveList))
	return mix
}

// func solve(&r.cube) string {
	
	// return solution
// }

// func printSolution(solution) {

// }

// randomMix returns a random 20 to 40 spin long mix
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
	n := rand.Intn(20) + 20
	for i := 0; i <= n; i++ {
		mix += 	spin[rand.Intn(len(spin))]
		if i != n {
			mix += " "
		}
	}
	return mix
}


func test() {//////
	fmt.Printf("test!\n")/////////
	var face uint32
	face = 1
	fmt.Printf("\nint before: %v\n", face)/////////
	face = bits.RotateLeft32(face, 2)
	fmt.Printf("\nint after: %v\n\n", face)/////////
	face = bits.RotateLeft32(face, -3)
	fmt.Printf("\nint after: %v\n\n", face)/////////
	fmt.Printf("test end!\n")/////////
}

func dumpCube(cube *[6]uint32) {
	fmt.Printf("r: %b\n", cube)
	fmt.Printf("face 0: %.32b\n", cube[0])
	fmt.Printf("face 1: %.32b\n", cube[1])
	fmt.Printf("face 2: %.32b\n", cube[2])
	fmt.Printf("face 3: %.32b\n", cube[3])
	fmt.Printf("face 4: %.32b\n", cube[4])
	fmt.Printf("face 5: %.32b\n", cube[5])
	
}

func RunRubik() {
	// test()//////

	mix := parseArg()
	// random := randomMix()//// make option, flag -r?
	// fmt.Printf("random: %s\n", random)//
	fmt.Printf("mix: %s\n", mix)//
	r := initRubik()
	dumpCube(&r.cube)////
	// spin(mix, r)
	// solution := solve(&r.cube)
	// printSolution(solution)
	// runGraphic()
}
