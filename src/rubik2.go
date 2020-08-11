package rubik

import (
	"fmt"
)

// S =	cornerPermutation	edgePermutation
// 		cornerOrientation	edgeOrientation

type cepo struct {
	cP [12]uint8	// cornerPermutation	(0-11)
	cO [12]uint8	// cornerOrientation	(0-2)	0 = good, 1 = twisted clockwise, 2 = twisted anti-clockwise
	eP [8]uint8		// edgePermutation		(0-7)
	eO [8]uint8		// edgeOrientation		(0-1)	0 = good, 1 = bad
}

// var cepo *cepo

func initCepo() *cepo {
	cepo := &cepo{}
	for i := range cepo.cP {
		cepo.cP[i] = uint8(i)
	}
	for i := range cepo.eP {
		cepo.eP[i] = uint8(i)
	}
	return cepo
}

func dumpCepo(cepo *cepo) {
	fmt.Printf("oh hi!\n")//
	// fmt.Printf(cepo)//
	for i, corner := range cepo.cP {
		fmt.Printf("Corner Permutation %v:\t%v\n", i, corner)//
	}
	fmt.Println()//
	for i, corner := range cepo.cO {
		fmt.Printf("Corner Orientation %v:\t%v\n", i, corner)//
	}
	fmt.Println()//
	for i, edge := range cepo.eP {
		fmt.Printf("Corner Orientation %v:\t%v\n", i, edge)//
	}
	fmt.Println()//
	for i, edge := range cepo.eO {
		fmt.Printf("Corner Orientation %v:\t%v\n", i, edge)//
	}
}

func RunRubik2() {
	cepo := initCepo()
	dumpCepo(cepo)
	// fmt.Printf(cepo)
}