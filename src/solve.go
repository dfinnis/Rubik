package rubik

import (
	"fmt"
)

func newNode(newCube *[6]uint32, move string) *rubik {
	return &rubik{
		cube:   *newCube,
		move:	move,
	}
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
	return 42
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
		// fmt.Printf("face: %v\n", face)//
		// fmt.Printf("cube[face]&0x70000000: %x\n", cube[face]&0x70000000)//
		// fmt.Printf("cube[face]&0x70000000 / uint32(math.Pow(16, 6): %x\n", cube[face]&0x70000000 / uint32(math.Pow(16, 6)))//
		// fmt.Printf("cube[face]&0x70000000  >> 6: %x\n", cube[face]&0x70000000 >> 24)//

		// fmt.Printf("cube[face]&0x70: %x\n", cube[face]&0x70)//
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

func inPath(node *rubik, path []rubik) bool {
	for i := range path {
		if node.cube == path[i].cube {
			return true
		}
	}
	return false
}

// func dumpPath(path []rubik) {
// 	for i := range path {
// 		fmt.Printf("------------------------------------\n")
// 		fmt.Printf("Move: %v\n", path[i].move)
// 		dumpCube(&path[i].cube)
// 	}
// }

func idaStar(r *rubik) string {
	// var solution *string
	var bound uint8 = heuristicG3(&r.cube)
	// fmt.Printf("bound: %v\n", bound)//
	var subgroup uint8 = 2 // 0!!! test heuristics to establish subgroup
	if heuristicG2(&r.cube) == 0 {
		subgroup = 3
	}
	// fmt.Printf("subgroup: %v\n", subgroup)//

	var path []rubik
	path = append(path, *r)
	for {
		cost, solution := search(path, 0, bound, subgroup)
		// if t = FOUND then return (path, bound)
		if cost == 255 {
			// fmt.Printf("***************	END	********************\n")//
			// dumpPath(path)//
			// return "Found" /// replace with solution!!
			return solution
		}
		// if t = ∞ then return NOT_FOUND
		bound = cost
		// fmt.Printf("bound 2: %v\n", bound)//
	}
	// return "Error"//
	// dumpCube(&path[0].cube)//
}

func listMoves(node *rubik, subgroup uint8) []string {
	// fmt.Printf("subgroup: %v\n", subgroup)//
	moves := []string{
		"U",
		"D",
		"R",
		"L",
		"F",
		"B",
	}
	if subgroup == 1 {
		moves = []string{
			"U2",
			"D2",
			"R",
			"L",
			"F",
			"B",
		}
		// fmt.Printf("Group 1!!!!!!!!!!!!\n")////
	} else if subgroup == 2 {
		moves = []string{
			"U2",
			"D2",
			"R",
			"L",
			"F2",
			"B2",
		}
		// fmt.Printf("Group 2!!!!!!!!!!!!\n")////
	} else if subgroup == 3 {
		moves = []string{
			"U2",
			"D2",
			"R2",
			"L2",
			"F2",
			"B2",
		}
		// fmt.Printf("Group 3!!!!!!!!!!!!\n")////
	}
	// fmt.Printf("move: %v\n", node.move)//
	if node.move != "" {
		for i, move := range moves {
			if move == node.move {
				moves = append(moves[:i], moves[i+1:]...)
				break
			}
		}
	}
	// fmt.Printf("moves: %v\n", moves)//
	return moves
}

func search(path []rubik, g uint8, bound uint8, subgroup uint8) (uint8, string) {
	node := path[len(path) - 1]
	// fmt.Printf("Move: %v\n", &path[i].move)//
	// dumpCube(&node.cube)//
	// heuristic := heuristic(&node.cube, subgroup)
	if heuristic(&node.cube, subgroup) == 0 && subgroup < 3 {
		subgroup++
	}
	f := g + heuristic(&node.cube, subgroup)
	// fmt.Printf("f: %v\n", f)
	if f > bound {
		return f, ""
	}
	if isSolved(&node.cube) {
		var solved string
		for i := 1; i < len(path); i++ {
			solved += path[i].move + " "
		}
		return 255, solved // FOUND
	}
	move := listMoves(&node, subgroup)
	var min uint8 = 255 // ∞
	for i:= 0; i < len(move); i++ {
		new := newNode(&node.cube, move[i])
		spin(move[i], &new.cube)
		// fmt.Printf("Move: %v\n", new.move)//
		// dumpCube(&new.cube)//
		if inPath(new, path) == false {
			path = append(path, *new)
			// dumpPath(path)//
			cost, solution := search(path, g + heuristic(&new.cube, subgroup), bound, subgroup)
			if cost == 255 {
				return 255, solution
			}
			if cost < min {
				min = cost
			}
			path = path[:len(path) - 1] // pop
		}
		// fmt.Printf("##############################\n")//
	}
	// fmt.Printf("len(path): %v\n", len(path))//
	// dumpPath(path)//
	return min, ""
}

func solve(r *rubik) string {
	if isSolved(&r.cube) {
		return ""
	}
	dumpCube(&r.cube)//
	var bound uint8 = heuristicG2(&r.cube)
	fmt.Printf("bound G2: %v\n", bound)//
	solution := idaStar(r)
	// fmt.Printf("solution: %v\n", solution)//
	// solution = randomMix()/////////
	// solution = ""//
	return solution
}