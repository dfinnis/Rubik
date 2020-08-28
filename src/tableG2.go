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
		index = index * (n + 1 - i)
		for j := i+1; j < n; j++ {
			if cube.cP[i] > cube.cP[j] {
				index++
			}
		}
	}
	return index
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
	var converted int16// = 1
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

func makeTableG2(tables *tables) {
	tableG2IdxConv(tables)
	cube := initCube()
	fmt.Printf("\ncPindex: %v\n", cP2index(cube))//
	fmt.Printf("\nePindex: %v\n", eP2index8(cube, tables))//
	spin("L F2 U2", cube)
	fmt.Printf("\ncPindex: %v\n", cP2index(cube))//
	fmt.Printf("\nePindex: %v\n", eP2index8(cube, tables))//
	fmt.Printf("\ntables.G2ePindex: %v\n", tables.G2ePindex)//
}