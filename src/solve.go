package rubik

import (
	"fmt"
	"time"
	"strings"
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
	var bound uint8 = heuristic(&r.cube, subgroup)
	var path []rubik
	path = append(path, *r)
	for {
		cost, solution := search(path, 0, bound, subgroup, 0)
		if cost == 255 {
			return solution
		}
		// if t = ∞ then return NOT_FOUND
		bound = cost
	}
	// dumpCube(&path[0].cube)//
}

func search(path []rubik, g uint8, bound uint8, subgroup uint8, depth uint8) (uint8, string) {
	node := path[len(path) - 1]
	f := g + heuristic(&node.cube, subgroup)
	// fmt.Printf("f: %v\n", f)
	if f > bound {
		return f, ""
	}
	// fmt.Printf("g = %v\n", g)//
	// fmt.Printf("depth = %v\n", depth)//
	if heuristic(&node.cube, subgroup) == 0 {
		var solvedPart string
		for i := 1; i < len(path); i++ {
			solvedPart += path[i].move + " "
		}
		// fmt.Printf("subgroup: %v\n", subgroup)//
		fmt.Printf("heuristic(&node.cube, subgroup): %v\n", heuristic(&node.cube, subgroup))////////
		fmt.Printf("heuristicG0(&node.cube): %v\n", heuristicG0(&node.cube))////////
		fmt.Printf("heuristicG1(&node.cube): %v\n", heuristicG1(&node.cube))////////
		fmt.Printf("heuristicG2(&node.cube): %v\n", heuristicG2(&node.cube))////////
		fmt.Printf("heuristicG3(&node.cube): %v\n", heuristicG3(&node.cube))////////
		// fmt.Printf("depth = %v\n", depth)//
		return 255, solvedPart
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
			cost, solution := search(path, g + heuristic(&new.cube, subgroup) + 1/* + 1 */, bound, subgroup, depth + 1) // g + h + 1?
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

func halfTurnMetric(sequence string) int {
	return len(strings.Fields(sequence))
}

func solve(r *rubik) string {
	// dumpCube(&r.cube)//
	dumpCube(&r.cube)//
	fmt.Printf("heuristicG0(&r.cube): %v\n", heuristicG0(&r.cube))////////
	fmt.Printf("heuristicG1(&r.cube): %v\n", heuristicG1(&r.cube))////////
	fmt.Printf("heuristicG2(&r.cube): %v\n", heuristicG2(&r.cube))////////
	fmt.Printf("heuristicG3(&r.cube): %v\n", heuristicG3(&r.cube))////////

	var solution string
	for subgroup := subgroup(&r.cube); subgroup < 4; subgroup++ {
		fmt.Printf("\nsubgroup: %v\n", subgroup)////////
		// fmt.Printf("heuristicG0(&r.cube): %v\n", heuristicG0(&r.cube))////////
		// fmt.Printf("heuristicG1(&r.cube): %v\n", heuristicG1(&r.cube))////////
		// fmt.Printf("heuristicG2(&r.cube): %v\n", heuristicG2(&r.cube))////////
		// fmt.Printf("heuristicG3(&r.cube): %v\n", heuristicG3(&r.cube))////////
		start := time.Now()//
		solutionPart := idaStar(r, subgroup)
		elapsed := time.Since(start)//
		fmt.Printf("Group Solve time: %v\n", elapsed)//
		fmt.Printf("solutionPart: %v\n", solutionPart)//
		fmt.Printf("Half Turn Metric = %v\n", halfTurnMetric(solutionPart))//
		spin(solutionPart, &r.cube)
		dumpCube(&r.cube)//
		solution += solutionPart
		if isSolved(&r.cube) {
			break
		}
		// if subgroup == 0 {//
		// 	break//
		// }//
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