package rubik

import (
	"os"
	"fmt"
	"math"
	"strconv"
)

func cO2index(cO [8]int8) int {
	var index int
	for i := 0; i < 7; i++ {
		index = index * 3
		index = index + int(cO[i])
	}
	return index
}

// // index2cO reverses cO2index, accepting 0-2186 returning [00000000]-[22222221] // debug tool
// func index2cO(index int) [8]int8 {
// 	var s int
// 	var or [8]int
// 	for i := 6; i >= 0; i-- {
// 		or[i] = index % 3
// 		s = s - or[i]
// 		if s < 0 {
// 			s = s + 3
// 		}
// 		index = (index - or[i]) / 3
// 	}
// 	or[7] = s
// 	var or8 [8]int8
// 	for i := range or8 {
// 		or8[i] = int8(or[i])
// 	}
// 	return or8
// }

// func testcO2index() {
// 	for i := 0; i <= 2186; i++{
// 		cO := index2cO(i)
// 		index := cO2index(cO)
// 		fmt.Printf("cO: %v: %v\n", i, cO)
// 		fmt.Printf("index: %v\n", index)
// 		if index != i {
// 			fmt.Printf("Error cO2index function\n")
// 		}
// 	}
// }

func eP2Binary(cube *cepo) [11]bool {
	var binary [11]bool
	for i := 0; i < 11; i++ {
		if cube.eP[i] > 7 {
			binary[i] = true
		}
	}
	return binary
}

func binaryBool2Decimal(binary [11]bool) int {
	var decimal int
	for i := 0; i < 11; i++ {
		if binary[i] == true {
			decimal += int(math.Pow(2, float64(10-i)))
		}
	}
	return decimal
}

func eP2index(cube *cepo, tables *tables) int16 {
	ePbinary := eP2Binary(cube)
	idxEP := binaryBool2Decimal(ePbinary)
	return tables.G1ePindex[idxEP]
}

// debug tool
func index2eP(index int16, tables *tables) [12]bool {
	var eP [12]bool
	for i, entry := range tables.G1ePindex {
		if entry == index && entry != 0 {
			fmt.Printf("entry: %v\n", entry)
			fmt.Printf("i: %v\n", i)
			binary := strconv.FormatInt(int64(i), 2)
			fmt.Printf("binary: %v\n", binary)
		}
	}
	return eP
}

func tableG1IdxConv(tables *tables) {
	var converted int16
	var idx int64
	for idx = 0; idx < 2048; idx++ {
		var count uint8
		binary := strconv.FormatInt(idx, 2)
		for _, bit := range binary {
			if bit == '1' {
				count++
			}
		}
		if count == 4 || count == 3 {
				tables.G1ePindex[idx] = converted
				converted++
		}
	}
}

func tableG1(tables *tables) {
	fmt.Printf("\nGenerating pruning table for G1")
	var parents []cepo
	parents = append(parents, *initCube())
	var cumulative int//
	var depth uint8
	for depth < 8 {// 9 !!!!
		depth++
		var count int//
		var children []cepo
		for _, parent := range parents {
			for _, move := range listMoves(&parent, 1) {
				// fmt.Printf("\nmove %v: %v\n", i, move)//
				child := newNode(&parent, move)
				spin(move, child)
				// dumpCube(child)//

				idxCO := cO2index(child.cO)
				idxEP := eP2index(child, tables)

				// if tables.G1[idxEP][idxCO] != 0 {
				// 	fmt.Printf("tables.G1[%v][%v]: %v\n", idxEP, idxCO, tables.G1[idxEP][idxCO])//
				// }
				if tables.G1[idxEP][idxCO] == 0 && !(idxEP == 0 && idxCO == 0) {
					tables.G1[idxEP][idxCO] = depth
					// if depth < 2 {
					// 	children = append(children, *child)//!!!!!!!
					// }
					count++//
					cumulative++//
				// } else if tables.G1[idxEP][idxCO] != 0 {
				// // } else if (idxEP == 0 && idxCO == 0) && isSolved(child) == false {//
				// 	children = append(children, *child)
					// fmt.Printf("child.move: %v\n", child.move)//
					// fmt.Printf("child.move2: %v\n", child.move2)//
					// fmt.Printf("OH MYYYY!!!\n")//
					// dumpCube(child)
				// } else if !(idxEP == 0 && idxCO == 0) {//
					// fmt.Printf("tables.G1[%v][%v]: %v\n", idxEP, idxCO, tables.G1[idxEP][idxCO])//
				}
				children = append(children, *child)//
			}
		}
		parents = children
		fmt.Printf(".")
		// fmt.Printf("depth: %v\n", depth)//
		// fmt.Printf("count: %v\n", count)//
		// fmt.Printf("cumulative: %v\n", cumulative)//
		// fmt.Printf("len(parents): %v\n\n", len(parents))//
	}

	var count int//
	for ePidx := 0; ePidx < 495; ePidx++ {
		for cOidx := 0; cOidx < 2187; cOidx++ {
			if tables.G1[ePidx][cOidx] == 0 && !(ePidx == 0 && cOidx == 0) {
				tables.G1[ePidx][cOidx] = 9 // 10 !!!??
				count++//
				cumulative++//
			}
		}
	}
	// fmt.Printf("depth: %v\n", 9)//
	// fmt.Printf("count: %v\n", count)//
	// fmt.Printf("cumulative: %v\n\n", cumulative)//
	// fmt.Printf("\n###########################################################################\n")//
}

func makeTableG1(tables *tables) {
	tableG1IdxConv(tables)
	// testcO2index()
	if _, err := os.Stat("tables/G1.txt"); os.IsNotExist(err) {
		tableG1(tables)
		file := createFile("tables/G1.txt")
		defer file.Close()
		for ePidx := 0; ePidx < 495; ePidx++ {
			for cOidx := 0; cOidx < 2187; cOidx++ {
				_, err = file.WriteString(fmt.Sprintf("%x", tables.G1[ePidx][cOidx]))
				if err != nil {
					errorExit("failed to write to file")
				}
			}
		}
	} else {
		file := readFile("tables/G1.txt")
		ePidx := 0
		cOidx := 0
		for _, depth := range file {
			tables.G1[ePidx][cOidx] = readHex(depth)
			cOidx++
			if cOidx >= 2187 {
				cOidx = 0
				ePidx++
			}
		}
	}
}