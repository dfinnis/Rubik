package rubik

import (
	"fmt"
	"time"
)

// S =	cornerPermutation	edgePermutation
// 		cornerOrientation	edgeOrientation

// corner[0] is face U top left cubie
// edge[0] is face U left cubie


type cepo struct {
	cP 		[8]int8		// cornerPermutation	(0-7)
	cO 		[8]int8		// cornerOrientation	(0-2)	0 = good, 1 = twisted clockwise, 2 = twisted anti-clockwise
	eP 		[12]int8	// edgePermutation		(0-11)
	eO 		[12]int8	// edgeOrientation		(0-1)	0 = good, 1 = bad // bool?
	move	string		// last move
	move2	string		// move before last
}

// var cepo *cepo

func initCepo() *cepo {
	cepo := &cepo{}
	for i := range cepo.cP {
		cepo.cP[i] = int8(i)
	}
	for i := range cepo.eP {
		cepo.eP[i] = int8(i)
	}
	return cepo
}

// func cornerFacelet(cepo *cepo, idx uint8, face uint8) (color uint8) {
// 	// fmt.Printf("OH HIIII")//
// 	permutation := cepo.cP[idx]
// 	fmt.Printf("permutation: %v\n", permutation)
// 	return 0
// }

//func edgeFacelet(cepo *cepo, idx uint8, face uint8) (color uint8) {
//}

func dumpCepo(cepo *cepo) {
	fmt.Printf("\n\n#### -- CUBE -- ####\n")


	
	// dumpFace(cube, 0)
	// dumpLFRB(cube)
	// dumpFace(cube, 5)
	// fmt.Printf("\n        ")

	// if cornerFacelet(cepo, 0, 0) == 5 {
	// 	fmt.Printf("%v5%v ", Yellow, Reset)
	// } else if cornerFacelet(cepo, 0, 0) == 4 {
	// 	fmt.Printf("%v4%v ", Blue, Reset)
	// } else if cornerFacelet(cepo, 0, 0) == 3 {
	// 	fmt.Printf("%v3%v ", Red, Reset)
	// } else if cornerFacelet(cepo, 0, 0) == 2 {
	// 	fmt.Printf("%v2%v ", Green, Reset)
	// } else if cornerFacelet(cepo, 0, 0) == 1 {
	// 	fmt.Printf("%v1%v ", Orange, Reset)
	// } else {
	// 	fmt.Printf("0 ")
	// }
	// fmt.Printf("%v")

	fmt.Printf("\n\n")//
	for i, corner := range cepo.cP {
		fmt.Printf("Corner Permutation %v:\t%v\n", i, corner)//
	}
	fmt.Println()//
	for i, corner := range cepo.cO {
		fmt.Printf("Corner Orientation %v:\t%v\n", i, corner)//
	}
	fmt.Println()//
	for i, edge := range cepo.eP {
		fmt.Printf("Edge Permutation %v:\t%v\n", i, edge)//
	}
	fmt.Println()//
	for i, edge := range cepo.eO {
		fmt.Printf("Edge Orientation %v:\t%v\n", i, edge)//
	}
}

func isSolvedCepo(cube *cepo) bool {
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

func printSolution2(solution string, elapsed time.Duration, cube *cepo) {
	fmt.Printf("\n########################################\n")
	if isSolvedCepo(cube) == false {
		fmt.Printf("%v\nError: Solution Incorrect :(%v\n", Red, Reset)
	} else {//
		fmt.Printf("%v\nSolution Correct, cube solved! :)\n%v", Green, Reset)//
	}//
	fmt.Printf("\nHalf Turn Metric: %v\n", halfTurnMetric(solution))
	fmt.Printf("\n%vSolution:\n%v%v\n\n", "\x1B[1m", "\x1B[0m", solution)
	fmt.Printf("Solve time:\n%v\n\n", elapsed)
}

func RunRubik2() {
	mix, visualizer, length := parseArg()
	if mix == "-r" || mix == "--random" {
		mix = randomMix(length)
	}
	tableG0 := tableGenerator()
	// fmt.Printf("tableG0: %v\n", tableG0)//
	cube := initCepo()
	// dumpCepo(cube)//
	spinCepo(mix, cube)
	// dumpCepo(cube)//

	start := time.Now()
	solution := solveCepo(cube, tableG0)
	// solution := "U U"//
	elapsed := time.Since(start)
	spinCepo(solution, cube)
	fmt.Printf("========================================\n")//
	dumpCepo(cube)//
	printSolution2(solution, elapsed, cube)
	runGraphic(mix, solution, visualizer)
}