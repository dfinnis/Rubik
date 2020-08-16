package rubik

import (
	"fmt"//
)

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

func solveG0(parent *cepo, tableG0 [4096]uint8) string {
	var solution string
	index := binaryToDecimal(parent.eO)
	parentDepth := tableG0[index]
	fmt.Printf("index: %v\n", index)//
	fmt.Printf("depth: %v\n\n", parentDepth)//

	// for isSubgroup(parent) == 0 {
	for i, move := range listMovesCepo(parent, 0) {
		fmt.Printf("move[%v]: %v\n", i, move)//
		child := newNodeCepo(parent, move)
		spinCepo(move, child)
		index := binaryToDecimal(child.eO)
		fmt.Printf("index: %v\n", index)//
		childDepth := tableG0[index]
		fmt.Printf("childDepth: %v\n", childDepth)//
		if childDepth < parentDepth {
			fmt.Printf("OH HIII!!\n\n")//
			solution += move + " "
			fmt.Printf("solution: %v\n", solution)////////
			parent = child
			parentDepth = childDepth
			break
		}
	}
	// }
	dumpCepo(parent)
	
	// dumpCepo(cube)//
	return solution
	// return "F U"
}

func solveCepo(cube *cepo, tableG0 [4096]uint8) string {
	// fmt.Printf("tableG0: %v\n", tableG0)//
	subgroup := isSubgroup(cube)
	fmt.Printf("\nsubgroup initally: %v\n", subgroup)//
	// solution := "F U"//

	var solution string
	for subgroup := isSubgroup(cube); subgroup < 4; subgroup++ {
		fmt.Printf("\nsubgroup: %v\n", subgroup)////////
		if subgroup == 0 {
			solutionPart := solveG0(cube, tableG0)
			fmt.Printf("solutionPart: %v\n", solutionPart)////////
		}
		// spinCepo(solutionPart, cube)
		// solution += solutionPart
		// if isSolvedCepo(cube) {
		// 	break
		// }
	}

	// fmt.Printf("\n\nSolution pre-trim: %v\n", solution)///
	// fmt.Printf("HTM pre-trim: %v\n", halfTurnMetric(solution))///
	// solution = trim(solution)











	return solution
}