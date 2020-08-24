package rubik

import (
	"os"
	"fmt"
	"math"
	"io/ioutil"
)

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

func binaryToDecimal(binary [12]int8) int {
	var decimal int
	for i, bit := range binary {
		decimal += int(bit) * int(math.Pow(2, float64(11-i)))
	}
	return decimal
}

// func tableFull(table [4096]uint8) bool {
// 	for i := 1; i < len(table); i++ {
// 		if table[i] == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func tableGeneratorG0() [4096]uint8 {
	fmt.Printf("\nGenerating pruning table for G0")
	var table [4096]uint8
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

func tableGenerator() [4096]uint8 {
	var tableG0 [4096]uint8//
	if _, err := os.Stat("tables/G0.txt"); os.IsNotExist(err) {
		tableG0 = tableGeneratorG0()
		file := createFile("tables/G0.txt")
		defer file.Close()
		for i := 0; i < len(tableG0); i++ {
			_, err = file.WriteString(fmt.Sprintf("%d", tableG0[i]))
			if err != nil {
				errorExit("failed to write to file")
			}
		}
	} else {
		file := readFile("tables/G0.txt")
		for i, depth := range file {
			tableG0[i] = depth - 48
		}
	}

	// var tableG1 [4096]uint8//
	return tableG0
}