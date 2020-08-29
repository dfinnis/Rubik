package rubik

import (
	"fmt"
	"math"
	"strconv"
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

func index2cP8(index int) [8]int {
	var cP [8]int
	cP[7] = 1
	for i := 7; i >= 0; i-- {
		cP[i] = 1 + (index % (8-i))
		index = (index - (index % (8-i)))/(8-i)
		for j := i + 1; j < 8; j++ {
			if cP[j] >= cP[i] {
				cP[j] = cP[j]+1
			}
		}
	}
	return cP
}

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

func eP2index8(cube *cepo, tables *tables) int16 {
	ePbinary := eP2Binary8(cube)
	idxEP := binaryBool2Decimal8(ePbinary)
	return tables.G2ePindex[idxEP]
}

func tableG2IdxConv(tables *tables) { // make file/read from file?
	var converted int16
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

func cornersInOrbit(cube *cepo) bool {
	for i := 0; i < 4; i++ {
		if cube.cP[i] > 3 {
			return false
		}
	}
	return true
}

func cPinList(cube *cepo, initial []cepo) bool {
	for _, permuation := range initial {
		if cP2index(cube) == cP2index(&permuation) {
			return true
		}
	}
	return false
}

func tableG2(tables *tables) {
	fmt.Printf("\nGenerating pruning table for G2")
	var initial []cepo
	var parents []cepo
	parents = append(parents, *initCube())
	// var count int//
	var depth uint8
	for depth < 4 {
		depth++
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
	parents = initial
	// fmt.Printf("\ncount: %v\n", count)//
	fmt.Printf("\nlen(initial): %v\n", len(initial))//
	fmt.Printf("len(parents): %v\n", len(parents))//


	fmt.Printf("\n\n###########################################################################\n")//
}

func makeTableG2(tables *tables) {
	tableG2IdxConv(tables)
	tableG2(tables)
	// cube := initCube()

	// cPindex := cP2index(cube)
	// ePindex := eP2index8(cube, tables)



	// fmt.Printf("\ncPindex: %v\n", cPindex)//
	// fmt.Println(reflect.TypeOf(cPindex))//
	// index2cP := index2cP8(cPindex)
	// fmt.Printf("index2cP: %v\n", index2cP)//

	// fmt.Printf("ePindex: %v\n\n", eP2index8(cube, tables))//
	// spin("L F2 U2 D2 R", cube)//
	// dumpCube(cube)//

}