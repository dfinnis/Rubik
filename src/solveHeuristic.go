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
	} else if subgroup == 2 {
		moves = []string{
			"U2",
			"D2",
			"R",
			"L",
			"F2",
			"B2",
		}
	} else { // subgroup = 3
		moves = []string{
			"U2",
			"D2",
			"R2",
			"L2",
			"F2",
			"B2",
		}
	}
	// if subgroup != 0 {
	// if node.move != "" {
	// 	for i, move := range moves {
	// 		if move == node.move {
	// 			moves = append(moves[:i], moves[i+1:]...)
	// 			break
	// 		}
	// 	}// remove opposite face move, not just last move?? i.e. avoid G0 R L R L????!!!
	// }
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

// Edge Orientation - 7 moves max
func heuristicG0(cube *[6]uint32) uint8 {
	var edgeOriented uint8
	var cubie uint32
	for _, face := range [2]uint8{1, 3} { // if face L or R, facelet not U or D
		var mask uint32 = 0x1000000
		for cubie = 0x7000000; cubie > 0; cubie /= 256 { // iterate 4 edges
			if cube[face]&cubie != 0 && cube[face]&cubie != mask * 5 {
				edgeOriented++
			}
			mask /= 256
		}
	}
	// fmt.Printf("edgeOriented L R: %v\n", edgeOriented)//
	for _, face := range [2]uint8{0, 5} { // if face U or D, facelet not L or R
		var mask uint32 = 0x1000000
		for cubie = 0x7000000; cubie > 0; cubie /= 256 { // iterate 4 edges
			if cube[face]&cubie != mask * 1 && cube[face]&cubie != mask * 3 {
				edgeOriented++
			}
			mask /= 256
		}
	}
	// fmt.Printf("edgeOriented U D: %v\n", edgeOriented)//
	for _, face := range [2]uint8{2, 4} { // if F or B
		var mask uint32 = 0x1000000
		for cubie = 0x7000000; cubie > 0; cubie = cubie >> (4 * 4) { // for top and bottom edge, not L or R
			if cube[face]&cubie != mask * 1 && cube[face]&cubie != mask * 3 {
				edgeOriented++
			}
			mask = mask >> (4 * 4)
		}
		// fmt.Printf("edgeOriented F B (top bottom): %v\n", edgeOriented)//
		mask = 0x10000
		for cubie = 0x70000; cubie > 0; cubie = cubie >> (4 * 4) {  // for right and left edge, not U or D
			if cube[face]&cubie != 0 && cube[face]&cubie != mask * 5 {
				edgeOriented++
			}
			mask = mask >> (4 * 4)
		}
		// fmt.Printf("edgeOriented F B (right left): %v\n", edgeOriented)//
	}
	// fmt.Printf("edgeOriented: %v\n", edgeOriented)//
	// return (24 - edgeOriented) / 2
	// fmt.Printf("heuristicG0: %v\n", 24 - edgeOriented)///
	return 24 - edgeOriented// + heuristicG1Edges(cube)
}

//  Corner Twist, & middle edges in middle layer - 10 moves max
func heuristicG1(cube *[6]uint32) uint8 {
	var color uint8
	var cubie uint32
	for _, face := range [2]uint8{0, 5} { // face U and D, facelets U or D
		var mask uint32 = 0x10000000
		for cubie = 0x70000000; cubie > 0; cubie /= 16 {
			if cube[face]&cubie == 0 || cube[face]&cubie == mask * 5 {
				color++
			}
			mask /= 16
		}
	}
	// fmt.Printf("heuristicG1: %v\n", 16 - color)///
	// return 16 - color
	// return (16 - color) / 2
	// fmt.Printf("heuristicG1: %v\n", 8 - color / 2)///
	return 8 - color / 2
}

// corners into their tetrads, edges into their slices, corner permutations into squares group - 13 moves max
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
	for _, face := range [4]uint8{1, 2, 3, 4} {
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

// 15 moves max
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
	// if subgroup == 0 {
	// 	return heuristicG0(cube)
	// } else if subgroup == 1 {
	// 	return heuristicG1(cube)
	// } else if subgroup == 2 {
	// 	return heuristicG2(cube)
	// } else { // subgroup = 3
	// 	return heuristicG3(cube)
	// }	

	// heuristic := heuristicG0(cube)
	// if subgroup >= 1 {
	// 	heuristic += heuristicG1(cube)
	// }
	// if subgroup >= 2 {
	// 	heuristic += heuristicG2(cube)
	// }
	// if subgroup == 3 {
	// 	heuristic += heuristicG3(cube)
	// }
	// return heuristic

	if subgroup == 0 {
		return heuristicG0(cube)// + heuristicG1Edges(cube)
	} else if subgroup == 1 {
		// return heuristicG1Edges(cube)
		return heuristicG1(cube)// + heuristicG0(cube)
		// return heuristicG1(cube) + heuristicG0(cube)
		// return heuristicG1Corners(cube)// + heuristicG0(cube)
		// return heuristicG1Corners(cube)

	} else if subgroup == 2 {
		return heuristicG2(cube)
		// return heuristicG2(cube) + heuristicG0(cube)
	} else { // subgroup = 3
		return heuristicG3(cube)
	}	
}

