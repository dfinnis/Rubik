package rubik

import (
	"fmt"
	"math"
)

// func conv12index(eP [12]int8) int {
// 	var decimal int
// 	for i := 0; i < 12; i++ {
// 		decimal += int(eP[i]) * int(math.Pow(2, float64(i)))
// 	}
// 	return decimal
// }

func ePindex4(cube *cepo) int {
	var index int
	for i := 0; i < 3; i++ {
		fmt.Printf("cube.eP[i] : %v\n", cube.eP[i])//
		index += int(cube.eP[i]) * int(math.Pow(2, float64(2-i)))
		fmt.Printf("math.Pow(2, float64(3-i)): %v\n", math.Pow(2, float64(2-i)))//
		fmt.Printf("index: %v\n", index)//
	}
	return index
}

func ePindexConverter()/* [6912]uint8*/ {
	fmt.Printf("\nOH HIII!\n")
	cube := initCube()
	index := ePindex4(cube)
	fmt.Printf("index done: %v\n", index)//
}

func cPindexConverter() [40320]uint8 {
	var initial []cepo
	var parents []cepo
	var cPindexConv [40320]uint8
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
					cPindexConv[cP2index(child)] = converted
					converted++
				}
				children = append(children, *child)
			}
		}
		parents = children
	}
	// fmt.Printf("count: %v\n", count)//
	return cPindexConv
}

func makeTableG3(tables *tables) {
	fmt.Printf("\nGenerating pruning table for G3")
	cPindexConv := cPindexConverter()
	ePindexConverter()
	cube := initCube()
	spin("U2 D2", cube)//
	cPindex := cP2index(cube)
	fmt.Printf("\ncPindexConv[cPindex]: %v\n", cPindexConv[cPindex])//
	// fmt.Printf("cPindexConv: %v\n", cPindexConv)//
}