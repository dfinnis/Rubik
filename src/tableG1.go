package rubik

import (
	"os"
	"fmt"
	"math"
	"strconv"
)

func cO2index(cube *cepo) int {
	var index int
	for i := 0; i < 7; i++ {
		index = index * 3
		index = index + int(cube.cO[i])
	}
	return index
}

func eP2Binary(cube *cepo) [12]bool {
	var binary [12]bool
	for i := 0; i < 12; i++ {
		if cube.eP[i] > 7 {
			binary[i] = true
		}
	}
	return binary
}

func binaryBool2Decimal(binary [12]bool) int {
	var decimal int
	for i := 0; i < 12; i++ {
		if binary[i] == true {
			decimal += int(math.Pow(2, float64(11-i)))
		}
	}
	return decimal
}

func eP2index(cube *cepo, tables *tables) int16 {
	ePbinary := eP2Binary(cube)
	idxEP := binaryBool2Decimal(ePbinary)
	return tables.G1ePindex[idxEP]
}

// func index2orientation(index int) [8]int {
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
// 	return or
// }

func tableG1IdxConv(tables *tables) { // make file/read from file?
	var converted int16// = 1
	var idx int64
	for idx = 0; idx <4096; idx++ {
		var count uint8
		binary := strconv.FormatInt(idx, 2)
		for _, bit := range binary {
			if bit == '1' {
				count++
			}
		}
		if count == 4 {
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

				idxCO := cO2index(child)
				idxEP := eP2index(child, tables)

				if tables.G1[idxEP][idxCO] == 0 && !(idxEP == 0 && idxCO == 0) {
					tables.G1[idxEP][idxCO] = depth
					count++//
					cumulative++//
					// children = append(children, *child)
				}
				children = append(children, *child)
			}
		}
		parents = children
		fmt.Printf(".")
		// fmt.Printf("depth: %v\n", depth)//
		// fmt.Printf("count: %v\n", count)//
		// fmt.Printf("cumulative: %v\n\n", cumulative)//
		// fmt.Printf("len(parents): %v\n", len(parents))//
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
	if _, err := os.Stat("tables/G1.txt"); os.IsNotExist(err) {
		tableG1(tables)
		file := createFile("tables/G1.txt")
		defer file.Close()
		for ePidx := 0; ePidx < 495; ePidx++ {
			for cOidx := 0; cOidx < 2187; cOidx++ {
				_, err = file.WriteString(fmt.Sprintf("%d", tables.G1[ePidx][cOidx]))
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
			tables.G1[ePidx][cOidx] = depth - 48
			cOidx++
			if cOidx >= 2187 {
				cOidx = 0
				ePidx++
			}
		}
	}
}