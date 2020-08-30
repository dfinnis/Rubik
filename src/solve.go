package rubik

import (
	"fmt"//
	"strings"
	"time"
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
	} else if subgroup == 2 {
		return tables.G2[cP2index(cube)][eP2index8(cube, tables)]
	} else { // subgroup = 3
		idxCP := tables.G3cPindex[cP2index(cube)]
		idxEP := ePindexConverter(cube)
		return tables.G3[idxCP][idxEP[0]][idxEP[1]][idxEP[2]]
	}
}

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

func solve(cube *cepo, tables *tables, group bool) string {
	var solution string
	for subgroup := isSubgroup(cube); subgroup < 4; subgroup++ {
		cube.move = ""
		cube.move2 = ""
		start := time.Now()
		solutionPart := idaStar(cube, subgroup, tables)
		elapsed := time.Since(start)//
		if group {
			fmt.Printf("\n%vSubgroup: %v%v\n", "\x1B[1m", subgroup, "\x1B[0m")////////
			fmt.Printf("Solution: %v\n", solutionPart)//
			fmt.Printf("HTM:      %v\n", halfTurnMetric(solutionPart))//
			fmt.Printf("Time:     %v\n", elapsed)//
		}
		spin(solutionPart, cube)
		solution += solutionPart
	}
	if group {
		fmt.Println()
	}
	solution = trim(solution)
	return solution
}
