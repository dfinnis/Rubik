package rubik

import (
	"os"
	"fmt"
	"math"
)

func listAllMoves() []string {
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
	// // for i := 0; i < n; i++ {
	// // 	move := dry[rand.Intn(len(dry))]
	// // 	mix += move
	// 	if move == "U" || move == "U'" || move == "U2" {
	// 		if stringInSlice("D", dry) {
	// 			dry = spin[3:]
	// 		} else {
	// 			dry = spin[6:]
	// 		}
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
	// 		}
	return moves
}

func newNodeCepo(parent *cepo, move string) *cepo {
	return &cepo{
		cP:   	parent.cP,
		cO:   	parent.cO,
		eP:   	parent.eP,
		eO:   	parent.eO,
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

func tableGenerator() {
	if _, err := os.Stat("tables/G0.txt"); os.IsNotExist(err) {
		var table [4096]uint8
		parent := initCepo()
		
		for i, move := range listAllMoves() {
			fmt.Printf("\nmove %v: %v\n", i, move)//
			child := newNodeCepo(parent, "")
			spinCepo(move, child)
			// dumpCepo(child)//
			index := binaryToDecimal(child.eO)
			if index != 0 && table[index] == 0 {
				table[index] = 1
			}
			fmt.Printf("index: %v\n", index)//
		}
		// fmt.Printf("listAllMoves: %v\n", listAllMoves())//
		fmt.Printf("table: %v\n", table)//

		f, err := os.Create("tables/G0.txt")
		if err != nil {
			fmt.Printf("error creating file: %v", err)
			return
		}
		defer f.Close()
		for i := 0; i < len(table); i++ {
			_, err = f.WriteString(fmt.Sprintf("%d", table[i]))
			if err != nil {
				fmt.Printf("error writing to file: %v", err)
			}
		}
	}
}