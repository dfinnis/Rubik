package rubik

import (
	"fmt"
	"math"
	"strconv"
	// "reflect"//
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
	cPindex := cP2index(cube)
	fmt.Printf("\ncPindex: %v\n", cPindex)//
	// fmt.Println(reflect.TypeOf(cPindex))//
	index2cP := index2cP8(cPindex)
	fmt.Printf("index2cP: %v\n", index2cP)//

	fmt.Printf("ePindex: %v\n\n", eP2index8(cube, tables))//
	spin("L F2 U2 D2 R", cube)
	dumpCepo(cube)//

	cPindex2 := cP2index(cube)
	fmt.Printf("cPindex: %v\n", cPindex2)//
	index2cP2 := index2cP8(cPindex2)
	// fmt.Println(reflect.TypeOf(cPindex))//
	fmt.Printf("index2cP: %v\n", index2cP2)//
	fmt.Printf("ePindex: %v\n", eP2index8(cube, tables))//
	// fmt.Printf("\ntables.G2ePindex: %v\n", tables.G2ePindex)//
}