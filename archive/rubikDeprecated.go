package rubik

// import (
// 	"fmt"
// 	"os"
// 	// "reflect"//
// 	"math/rand"
// 	"time"
// )

// type face struct {
// 	pieces	[3][3]uint8
// }

// // rubik struct contains all information about current rubik state
// type rubik struct {
// 	cube	[6]face
// }

// // rubik var r contains all information about current rubik state
// var r *rubik

// func errorExit(message string) {
// 	fmt.Printf("Error: %s\n", message)
// 	os.Exit(1)
// }

// func initRubik() *rubik {
// 	r = &rubik{}
// 	var face uint8
// 	for face = 0; face < 6; face++ {
// 		for y := 0; y < 3; y++ {
// 			for x := 0; x < 3; x++ {
// 				r.cube[face].pieces[y][x] = face
// 			}
// 		}
// 	}
// 	return r
// }

// func parseArg() string {
// 	args := os.Args[1:]
// 	if len(args) == 0 {
// 		errorExit("not enough arguments, no mix given")
// 	} else if len(args) > 1 {
// 		errorExit("too many arguments")
// 	}
// 	mix := args[0]
// 	// fmt.Println(reflect.TypeOf(moveList))
// 	return mix
// }

// // func solve(&r.cube) string {
	
// 	// return solution
// // }

// // func printSolution(solution) {

// // }

// // randomMix returns a random 20 to 40 spin long mix
// func randomMix() string {
// 	var mix string
// 	spin := []string{
// 		"U",
// 		"U'",
// 		"U2",
// 		"D",
// 		"D'",
// 		"D2",
// 		"R",
// 		"R'",
// 		"R2",
// 		"L",
// 		"L'",
// 		"L2",
// 		"F",
// 		"F'",
// 		"F2",
// 		"B",
// 		"B'",
// 		"B2",
// 	}
// 	rand.Seed(time.Now().UnixNano())
// 	n := rand.Intn(20) + 20
// 	for i := 0; i <= n; i++ {
// 		mix += 	spin[rand.Intn(len(spin))]
// 		if i != n {
// 			mix += " "
// 		}
// 	}
// 	return mix
// }

// func RunRubikDeprecated() {
// 	mix := parseArg()
// 	// mix := randomMix()
// 	random := randomMix()//// make option, flag -r?
// 	fmt.Printf("random: %s\n", random)//
// 	fmt.Printf("mix: %s\n", mix)//
// 	r := initRubik()
// 	dumpCube(&r.cube)////
// 	spin(mix, r)
// 	// spin(random, r)
// 	// solution := solve(&r.cube)
// 	// printSolution(solution)
// 	// spin(solution, r)
// 	// rubik.runGraphic()
// 	// dumpCube(&r.cube)////
// }
