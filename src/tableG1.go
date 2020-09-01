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

// index2cO reverses cO2index, accepting 0-2186 returning [00000000]-[22222221]
func index2cO(index int) [8]int8 {
	var s int
	var or [8]int
	for i := 6; i >= 0; i-- {
		or[i] = index % 3
		s = s - or[i]
		if s < 0 {
			s = s + 3
		}
		index = (index - or[i]) / 3
	}
	or[7] = s
	var or8 [8]int8
	for i := range or8 {
		or8[i] = int8(or[i])
	}
	return or8
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
	// fmt.Printf("idxEP pre-convert: %v\n", idxEP)//
	// fmt.Printf("tables.G1ePindex[idxEP]: %v\n", tables.G1ePindex[idxEP])//
	return tables.G1ePindex[idxEP]
}

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
	// fmt.Printf("converted: %v\n", converted)//
}

// func initial24cubes(tables *tables) []cepo {
// 	var initial []cepo
// 	var parents []cepo
// 	initial = append(initial, *initCube())
// 	parents = append(parents, *initCube())
// 	for depth := 0; depth < 3; depth++ {
// 		var children []cepo
// 		for _, parent := range parents {
// 			for _, move := range listMoves(&parent, 3) {
// 				child := newNode(&parent, move)
// 				spin(move, child)
// 				if /*isSubgroup(child) == 2 && */eP2index(child, tables) == 0 && cO2index(child) == 0 && inPath(child, initial) == false {
// 					initial = append(initial, *child)
// 					// dumpCube(child)//
// 				}
// 				children = append(children, *child)
// 			}
// 		}
// 		parents = children
// 	}
// 	fmt.Printf("len(initial): %v\n", len(initial))
// 	return initial
// }

func tableG1(tables *tables) {
	fmt.Printf("\nGenerating pruning table for G1")
	var parents []cepo
	parents = append(parents, *initCube())
	// parents = initial24cubes(tables)
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

				if tables.G1[idxEP][idxCO] == 0 && !(idxEP == 0 && idxCO == 0) {
					tables.G1[idxEP][idxCO] = depth
					// children = append(children, *child)//!!!!!!!
					count++//
					cumulative++//
				// } else if tables.G1[idxEP][idxCO] != 0 {
				// // } else if (idxEP == 0 && idxCO == 0) && isSolved(child) == false {//
				// 	children = append(children, *child)
					// fmt.Printf("child.move: %v\n", child.move)//
					// fmt.Printf("child.move2: %v\n", child.move2)//
					// fmt.Printf("OH MYYYY!!!\n")//
					// dumpCube(child)
				} else if !(idxEP == 0 && idxCO == 0) {//
					// fmt.Printf("idxCO: %v\n", idxCO)//
					// fmt.Printf("idxEP: %v\n", idxEP)//
					// fmt.Printf("tables.G1[idxEP][idxCO]: %v\n", tables.G1[idxEP][idxCO])//
				}
				children = append(children, *child)//
			}
		}
		parents = children
		fmt.Printf(".")
		fmt.Printf("depth: %v\n", depth)//
		fmt.Printf("count: %v\n", count)//
		fmt.Printf("cumulative: %v\n\n", cumulative)//
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

	// for i := 0; i <= 2186; i++{
	// 	cO := index2cO(i)
	// 	fmt.Printf("cO %v: %v\n", i, cO)//
	// 	index := cO2index(cO)
	// 	if index != i {
	// 		fmt.Printf("WTF!!!!!!!!!!!!!!!!!!!!!\n")//
	// 	}
	// 	fmt.Printf("index: %v\n", index)//
	// }

	// cube := initCube()//
	// // dumpCube(cube)//
	// spin("R L2 U2 F B U2 D2 F B", cube)
	// dumpCube(cube)//
	// fmt.Println()//
	// // binary := eP2Binary(cube)
	// var binary [12]bool
	// binary[0] = true
	// binary[1] = true
	// binary[2] = true
	// binary[6] = true
	// fmt.Printf("binary: %v\n", binary)//
	// idxEPpre := binaryBool2Decimal(binary)
	// fmt.Printf("idxEP pre-convert: %v\n", idxEPpre)//
	// fmt.Printf("tables.G1ePindex[idxEP]: %v\n", tables.G1ePindex[idxEPpre])//

	// idxEP := eP2index(cube, tables)//
	// fmt.Printf("idxEP: %v\n", idxEP)//
	// index2eP(idxEP, tables)
	// // fmt.Printf("index2eP(idxEP): %v\n", index2eP(idxEP, tables))//

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