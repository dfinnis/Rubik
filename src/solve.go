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

func inPath(node *rubik, path []rubik) bool {
	for i := range path {
		if node.cube == path[i].cube {
			return true
		}
	}
	return false
}

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
		if cost == 255 {
			// fmt.Printf("***************	END	********************\n")//
			// dumpPath(path)//
			return solution
		}
		// if t = ∞ then return NOT_FOUND
		bound = cost
		// fmt.Printf("bound 2: %v\n", bound)//
	}
	// return "Error"//
	// dumpCube(&path[0].cube)//
}

func search(path []rubik, g uint8, bound uint8, subgroup uint8) (uint8, string) {
	node := path[len(path) - 1]
	// fmt.Printf("Move: %v\n", &path[i].move)//
	// dumpCube(&node.cube)//
	// heuristic := heuristic(&node.cube, subgroup)
	if heuristic(&node.cube, subgroup) == 0 && subgroup < 3 {
		fmt.Printf("subgroup: %v\n", subgroup)//
		subgroup++
		var solvedPart string//
		for i := 1; i < len(path); i++ {//
			solvedPart += path[i].move + " "//
		}//
		fmt.Printf("solvedPart: %v\n", solvedPart)//
		// fmt.Printf("subgroup: %v\n", subgroup)//
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