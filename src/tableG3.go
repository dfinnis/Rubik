package rubik

import (
	"fmt"
	"math"
)

func ePindexConverter(cube *cepo) int {
	// fmt.Printf("\nOH HIII!\n")
	// cube := initCube()
	ePindex := 1
	for slice := 0; slice < 3; slice++ {
		var sliceIndex int
		for i := slice * 4; i < slice * 4 + 3; i++ {
			// fmt.Printf("cube.eP[i] : %v\n", cube.eP[i])//
			sliceIndex += (int(cube.eP[i]) - slice * 4) * int(math.Pow(2, float64(2 - i + slice * 4)))
			// fmt.Printf("math.Pow(2, float64(3-i)): %v\n", math.Pow(2, float64(2-i + slice * 4)))//
			// fmt.Printf("index: %v\n", index)//
	
		}
		ePindex *= sliceIndex
		// fmt.Printf("index slice: %v\n\n", index)//
	}
	// fmt.Printf("index done: %v\n", index)//
	return ePindex
}

func cPtableIndex() [40320]uint8 {
	var cPtableIndex [40320]uint8
	var initial []cepo
	var parents []cepo
	var converted uint8 = 1
	initial = append(initial, *initCube())
	parents = append(parents, *initCube())
	for depth := 0; depth < 4; depth++ {
		var children []cepo
		for _, parent := range parents {
			for _, move := range listMoves(&parent, 2) {
				child := newNode(&parent, move)
				spin(move, child)
				if cornersInOrbit(child) == true && cPinList(child, initial) == false {
					initial = append(initial, *child)
					cPtableIndex[cP2index(child)] = converted
					converted++
				}
				children = append(children, *child)
			}
		}
		parents = children
	}
	// fmt.Printf("count: %v\n", count)//
	return cPtableIndex
}

func makeTableG3(tables *tables) {
	fmt.Printf("\nGenerating pruning table for G3")
	cPtableIndex := cPtableIndex()
	cube := initCube()
	ePindexConverter(cube)
	spin("U2 D2 R2", cube)//
	// cPindex := cP2index(cube)//
	fmt.Printf("\ncPtableIndex[cPindex]: %v\n", cPtableIndex[cP2index(cube)])//
	fmt.Printf("\nePindexConverter(cube): %v\n", ePindexConverter(cube))//
	// fmt.Printf("cPtableIndex: %v\n", cPtableIndex)//
}