package rubik

import (
	// "fmt"
)

func listMoves(node *rubik, subgroup uint8) []string {
	moves := []string{}
	if subgroup == 0 {
		moves = []string{
			"U",
			"D",
			"R",
			"L",
			"F",
			"B",
		}
	} else if subgroup == 1 {
		moves = []string{
			"U2",
			"D2",
			"R",
			"L",
			"F",
			"B",
		}
		// fmt.Printf("Oh HIIII!!!!\n")//
	} else if subgroup == 2 {
		moves = []string{
			"U2",
			"D2",
			"R",
			"L",
			"F2",
			"B2",
		}
	} else { //if subgroup == 3 {
		moves = []string{
			"U2",
			"D2",
			"R2",
			"L2",
			"F2",
			"B2",
		}
	}
	if node.move != "" {
		for i, move := range moves {
			if move == node.move {
				moves = append(moves[:i], moves[i+1:]...)
				break
			}
		}// remove opposite face move, not just last move?? i.e. avoid G0 R L R L????!!!
	}
	// fmt.Printf("moves: %v\n", moves)//
	return moves
}

func oppositeFace(face uint32) uint32 {
	if face == 0 {
		return 5
	} else if face == 1 {
		return 3
	} else if face == 2 {
		return 4
	} else if face == 3 {
		return 1
	} else if face == 4 {
		return 2
	} else {
		return 0
	}
}

func heuristicG0(cube *[6]uint32) uint8 {
	return 42
}

func heuristicG1(cube *[6]uint32) uint8 {
	// fmt.Printf("OH HIII!!!\n")////
	var color uint8
	var cubie uint32
	for _, face := range [2]uint8{1, 3} {
		// fmt.Printf("face: %v\n", face)//
		var mask uint32 = 0x10000000
		for cubie = 0x70000000; cubie > 0; cubie /= 16 {
			// fmt.Printf("count\n")//
			// fmt.Printf("mask: %v\n", mask)//
			if cube[face]&cubie == mask || cube[face]&cubie == mask * 3 {
				color++
				// fmt.Printf("color++\n")//
			}
			mask /= 16
		}
	}
	// fmt.Printf("color: %v\n\n", color)//
	return (16 - color) / 2
}

// 15 moves max
func heuristicG2(cube *[6]uint32) uint8 {
	var color uint8
	var parity uint8
	var face uint32
	for face = 0; face < 6; face++ {
		var cubie uint32
		var mask uint32 = 0x10000000
		for cubie = 0x70000000; cubie > 0; cubie /= 16 {
			if cube[face]&cubie == mask * face || cube[face]&cubie == mask * oppositeFace(face) {
				color++
			}
			mask /= 16
		}
	}
	// fmt.Printf("color: %v\n", color)//
	for _, face := range [4]uint8{0, 2, 4, 5} {
		if cube[face]&0x70000000 >> (6 * 4) == cube[face]&0x70 {
			parity++
		}
		if cube[face]&0x700000 >> (2 * 4) == cube[face]&0x7000 {
			parity++
		}
	}
	// fmt.Printf("parity: %v\n", parity)//
	if parity == 0 {
		parity = 8
	}
	if parity != 8 {
		parity = 0
	}
	return (56 - (color + parity)) / 4
}

// 17 moves max
func heuristicG3(cube *[6]uint32) uint8 {
	var correct uint8
	var face uint32
	for face = 0; face < 6; face++ {
		var cubie uint32
		var mask uint32 = 0x10000000
		for cubie = 0x70000000; cubie > 0; cubie /= 16 {
			if cube[face]&cubie == mask * face {
				correct++
			}
			mask /= 16
		}
	}
	return (48 - correct) / 4
}

func heuristic(cube *[6]uint32, subgroup uint8) uint8 {
	if subgroup == 0 {
		return heuristicG0(cube)
	} else if subgroup == 1 {
		return heuristicG1(cube)
	} else if subgroup == 2 {
		return heuristicG2(cube)
	} else { // if subgroup == 3 {
		return heuristicG3(cube)
	}	
}