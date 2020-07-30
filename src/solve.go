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

// 15 moves max
func heuristicG2(cube *[6]uint32) uint8 {
	var color uint8
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
	fmt.Printf("color: %v\n", color)//
	return 42
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
	var path []rubik
	path = append(path, *r)
	for {
		cost, solution := search(path, 0, bound)
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

func listMoves(node *rubik) []string {
	moves := []string{
		"U2",
		"D2",
		"R2",
		"L2",
		"F2",
		"B2",
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

func search(path []rubik, g uint8, bound uint8) (uint8, string) {
	node := path[len(path) - 1]
	// fmt.Printf("Move: %v\n", &path[i].move)//
	// dumpCube(&node.cube)//
	f := g + heuristicG3(&node.cube)
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
	move := listMoves(&node)
	var min uint8 = 255 // ∞
	for i:= 0; i < len(move); i++ {
		new := newNode(&node.cube, move[i])
		spin(move[i], &new.cube)
		// fmt.Printf("Move: %v\n", new.move)//
		// dumpCube(&new.cube)//
		if inPath(new, path) == false {
			path = append(path, *new)
			// dumpPath(path)//
			cost, solution := search(path, g + heuristicG3(&new.cube), bound)
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
	var bound uint8 = heuristicG2(&r.cube)
	fmt.Printf("bound G2: %v\n", bound)//
	solution := idaStar(r)
	// fmt.Printf("solution: %v\n", solution)//
	// solution = randomMix()/////////
	return solution
}