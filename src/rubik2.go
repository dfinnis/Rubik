package rubik

import (
	"fmt"
)

// S =	cornerPermutation	edgePermutation
// 		cornerOrientation	edgeOrientation

// corner[0] is face U top left cubie
// edge[0] is face U left cubie


type cepo struct {
	cP [8]int8			// cornerPermutation	(0-7)
	cO [8]int8			// cornerOrientation	(0-2)	0 = good, 1 = twisted clockwise, 2 = twisted anti-clockwise
	eP [12]int8		// edgePermutation		(0-11)
	eO [12]int8		// edgeOrientation		(0-1)	0 = good, 1 = bad // bool?
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

func RunRubik2() {
	mix, visualizer, length := parseArg()
	if mix == "-r" || mix == "--random" {
		mix = randomMix(length)
	}
	cube := initCepo()
	dumpCepo(cube)//
	// fmt.Printf(cepo)
	spinCepo(mix, cube)
	dumpCepo(cube)//

	// start := time.Now()
	// solution := solve(r)
	solution := "F U"
	// elapsed := time.Since(start)
	// printSolution(solution, elapsed, &r.cube)
	runGraphic(mix, solution, visualizer)
}