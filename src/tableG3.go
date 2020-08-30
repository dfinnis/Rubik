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


func tableG3(tables *tables, cPtableIndex [40320]uint8) {
	fmt.Printf("\nGenerating pruning table for G3")
	var parents []cepo
	parents = append(parents, *initCube())
	var depth uint8
	var cumulative int//
	for depth < 15 {
		depth++
		var count int//
		var children []cepo
		for _, parent := range parents {
			for _, move := range listMoves(&parent, 3) {
				// fmt.Printf("\nmove %v: %v\n", i, move)//
				child := newNode(&parent, move)
				spin(move, child)
				// dumpCube(child)//

				idxCP := cPtableIndex[cP2index(child)]
				idxEP := ePindexConverter(child)

				if tables.G3[idxCP][idxEP] == 0 && !(idxCP == 0 && idxEP == 0) {
					tables.G3[idxCP][idxEP] = depth
					count++//
					cumulative++//
					// children = append(children, *child)
				}
				children = append(children, *child)
			}
		}
		parents = children
		fmt.Printf(".")
		fmt.Printf("depth: %v\n", depth)//
		fmt.Printf("count: %v\n", count)//
		fmt.Printf("cumulative: %v\n\n", cumulative)//
		// // fmt.Printf("len(parents): %v\n", len(parents))//
		// // fmt.Printf("tables.G3[0][0]: %v\n\n", tables.G3[0][0])//
		// // fmt.Printf("tables.G3[0][1]: %v\n\n", tables.G3[0][1])//
		// // fmt.Printf("tables.G3[0][2]: %v\n\n", tables.G3[0][2])//
		// // fmt.Printf("tables.G3[40319][69]: %v\n\n", tables.G3[40319][69])//
	}
}



func makeTableG3(tables *tables) {
	fmt.Printf("\nGenerating pruning table for G3")
	cPtableIndex := cPtableIndex()
	tableG3(tables, cPtableIndex)
	// cube := initCube()
	// ePindexConverter(cube)
	// spin("U2 D2 R2", cube)//
	// // cPindex := cP2index(cube)//
	// fmt.Printf("\ncPtableIndex[cPindex]: %v\n", cPtableIndex[cP2index(cube)])//
	// fmt.Printf("\nePindexConverter(cube): %v\n", ePindexConverter(cube))//
	// fmt.Printf("cPtableIndex: %v\n", cPtableIndex)//
}