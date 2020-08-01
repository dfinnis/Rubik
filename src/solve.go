package rubik

import (
	"fmt"
	// "os"//
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

func idaStar(r *rubik, subgroup uint8) string {
	// var subgroup uint8 = 0 // test heuristics to establish subgroup
	// if heuristicG0(&r.cube) == 0 {
	// 	subgroup = 1
	// 	if heuristicG1(&r.cube) == 0 {
	// 		subgroup = 2
	// 		if heuristicG2(&r.cube) == 0 {
	// 			subgroup = 3
	// 		}
	// 	}
	// }
	// fmt.Printf("\nsubgroup: %v\n", subgroup)////////
	var bound uint8 = heuristic(&r.cube, subgroup)
	// fmt.Printf("bound: %v\n", bound)//
	// dumpCube(&r.cube)////////

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
		// break/////
	}
	// dumpCube(&path[0].cube)//
}

func search(path []rubik, g uint8, bound uint8, subgroup uint8) (uint8, string) {
	node := path[len(path) - 1]
	// fmt.Printf("Move: %v\n", &path[i].move)//
	// dumpCube(&node.cube)//
	// heuristic := heuristic(&node.cube, subgroup)
	// if heuristic(&node.cube, subgroup) == 0 && subgroup < 3 {
	// 	fmt.Printf("subgroup: %v\n", subgroup)//
	// 	subgroup++
	// 	// var solvedPart string//
	// 	// for i := 1; i < len(path); i++ {//
	// 	// 	solvedPart += path[i].move + " "//
	// 	// }//
	// 	// fmt.Printf("solvedPart: %v\n", solvedPart)//
	// 	// fmt.Printf("subgroup: %v\n", subgroup)//
	// }
	f := g + heuristic(&node.cube, subgroup)
	// fmt.Printf("f: %v\n", f)
	if f > bound {
		return f, ""
	}
	if heuristic(&node.cube, subgroup) == 0 {
		var solvedPart string
		for i := 1; i < len(path); i++ {
			solvedPart += path[i].move + " "
		}
		// fmt.Printf("solvedPart: %v\n", solvedPart)//
		return 255, solvedPart
	}
	// if isSolved(&node.cube) {
	// 	var solved string
	// 	for i := 1; i < len(path); i++ {
	// 		solved += path[i].move + " "
	// 	}
	// 	return 255, solved // FOUND
	// }
	move := listMoves(&node, subgroup)
	var min uint8 = 255 // ∞
	for i:= 0; i < len(move); i++ {
		new := newNode(&node.cube, move[i])
		spin(move[i], &new.cube)
		// fmt.Printf("Move: %v\n", new.move)//
		// dumpCube(&new.cube)//
		// fmt.Printf("heursiticG1: %v\n", heuristicG1(&new.cube))
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
	// os.Exit(1)///
	return min, ""
}

// test heuristics to establish subgroup
func subgroup(cube *[6]uint32) uint8 {
	var subgroup uint8
	if heuristicG0(&r.cube) == 0 {
		subgroup = 1
		if heuristicG1(&r.cube) == 0 {
			subgroup = 2
			if heuristicG2(&r.cube) == 0 {
				subgroup = 3
			}
		}
	}
	return subgroup
}

func solve(r *rubik) string {
	// dumpCube(&r.cube)//
	//////establish subgroup first then give subgroup to idastar()!!!!!!!!
	var solution string

	for subgroup := subgroup(&r.cube); subgroup < 4; subgroup++ {
		fmt.Printf("\nsubgroup: %v\n", subgroup)////////
		solutionPart := idaStar(r, subgroup)
		fmt.Printf("solutionPart: %v\n", solutionPart)//
		dumpCube(&r.cube)//
		spin(solutionPart, &r.cube)
		solution += solutionPart
		if isSolved(&r.cube) {
			break
		}
	}

	// var solution string
	// for isSolved(&r.cube) == false {
	// 	solutionPart := idaStar(r)
	// 	// fmt.Printf("solutionPart: %v\n", solutionPart)//
	// 	// dumpCube(&r.cube)//
	// 	spin(solutionPart, &r.cube)
	// 	solution += solutionPart
	// }
	return solution
}