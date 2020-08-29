package rubik

import (
	"fmt"//
	"strings"
)

func halfTurnMetric(sequence string) int {
	return len(strings.Fields(sequence))
}

func isSubgroup(cube *cepo) int8 {
	for i := range cube.eO { // edges not oriented -> 0
		if cube.eO[i] != 0 {
			return 0
		}
	}
	for i := range cube.cO { // corners not oriented -> 1
		if cube.eO[i] != 0 {
			return 1
		}
	}
	for i := 8; i <= 11; i++ { // LR-edges(8-11) not in their slice -> 1
		if cube.eP[i] < 8 {
			return 1
		}
	}
	for i := 0; i <= 3; i++ { // corners not in tetrads -> 2
		if cube.cP[i] > 3 {
			return 2
		}
	}
	for i := 4; i <= 7; i++ { // corners not in tetrads -> 2
		if cube.cP[i] < 4 {
			return 2
		}
	}
	for i := 0; i <= 3; i++ { // edges not in their slice -> 2
		if cube.eP[i] > 3 {
			return 2
		}
	}
	for i := 4; i <= 7; i++ { // edges not in their slice -> 2
		if cube.eP[i] < 4 || cube.eP[i] > 7 {
			return 2
		}
	}
	var parity int8
	for i := range cube.cP { // corners parity odd -> 2
		if cube.cP[i] == int8(i) {
			parity++
		}
	}
	if parity % 2 != 0 {
		return 2
	}
	// parity = 0
	// for i := range cube.eP { // edges parity odd -> 2
	// 	if cube.eP[i] == int8(i) {
	// 		parity++
	// 	}
	// }
	// if parity % 2 != 0 {
	// 	return 2
	// }
	if isSolvedCepo(cube) == false {
		return 3
	}
	return 4
}

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

func heuristic(cube *cepo, subgroup int8, tables *tables) uint8 {
	if subgroup == 0 {
		return tables.G0[binaryToDecimal(cube.eO)]
	} else if subgroup == 1 {
		return tables.G1[eP2index(cube, tables)][cO2index(cube)]
	} else {// if subgroup == 2 {
		return tables.G2[cP2index(cube)][eP2index8(cube, tables)]
	}
}

func search(path []cepo, g uint8, bound uint8, subgroup int8, depth uint8, tables *tables) (uint8, string) {
	node := path[len(path) - 1]
	// f := g + tables.G0[binaryToDecimal(node.eO)]
	f := g + heuristic(&node, subgroup, tables)
	// fmt.Printf("g: %v\n", g)//
	// fmt.Printf("f: %v\n", f)//

	if f > bound {
		return f, ""
	}
	// fmt.Printf("g = %v\n", g)//
	// fmt.Printf("subgroup = %v\n", subgroup)//
	// fmt.Printf("depth = %v\n", depth)//
	// if tables.G0[binaryToDecimal(node.eO)] == 0 {
	if heuristic(&node, subgroup, tables) == 0 {
		var solvedPart string
		for i := 1; i < len(path); i++ {
			solvedPart += path[i].move + " "
		}
		// fmt.Printf("subgroup: %v\n", subgroup)//
		// fmt.Printf("depth = %v\n", depth)//
		return 255, solvedPart
	}
	moves := listMoves(&node, subgroup)
	var min uint8 = 255 // ∞
	for i:= 0; i < len(moves); i++ {
		new := newNode(&node, moves[i])
		spin(moves[i], new)
		// fmt.Printf("Move: %v\n", new.move)//
		// dumpCube(&new.cube)//
		if inPath(new, path) == false {
			path = append(path, *new)
			// dumpPath(path)//
			cost, solution := search(path, g + heuristic(new, subgroup, tables), bound, subgroup, depth + 1, tables) // g + h + 1?
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

func idaStar(cube *cepo, subgroup int8, tables *tables) string {
	bound := heuristic(cube, subgroup, tables)
	var path []cepo
	path = append(path, *cube)
	for {
		cost, solution := search(path, 0, bound, subgroup, 0, tables)
		// fmt.Printf("cost: %v\n", cost)///
		if cost == 255 {
			return solution
		}
		bound = cost
	}
}

func solve(cube *cepo, tables *tables) string {
	// fmt.Printf("tableG0: %v\n", tableG0)//
	subgroup := isSubgroup(cube)
	fmt.Printf("\nsubgroup initally: %v\n", subgroup)//
	// solution := "F U"//

	var solution string
	for subgroup := isSubgroup(cube); subgroup < 4; subgroup++ {
		cube.move = ""
		cube.move2 = ""
		fmt.Printf("\nsubgroup: %v\n", subgroup)////////
		// dumpCube(cube)////////////////////////////////////////////////////##########
		// start := time.Now()
		solutionPart := idaStar(cube, subgroup, tables)
		spin(solutionPart, cube)
		solution += solutionPart
		// elapsed := time.Since(start)//
		// fmt.Printf("Group Solve time: %v\n", elapsed)//
		fmt.Printf("solutionPart: %v\n", solutionPart)//
		fmt.Printf("Half Turn Metric = %v\n", halfTurnMetric(solutionPart))//
		// if isSolved(cube) {
		// 	break
		// }
		if subgroup > 2 {//
			// fmt.Printf("G3: binaryToDecimal12: %v\n", binaryToDecimal12(cube.eP))//
			break
		}
	}

	fmt.Printf("\n\nSolution pre-trim: %v\n", solution)///
	fmt.Printf("HTM pre-trim: %v\n", halfTurnMetric(solution))///
	solution = trim(solution)

	return solution
}
