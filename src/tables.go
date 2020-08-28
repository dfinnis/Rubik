package rubik

import (
	"os"
	"fmt"
	"math"
	"io/ioutil"
	"strconv"
)

type tables struct {
	G0 [2048]uint8
	colIndex [4096]int16
	G1cO [2187]uint8
	G1eP [495]uint8
}

func move2nul(move string, move2 string) string {
	if move != "" && move2 != "" {
		if move[0] == 'U' && move2[0] != 'D' || 
			move[0] == 'D' && move2[0] != 'U' ||
			move[0] == 'F' && move2[0] != 'B' || 
			move[0] == 'B' && move2[0] != 'F' ||
			move[0] == 'L' && move2[0] != 'R' || 
			move[0] == 'R' && move2[0] != 'L' {
			move2 = ""
		}
	}
	return move2
}

func newNode(parent *cepo, move string) *cepo {
	return &cepo{
		cP:   	parent.cP,
		cO:   	parent.cO,
		eP:   	parent.eP,
		eO:   	parent.eO,
		move2:	move2nul(move, parent.move),
		move:	move,
	}
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
	// fmt.Printf("edgePermutationBin: %v\n", edgePermutationBin)
	idxEP := binaryBool2Decimal(ePbinary)
	// fmt.Printf("index: %v\n", index)//
	return tables.colIndex[idxEP]
}


func binaryToDecimal(binary [12]int8) int {
	var decimal int
	for i := 0; i < 11; i++ {
		decimal += int(binary[i]) * int(math.Pow(2, float64(10-i)))
	}
	return decimal
}

func cO2index(cube *cepo) int {
	var index int
	for i := 0; i < 7; i++ {
		index = index * 3
		index = index + int(cube.cO[i])
	}
	return index
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


// func tableFull(table [2048]uint8) bool {
// 	for i := 1; i < len(table); i++ {
// 		if table[i] == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func tableG0() [2048]uint8 {
	fmt.Printf("\nGenerating pruning table for G0")
	var table [2048]uint8
	var depth uint8
	var parents []cepo
	parents = append(parents, *initCube())
	
	for depth < 6 {
		// for tableFull(table) == false {
		var children []cepo
		// var count int//
		depth++
		for _, parent := range parents {
			for _, move := range listMoves(&parent, 0) {
				// fmt.Printf("\nmove %v: %v\n", i, move)//
				child := newNode(&parent, move)
				spin(move, child)
				// dumpCepo(child)//
				index := binaryToDecimal(child.eO)
				if index != 0 && table[index] == 0 {
					table[index] = depth
					// count++//
				}
				children = append(children, *child)
			}
		}
		parents = children
		fmt.Printf(".")
		// fmt.Printf("depth: %v\n", depth)//
		// fmt.Printf("count: %v\n", count)//
		// fmt.Printf("len(parents): %v\n", len(parents))//
	}
	for i, depth := range table {
		if i > 0 && depth == 0 {
			table[i] = 7
		}
	}
	// fmt.Printf("table: %v\n", table)//
	fmt.Printf("\n")
	return table
}

func createFile(filepath string) *os.File {
	file, err := os.Create(filepath)
	if err != nil {
		errorExit("failed to create file")
	}
	return file
}

func readFile(filepath string) []byte {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		errorExit("failed to read pruning table file")
	}
	return file
}

func makeTableG0(tables *tables) {
	if _, err := os.Stat("tables/G0.txt"); os.IsNotExist(err) {
		tables.G0 = tableG0()
		file := createFile("tables/G0.txt")
		defer file.Close()
		for i := 0; i < len(tables.G0); i++ {
			_, err = file.WriteString(fmt.Sprintf("%d", tables.G0[i]))
			if err != nil {
				errorExit("failed to write to file")
			}
		}
	} else {
		file := readFile("tables/G0.txt")
		for i, depth := range file {
			tables.G0[i] = depth - 48
		}
	}
}

func tableG1IdxConv(tables *tables) { // make file/read from file?
	var converted int16// = 1
	var idx int64
	// for idx = 7; idx <=1920; idx++ {
	for idx = 0; idx <4096; idx++ {
		var count uint8
		// fmt.Printf("idx: %v\n", idx) // 
		binary := strconv.FormatInt(idx, 2)
		for _, bit := range binary {
			if bit == '1' {
				count++
			}
		}
		if count == 4 {
				// fmt.Printf("binary: %v\n", binary)///
				// fmt.Printf("idx: %v\n", idx) // 
				// fmt.Printf("converted: %v\n", converted)///
				tables.colIndex[idx] = converted
				converted++
		}
	}
	// for i, index := range colIndex {
	// 	if index != 0 {
	// 		fmt.Printf("index %v: %v\n", i, index) // 
	// 	}
	// }
}

func tableG1(tables *tables) {
	// var tableG1 [495][2187]uint8
	// var tableG1cO [2187]uint8
	// var tableG1eP [495]uint8
	fmt.Printf("\nGenerating pruning table for G1")
	var depth uint8
	var parents []cepo
	parents = append(parents, *initCube())
	for depth < 6 {//10//////////////////////////////////////////////////////
		// for tableFull(table) == false {
		var children []cepo
		var count int//
		depth++
		// fmt.Printf("len(parents): %v\n", len(parents))//
		for _, parent := range parents {
			for _, move := range listMoves(&parent, 1) {
				// fmt.Printf("\nmove %v: %v\n", i, move)//
				child := newNode(&parent, move)
				spin(move, child)
				// dumpCepo(child)//
				idxCO := cO2index(child)
				idxEP := eP2index(child, tables)

				if idxCO != 0 && tables.G1cO[idxCO] == 0 {
					tables.G1cO[idxCO] = depth
					count++//
				}
				if idxEP != 0 && tables.G1eP[idxEP] == 0 {
					tables.G1eP[idxEP] = depth
				}
				// if tableG1[idxEPconverted][idxCO] == 0 && !(idxEPconverted == 0 && idxCO == 0) {
				// if tableG1[idxEPconverted][idxCO] == 0 && (idxEPconverted != 0 || idxCO != 0) {
				// 	tableG1[idxEPconverted][idxCO] = depth
				// 	fmt.Printf("idxEPconverted: %v\n", idxEPconverted)//
				// 	fmt.Printf("idxCO: %v\n", idxCO)//
				// 	count++//
				// 	// fmt.Printf("count++\n\n")//
				// 	// children = append(children, *child)
				// }
				children = append(children, *child)
			}
		}
		parents = children
		fmt.Printf(".")
		fmt.Printf("depth: %v\n", depth)//
		fmt.Printf("count: %v\n\n", count)//
		// fmt.Printf("len(parents): %v\n", len(parents))//
	}
	// for i, depth := range table {
	// 	if i > 0 && depth == 0 {
	// 		table[i] = 10
	// 	}
	// }
	for i, entry := range tables.G1cO {
		if i > 0 {
			if entry == 0 {
				tables.G1cO[i] = 7
				// fmt.Printf("OH HIIII\n")//
			}
		}
	}
	fmt.Printf("\ntableG1cO: %v\n\n", tables.G1cO)//
	fmt.Printf("\ntableG1eP: %v\n", tables.G1eP)//
	fmt.Printf("\n###########################################################################\n")
}

func makeTableG1(tables *tables) {
	tableG1IdxConv(tables)
	tableG1(tables)
}

// func makeTables() [2048]uint8 {
func makeTables() *tables {
	tables := &tables{}
	makeTableG0(tables)
	makeTableG1(tables)


	return tables
}