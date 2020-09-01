package rubik

import (
	"fmt"
	"math"
	"strconv"
	"os"
)

func cP2index(cube *cepo) int {
	n := 8
	index := 0
	for i := 0; i < n; i++ {
		index = index * (n - i)
		for j := i+1; j < n; j++ {
			if cube.cP[i] > cube.cP[j] {
				index++
			}
		}
	}
	return index
}

// // index2cP8 reverses cP2index // debug tool
// func index2cP8(index int) [8]int {
// 	var cP [8]int
// 	cP[7] = 1
// 	for i := 7; i >= 0; i-- {
// 		cP[i] = 1 + (index % (8-i))
// 		index = (index - (index % (8-i)))/(8-i)
// 		for j := i + 1; j < 8; j++ {
// 			if cP[j] >= cP[i] {
// 				cP[j] = cP[j]+1
// 			}
// 		}
// 	}
// 	return cP
// }

func eP2Binary8(cube *cepo) [8]bool {
	var binary [8]bool
	for i := 0; i < 8; i++ {
		if cube.eP[i] > 3 {
			binary[i] = true
		}
	}
	return binary
}

func binaryBool2Decimal8(binary [8]bool) int {
	var decimal int
	for i := 0; i < 8; i++ {
		if binary[i] == true {
			decimal += int(math.Pow(2, float64(7-i)))
		}
	}
	return decimal
}

func eP2index8(cube *cepo, tables *tables) uint8 {
	ePbinary := eP2Binary8(cube)
	idxEP := binaryBool2Decimal8(ePbinary)
	return tables.G2ePindex[idxEP]
}

func tableG2IdxConv(tables *tables) { // make file/read from file?
	var converted uint8
	var idx int64
	for idx = 0; idx <255; idx++ {
		var count uint8
		binary := strconv.FormatInt(idx, 2)
		for _, bit := range binary {
			if bit == '1' {
				count++
			}
		}
		if count == 4 {
				tables.G2ePindex[idx] = converted
				converted++
		}
	}
}

// cornersInOrbit returns true if corners 0-3 are in position 0-3, and 4-7 in 4-7
// (can be solved using only G3 moves, double face turns) 
func cornersInOrbit(cube *cepo) bool {
	for i := 0; i < 4; i++ {
		if cube.cP[i] > 3 {
			return false
		}
	}
	return true
}

// cPinList returns true if given cube corner permutation is in list of inital cubes
func cPinList(cube *cepo, initial []cepo) bool {
	for _, permuation := range initial {
		if cP2index(cube) == cP2index(&permuation) {
			return true
		}
	}
	return false
}

// initial96cubes returns the 96 cubes with corners in orbit
func initial96cubes() []cepo {
	var initial []cepo
	var parents []cepo
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
				}
				children = append(children, *child)
			}
		}
		parents = children
	}
	return initial
}

func tableG2(tables *tables) {
	fmt.Printf("\nGenerating pruning table for G2")
	parents := initial96cubes()
	var depth uint8
	for depth < 13 {
		depth++
		var children []cepo
		for _, parent := range parents {
			for _, move := range listMoves(&parent, 2) {
				child := newNode(&parent, move)
				spin(move, child)
				idxCP := cP2index(child)
				idxEP := eP2index8(child, tables)
				if tables.G2[idxCP][idxEP] == 0 && !(idxCP == 0 && idxEP == 0) {
					tables.G2[idxCP][idxEP] = depth
					children = append(children, *child)
				}
			}
		}
		parents = children
		fmt.Printf(".")
	}
}

func readHex(char uint8) uint8 {
	if char < 97 {
		return char - 48
	} else {
		return char - 87
	}
}

func makeTableG2(tables *tables) {
	tableG2IdxConv(tables)
	if _, err := os.Stat("tables/G2.txt"); os.IsNotExist(err) {
		tableG2(tables)
		file := createFile("tables/G2.txt")
		defer file.Close()
		for cPidx := 0; cPidx < 40320; cPidx++ {
			for ePidx := 0; ePidx < 70; ePidx++ {
				_, err = file.WriteString(fmt.Sprintf("%x", tables.G2[cPidx][ePidx]))
				if err != nil {
					errorExit("failed to write to file")
				}
			}
		}
	} else {
		file := readFile("tables/G2.txt")
		cPidx := 0
		ePidx := 0
		for _, depth := range file {
			tables.G2[cPidx][ePidx] = readHex(depth)
			ePidx++
			if ePidx >= 70 {
				ePidx = 0
				cPidx++
			}
		}
	}
}