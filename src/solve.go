package rubik

import (
	"fmt"//
	"strings"
	"time"
)

// halfTurnMetric returns wc -w for given solution
func halfTurnMetric(sequence string) int {
	return len(strings.Fields(sequence))
}

// inPath returns true if given cube is already in the search path
func inPath(node *cepo, path []cepo) bool {
	for idx := range path {
		var different bool
		for i := range node.cP {
			if node.cP[i] != path[idx].cP[i] {
				different = true
				break
			}
		}
		if different == true {
			break
		}
		for i := range node.cO {
			if node.cO[i] != path[idx].cO[i] {
				different = true
				break
			}
		}
		if different == true {
			break
		}
		for i := range node.eP {
			if node.eP[i] != path[idx].eP[i] {
				different = true
				break
			}
		}
		if different == true {
			break
		}
		for i := range node.eO {
			if node.eO[i] != path[idx].eO[i] {
				different = true
				break
			}
		}
		if different == false {
			return true
		}
	}
	return false
}

// heuristic returns pruning table depth for given cube and subgroup
func heuristic(cube *cepo, subgroup int8, tables *tables) uint8 {
	if subgroup == 0 {
		return tables.G0[binaryToDecimal(cube.eO)]
	} else if subgroup == 1 {
		return tables.G1[eP2index(cube, tables)][cO2index(cube.cO)]
	} else if subgroup == 2 {
		return tables.G2[cP2index(cube)][eP2index8(cube, tables)]
	} else { // subgroup = 3
		idxCP := tables.G3cPindex[cP2index(cube)]
		idxEP := ePindexConverter(cube)
		return tables.G3[idxCP][idxEP[0]][idxEP[1]][idxEP[2]]
	}
}

// search recursively follows the pruning table heuristic until solved
func search(path []cepo, g uint8, bound uint8, subgroup int8, depth uint8, tables *tables) (uint8, string) {
	node := path[len(path) - 1]
	f := g + heuristic(&node, subgroup, tables)
	if f > bound {
		return f, ""
	}
	if heuristic(&node, subgroup, tables) == 0 {
		var solvedPart string
		for i := 1; i < len(path); i++ {
			solvedPart += path[i].move + " "
		}
		return 255, solvedPart
	}
	moves := listMoves(&node, subgroup)
	var min uint8 = 255 // ∞
	for i:= 0; i < len(moves); i++ {
		new := newNode(&node, moves[i])
		spin(moves[i], new)
		if inPath(new, path) == false {
			path = append(path, *new)
			cost, solution := search(path, g + heuristic(new, subgroup, tables), bound, subgroup, depth + 1, tables)
			if cost == 255 {
				return 255, solution
			}
			if cost < min {
				min = cost
			}
			path = path[:len(path) - 1] // pop
		}
	}
	return min, ""
}

// idaStar searches for a solution to the subgroup
func idaStar(cube *cepo, subgroup int8, tables *tables) string {
	bound := heuristic(cube, subgroup, tables)
	var path []cepo
	path = append(path, *cube)
	for {
		cost, solution := search(path, 0, bound, subgroup, 0, tables)
		if cost == 255 {
			return solution
		}
		bound = cost
	}
}

// solve calls idaStar search for each subgroup
func solve(cube *cepo, tables *tables, group bool) string {
	fmt.Printf("\nSolving")
	var solution string
	for subgroup := isSubgroup(cube); subgroup < 4; subgroup++ {
		cube.move = ""
		cube.move2 = ""
		start := time.Now()
		solutionPart := idaStar(cube, subgroup, tables)
		if group {
			elapsed := time.Since(start)
			dumpCube(cube)//
			fmt.Printf("\n%vSubgroup: %v%v\n", Bright, subgroup, Reset)
			fmt.Printf("Solution: %v\n", solutionPart)
			fmt.Printf("HTM:      %v\n", halfTurnMetric(solutionPart))
			fmt.Printf("Time:     %v\n", elapsed)
		} else {
			fmt.Printf(".")
		}
		spin(solutionPart, cube)
		solution += solutionPart
	}
	if group {
		dumpCube(cube)//
		fmt.Println()
	}
	solution = trim(solution)
	return solution
}
