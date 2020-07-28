package rubik

import (
	"fmt"
)

func newNode(newCube *[6]uint32) *rubik {
	return &rubik{
		cube:   *newCube,
	}
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

func solve(r *rubik) string {
	if isSolved(&r.cube) {
		return ""
	}
	idaStar(r)
	solution := randomMix()/////////
	return solution
}

func inPath(node *rubik, path []rubik) bool {
	for i := range path {
		if node.cube == path[i].cube {
			return true
		}
	}
	return false
}

func dumpPath(path []rubik) {
	for i := range path {
		dumpCube(&path[i].cube)
	}
}

func idaStar(r *rubik) string {
	var bound uint8 = heuristicG3(&r.cube)
	fmt.Printf("bound: %v\n", bound)//
	var path []rubik
	path = append(path, *r)
	for {
		cost := search(path, 0, bound)
		// if t = FOUND then return (path, bound)
		if cost == 255 {
			return "Found" /// replace with solution!!
		}
		// if t = ∞ then return NOT_FOUND
		bound = cost
		fmt.Printf("bound 2: %v\n", bound)//
		break//
	}
	return "Error"
	// dumpCube(&path[0].cube)//
}

func search(path []rubik, g uint8, bound uint8) uint8 {
	node := path[len(path) - 1]
	dumpCube(&node.cube)//
	f := g + heuristicG3(&node.cube)
	fmt.Printf("f: %v\n", f)
	if f > bound {
		return f
	}
	if isSolved(&node.cube) {
		return 255 // FOUND
	}
	move := []string{
		"U2",
		"D2",
		"R2",
		"L2",
		"F2",
		"B2",
	}
	var min uint8 = 255 // ∞
	for i:= 0; i < 6; i++ {
		new := newNode(&node.cube)
		spin(move[i], &new.cube)
		// dumpCube(&new.cube)//
		if inPath(new, path) == false {
			path = append(path, *new)
			cost := search(path, g + heuristicG3(&new.cube), bound)
			if cost == 255 {
				return 255
			}
			if cost < min {
				min = cost
			}
			path = path[1:] // pop
		}
		// heuristic := heuristicG3(&new.cube)
	}
	fmt.Printf("len(path): %v\n", len(path))//
	dumpPath(path)//
	return min
}