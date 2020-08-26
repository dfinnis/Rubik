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
}


func listAllMoves(cube *cepo) []string {
	moves := []string{
			"U",
			"U'",
			"U2",
			"D",
			"D'",
			"D2",
			"R",
			"R'",
			"R2",
			"L",
			"L'",
			"L2",
			"F",
			"F'",
			"F2",
			"B",
			"B'",
			"B2",
	}
	// dry := moves
	// // // for i := 0; i < n; i++ {
	// // // 	move := dry[rand.Intn(len(dry))]
	// // // 	mix += move
	// if cube.move == "U" || cube.move == "U'" || cube.move == "U2" {
	// 	if cube.move2 == "D" || cube.move2 == "D'" || cube.move2 == "D2" {
	// 		dry = moves[3:]
	// 	} else {
	// 		dry = moves[6:]
	// 	}
// 	} else if move == "D" || move == "D'" || move == "D2" {
// 		if stringInSlice("U", dry) {
// 			dry = append([]string{}, spin[:3]...)
// 			dry = append(dry, spin[6:]...)
// 		} else {
// 			dry = spin[6:]
// 		}
// 	} else if move == "R" || move == "R'" || move == "R2" {
// 		dry = append([]string{}, spin[:6]...)
// 		if stringInSlice("L", dry) {
// 			dry = append(dry, spin[9:]...)
// 		} else {
// 			dry = append(dry, spin[12:]...)
// 		}
// 	} else if move == "L" || move == "L'" || move == "L2" {
// 		if stringInSlice("R", dry) {
// 			dry = append([]string{}, spin[:9]...)
// 		} else {
// 			dry = append([]string{}, spin[:6]...)
// 		}
// 		dry = append(dry, spin[12:]...)
// 	} else if move == "F" || move == "F'" || move == "F2" {
// 		dry = append([]string{}, spin[:12]...)
// 		if stringInSlice("B", dry) {
// 			dry = append(dry, spin[15:]...)
// 		}
// 	} else if move == "B" || move == "B'" || move == "B2" {
// 		if stringInSlice("F", dry) {
// 			dry = spin[:15]
// 		} else {
// 			dry = spin[:12]
	// }
	return moves
}

func newNodeCepo(parent *cepo, move string) *cepo {
	return &cepo{
		cP:   	parent.cP,
		cO:   	parent.cO,
		eP:   	parent.eP,
		eO:   	parent.eO,
		move2:	parent.move,
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

func binaryToDecimal(binary [12]int8) int {
	var decimal int
	for i := 0; i < 11; i++ {
		decimal += int(binary[i]) * int(math.Pow(2, float64(10-i)))
	}
	return decimal
}

func orientation2index(cube *cepo) int {
	var index int
	for i := 0; i < 7; i++ {
		index = index * 3
		index = index + int(cube.cO[i])
	}
	return index
}

func index2orientation(index int) [8]int {
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
	return or
}


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
	parents = append(parents, *initCepo())
	
	for depth < 6 {
		// for tableFull(table) == false {
		var children []cepo
		// var count int//
		depth++
		for _, parent := range parents {
			for _, move := range listAllMoves(&parent) {
				// fmt.Printf("\nmove %v: %v\n", i, move)//
				child := newNodeCepo(&parent, move)
				spinCepo(move, child)
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

func tableGeneratorG0(tables *tables) {
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
	var tableG1 [495][2187]uint8
	fmt.Printf("\nGenerating pruning table for G1")
	var depth uint8
	var parents []cepo
	parents = append(parents, *initCepo())
	for depth < 1 {//10
		// for tableFull(table) == false {
		var children []cepo
		var count int//
		depth++
		fmt.Printf("len(parents): %v\n", len(parents))//
		for _, parent := range parents {
			for _, move := range listAllMoves(&parent) {
				// fmt.Printf("\nmove %v: %v\n", i, move)//
				child := newNodeCepo(&parent, move)
				spinCepo(move, child)
				// dumpCepo(child)//
				// index := binaryToDecimal(child.eO)
				idxCO := orientation2index(child)
				// fmt.Printf("orientation2index: %v\n", index)
				// fmt.Printf("orientation2index2: %v\n", orientation2index(cube))
				// fmt.Printf("index2orientation: %v\n", index2orientation(index))
				// fmt.Printf("orientation2index max: %v\n", orientation2index(cube))
				ePBinary := eP2Binary(child)
				// fmt.Printf("edgePermutationBin: %v\n", edgePermutationBin)
				idxEP := binaryBool2Decimal(ePBinary)
				// fmt.Printf("index: %v\n", index)//
				idxEPconverted := tables.colIndex[idxEP]
				fmt.Printf("idxEPconverted: %v\n", idxEPconverted)//
				fmt.Printf("idxCO: %v\n", idxCO)//

				if !(idxEPconverted == 0 && idxCO == 0) && tableG1[idxEPconverted][idxCO] == 0 {
					tableG1[idxEPconverted][idxCO] = depth
					count++//
					fmt.Printf("count++\n\n")//
				}
				children = append(children, *child)
			}
		}
		parents = children
		fmt.Printf(".")
		fmt.Printf("depth: %v\n", depth)//
		fmt.Printf("count: %v\n", count)//
		// fmt.Printf("len(parents): %v\n", len(parents))//
	}
	// for i, depth := range table {
	// 	if i > 0 && depth == 0 {
	// 		table[i] = 10
	// 	}
	// }
	// fmt.Printf("\ntableG1: %v\n", tableG1)//
	fmt.Printf("\n")
}

func tableGeneratorG1(tables *tables) {
	tableG1IdxConv(tables)
	tableG1(tables)
}

// func tableGenerator() [2048]uint8 {
func tableGenerator() *tables {
	tables := &tables{}
	tableGeneratorG0(tables)
	tableGeneratorG1(tables)


	return tables
}