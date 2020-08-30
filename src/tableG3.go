package rubik

import (
	"fmt"
	// "math"
)

// func ePindexConverter(cube *cepo) [3]int {
// 	// fmt.Printf("\nOH HIII!\n")
// 	// cube := initCube()
// 	// ePindex := 1
// 	var sliceIndex [3]int
// 	for slice := 0; slice < 3; slice++ {
// 		// var sliceIndex int
// 		for i := slice * 4; i < slice * 4 + 3; i++ {
// 			// fmt.Printf("cube.eP[i] : %v\n", cube.eP[i])//
// 			// sliceIndex += (int(cube.eP[i]) - slice * 4) * int(math.Pow(2, float64(2 - i + slice * 4)))
// 			sliceIndex[slice] += (int(cube.eP[i]) - slice * 4) * int(math.Pow(2, float64(2 - i + slice * 4)))
// 			// fmt.Printf("math.Pow(2, float64(3-i)): %v\n", math.Pow(2, float64(2-i + slice * 4)))//
// 			// fmt.Printf("index: %v\n", index)//
	
// 		}
// 		// ePindex *= sliceIndex
// 		// fmt.Printf("index slice: %v\n\n", index)//
// 	}
// 	// fmt.Printf("index done: %v\n", index)//
// 	// return ePindex
// 	return sliceIndex
// }

func ePindexConverter(cube *cepo) [3]uint8 {
	var sliceIndex [3]uint8
	if cube.eP[0] == 0 {
		if cube.eP[1] == 1 {
			if cube.eP[2] == 2 {
				sliceIndex[0] = 0 	//	0123
			} else { // cube.eP[2] = 3
				sliceIndex[0] = 1	//	0132
			}
		} else if cube.eP[1] == 2 {
			if cube.eP[2] == 1 {
				sliceIndex[0] = 2	//	0213
			} else { // cube.eP[2] = 3
				sliceIndex[0] = 3	//	0231
			}
		} else { // if cube.eP[1] == 3 {
			if cube.eP[2] == 1 {
				sliceIndex[0] = 4	//	0312
			} else { // cube.eP[2] = 2
				sliceIndex[0] = 5	//	0321
			}
		}
	} else if cube.eP[0] == 1 {
		if cube.eP[1] == 0 {
			if cube.eP[2] == 2 {
				sliceIndex[0] = 6 	//	1023
			} else { // cube.eP[2] = 3
				sliceIndex[0] = 7	//	1032
			}
		} else if cube.eP[1] == 2 {
			if cube.eP[2] == 0 {
				sliceIndex[0] = 8	//	1203
			} else { // cube.eP[2] = 3
				sliceIndex[0] = 9	//	1230
			}
		} else { // if cube.eP[1] == 3 {
			if cube.eP[2] == 0 {
				sliceIndex[0] = 10	//	1302
			} else { // cube.eP[2] = 2
				sliceIndex[0] = 11	//	1320
			}
		}
	} else if cube.eP[0] == 2 {
		if cube.eP[1] == 0 {
			if cube.eP[2] == 1 {
				sliceIndex[0] = 12 	//	2013
			} else { // cube.eP[2] = 3
				sliceIndex[0] = 13	//	2031
			}
		} else if cube.eP[1] == 1 {
			if cube.eP[2] == 0 {
				sliceIndex[0] = 14	//	2103
			} else { // cube.eP[2] = 3
				sliceIndex[0] = 15	//	2130
			}
		} else { // cube.eP[1] == 3 {
			if cube.eP[2] == 0 {
				sliceIndex[0] = 16	//	2301
			} else { // cube.eP[2] = 1
				sliceIndex[0] = 17	//	2310
			}
		}
	} else { // if cube.eP[0] == 3 {
		if cube.eP[1] == 0 {
			if cube.eP[2] == 1 {
				sliceIndex[0] = 18 	//	3012
			} else { // cube.eP[2] = 2
				sliceIndex[0] = 19 	//	3021
			}
		} else if cube.eP[1] == 1 {
			if cube.eP[2] == 0 {
				sliceIndex[0] = 20 	//	3102
			} else { // cube.eP[2] = 2
				sliceIndex[0] = 21 	//	3120
			}
		} else { // cube.eP[1] == 2 {
			if cube.eP[2] == 0 {
				sliceIndex[0] = 22	//	3201
			} else { // cube.eP[2] = 1
				sliceIndex[0] = 23	//	3210
			}
		}
	}
	
	fmt.Printf("cube.eP: %v\n", cube.eP)//
	fmt.Printf("sliceIndex[0]: %v\n", sliceIndex[0])//
	return sliceIndex
}

func cPtableIndex(tables *tables) {
	// var cPtableIndex [40320]uint8
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
					tables.G3cPindex[cP2index(child)] = converted
					converted++
				}
				children = append(children, *child)
			}
		}
		parents = children
	}
	// fmt.Printf("count: %v\n", count)//
	// return cPtableIndex
}


func tableG3(tables *tables) {
	fmt.Printf("\nGenerating pruning table for G3")
	var parents []cepo
	parents = append(parents, *initCube())
	var depth uint8
	var cumulative int//
	for depth < 15  {// 15
		depth++
		fmt.Printf("len(parents): %v\n", len(parents))//
		var count int//
		var children []cepo
		for _, parent := range parents {
			for _, move := range listMoves(&parent, 3) {
				// fmt.Printf("\nmove %v: %v\n", i, move)//
				child := newNode(&parent, move)
				spin(move, child)
				// dumpCube(child)//

				idxCP := tables.G3cPindex[cP2index(child)]
				idxEP := ePindexConverter(child)

				if tables.G3[idxCP][idxEP[0]][idxEP[1]][idxEP[2]] == 0 /*&& !(idxCP == 0 && idxEP[0] == 0 && idxEP[1] == 0 && idxEP[2] == 0)*/ {
					tables.G3[idxCP][idxEP[0]][idxEP[1]][idxEP[2]] = depth
					count++//
					cumulative++//
					children = append(children, *child)
				}
				// children = append(children, *child)
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
	cPtableIndex(tables)
	tableG3(tables)
	// cube := initCube()
	// ePindexConverter(cube)
	// spin("U2 D2 R2", cube)//
	// // cPindex := cP2index(cube)//
	// fmt.Printf("\ncPtableIndex[cPindex]: %v\n", cPtableIndex[cP2index(cube)])//
	// fmt.Printf("\nePindexConverter(cube): %v\n", ePindexConverter(cube))//
	// fmt.Printf("cPtableIndex: %v\n", cPtableIndex)//
}