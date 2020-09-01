package rubik

import (
	"fmt"
	"os"
)

func ePindexSlice(cube *cepo, slice uint8) uint8 {
	slice4 := int8(slice * 4)
	if cube.eP[0 + slice4] == 0 + slice4 {
		if cube.eP[1 + slice4] == 1 + slice4 {
			if cube.eP[2 + slice4] == 2 + slice4 {
				return 0 			//	0123
			} else {
				return 1			//	0132
			}
		} else if cube.eP[1 + slice4] == 2 + slice4 {
			if cube.eP[2 + slice4] == 1 + slice4 {
				return 2			//	0213
			} else {
				return 3			//	0231
			}
		} else {
			if cube.eP[2 + slice4] == 1 + slice4 {
				return 4			//	0312
			} else {
				return 5			//	0321
			}
		}
	} else if cube.eP[0 + slice4] == 1 + slice4 {
		if cube.eP[1 + slice4] == 0 + slice4 {
			if cube.eP[2 + slice4] == 2 + slice4 {
				return 6 			//	1023
			} else {
				return 7			//	1032
			}
		} else if cube.eP[1 + slice4] == 2 + slice4 {
			if cube.eP[2 + slice4] == 0 + slice4 {
				return 8			//	1203
			} else {
				return 9			//	1230
			}
		} else {
			if cube.eP[2 + slice4] == 0 + slice4 {
				return 10			//	1302
			} else {
				return 11			//	1320
			}
		}
	} else if cube.eP[0 + slice4] == 2 + slice4 {
		if cube.eP[1 + slice4] == 0 + slice4 {
			if cube.eP[2 + slice4] == 1 + slice4 {
				return 12 			//	2013
			} else {
				return 13			//	2031
			}
		} else if cube.eP[1 + slice4] == 1 + slice4 {
			if cube.eP[2 + slice4] == 0 + slice4 {
				return 14			//	2103
			} else {
				return 15			//	2130
			}
		} else {
			if cube.eP[2 + slice4] == 0 + slice4 {
				return 16			//	2301
			} else {
				return 17			//	2310
			}
		}
	} else {
		if cube.eP[1 + slice4] == 0 + slice4 {
			if cube.eP[2 + slice4] == 1 + slice4 {
				return 18 			//	3012
			} else {
				return 19 			//	3021
			}
		} else if cube.eP[1 + slice4] == 1 + slice4 {
			if cube.eP[2 + slice4] == 0 + slice4 {
				return 20 			//	3102
			} else {
				return 21 			//	3120
			}
		} else {
			if cube.eP[2 + slice4] == 0 + slice4 {
				return 22			//	3201
			} else {
				return 23			//	3210
			}
		}
	}
}

func ePindexConverter(cube *cepo) [3]uint8 {
	var sliceIndex [3]uint8
	var slice uint8
	for slice = 0; slice < 3; slice++ {
		sliceIndex[slice] = ePindexSlice(cube, slice)
	}
	return sliceIndex
}

func cPtableIndex(tables *tables) {
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
}

func tableG3(tables *tables) {
	fmt.Printf("\nGenerating pruning table for G3")
	var parents []cepo
	parents = append(parents, *initCube())
	var depth uint8
	for depth < 15  {
		depth++
		var children []cepo
		for _, parent := range parents {
			for _, move := range listMoves(&parent, 3) {
				child := newNode(&parent, move)
				spin(move, child)
				idxCP := tables.G3cPindex[cP2index(child)]
				idxEP := ePindexConverter(child)
				if tables.G3[idxCP][idxEP[0]][idxEP[1]][idxEP[2]] == 0 {
					tables.G3[idxCP][idxEP[0]][idxEP[1]][idxEP[2]] = depth
					children = append(children, *child)
				}
			}
		}
		parents = children
		fmt.Printf(".")
	}
}

func makeTableG3(tables *tables) {
	cPtableIndex(tables)
	if _, err := os.Stat("tables/G3.txt"); os.IsNotExist(err) {
		tableG3(tables)
		file := createFile("tables/G3.txt")
		defer file.Close()
		for cPidx := 0; cPidx < 96; cPidx++ {
			for ePidx0 := 0; ePidx0 < 24; ePidx0++ {
				for ePidx1 := 0; ePidx1 < 24; ePidx1++ {
					for ePidx2 := 0; ePidx2 < 24; ePidx2++ {
						_, err = file.WriteString(fmt.Sprintf("%x", tables.G3[cPidx][ePidx0][ePidx1][ePidx2]))
						if err != nil {
							errorExit("failed to write to file")
						}
					}
				}
			}
		}
	} else {
		file := readFile("tables/G3.txt")
		cPidx := 0
		ePidx0 := 0
		ePidx1 := 0
		ePidx2 := 0
		for _, depth := range file {
			tables.G3[cPidx][ePidx0][ePidx1][ePidx2] = readHex(depth)
			ePidx2++
			if ePidx2 >= 24 {
				ePidx2 = 0
				ePidx1++
				if ePidx1 >= 24 {
					ePidx1 = 0
					ePidx0++
					if ePidx0 >= 24 {
						ePidx0 = 0
						cPidx++
					}
				}
			}
		}
	}
}