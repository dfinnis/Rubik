package rubik

import (
	"fmt"
)

// type node struct {
// 	// id			int
// 	cube		[6]uint32
// 	// parent 		*node
// 	// children	[]*node
// }

// func newNode(/*id int, */newCube *[6]uint32/*, parent *node*/) *node {
// 	return &node{
// 		// id:     id,
// 		cube:   *newCube,
// 		// parent: parent,
// 	}
// }

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

// // generates a cube for each of the 6 possible moves
// func generateMovesG3(cube *[6]uint32) string {
// 	root := newNode(/*0, */cube/*, nil*/)
// 	move := []string{
// 		"U2",
// 		"D2",
// 		"R2",
// 		"L2",
// 		"F2",
// 		"B2",
// 	}
// 	var bestHeuristic uint8 = 255
// 	var bestCube [6]uint32
// 	var bestMove string
// 	// dumpCube(&root.cube)//
// 	for i:= 0; i < 6; i++ {
// 		new := root.cube
// 		spin(move[i], &new)
// 		heuristic := heuristicG3(&new)
// 		if heuristic < bestHeuristic {
// 			bestHeuristic = heuristic
// 			bestCube = new
// 			bestMove = move[i]
// 			fmt.Printf("\n\n---- best = %v ----", bestHeuristic)//

// 		}
// 		fmt.Printf("\nMove: %v\n", move[i])//
// 		fmt.Printf("heuristic = %v\n", heuristic)///
// 		// dumpCube(&new)//
// 		// fmt.Printf("Best:")///
// 		// fmt.Printf("Best Move: %v\n", bestMove)///
// 		// dumpCube(&bestCube)//
// 		// fmt.Printf("#######################################\n")///
// 		// newNode := newNode(i+1, &new, root)
// 		// root.children = append(root.children, newNode)
// 	}
// 	fmt.Printf("Best Move: %v\n", bestMove)///
// 	fmt.Printf("#######################################\n")///

// 	if bestHeuristic != 0 {
// 		// fmt.Printf("\nAAAAAAARRRRGGGGHHHHHH! not solved yet\n")//
// 		recursive := generateMovesG3(&bestCube)
// 		// fmt.Printf("\nrecursive solution: %v\n", recursive)//
// 		bestMove += " " + recursive
// 	}
// 	// fmt.Printf("#######################################\n")///å
// 	return bestMove
// }

// func tree(cube *[6]uint32) {
// 	// root := newNode(/*0, */cube/*, nil*/)
// 	generateMovesG3(cube)
// 	// printTree(root)
// 	fmt.Printf("Oh finished!\n")
// }

// func solve(cube *[6]uint32) string {
// 	if isSolved(cube) {
// 		return ""
// 	}
// 	// g4 := heuristicG3(cube)
// 	// fmt.Printf("G4: %v\n", g4)
// 	// tree(cube)
// 	idaStar(cube)
// 	// recursiveG3(cube)
// 	// solution := generateMovesG3(cube)
// 	// fmt.Printf("\n\nSOLUTION HERE : %v\n", solution)/////
// 	solution := randomMix()/////////
// 	return solution
// }

func solve(r *rubik) string {
	if isSolved(&r.cube) {
		return ""
	}
	idaStar(r)
	solution := randomMix()/////////
	return solution
}

func idaStar(r *rubik) {
	var path []rubik
	var bound uint8 = heuristicG3(&r.cube)
	fmt.Printf("bound: %v\n", bound)//
	path = append(path, *r)
	for {
		cost := search(path, 0, bound)
		// if t = FOUND then return (path, bound)
		// if t = ∞ then return NOT_FOUND
		bound = cost
		fmt.Printf("bound 2: %v\n", bound)//
		break//
	}
	dumpCube(&path[0].cube)//
}

func search(path []rubik, g uint8, bound uint8) uint8 {
	return 42
	// node := path // last
	// f := g + heuristicG3(node)
	// if f > bound {
	// 	return f
	// }
	// if isSolved(node) {
	// 	return FOUND
	// }
	// min := 0x77777777
	// root := newNode(/*0, */node/*, nil*/)
// 	move := []string{
// 		"U2",
// 		"D2",
// 		"R2",
// 		"L2",
// 		"F2",
// 		"B2",
// 	}
// 	for i:= 0; i < 6; i++ {
// 		new := root
// 		spin(move[i], &new)
// 		heuristic := heuristicG3(&new)
// 	}
}

// func idaStar(root *[6]uint32) {
// 	// var solution []rubik
// 	var path []rubik
// 	// solution = append(solution, root)
// 	var bound uint8 = heuristicG3(root)
// 	fmt.Printf("bound: %v\n", bound)//
// 	// path = append(path, root)
// 		// var solution []rubik
// 	// path = append(path, root)
// 	for {
// 		cost := search(path, 0, bound)
// 		// if t = FOUND then return (path, bound)
// 		// if t = ∞ then return NOT_FOUND
// 		bound = cost
// 		break//
// 	}
// 	dumpCube(path)
// }

// func search(path, g, bound) {
// 	// node := path // last
// 	// f := g + heuristicG3(node)
// 	// if f > bound {
// 	// 	return f
// 	// }
// 	// if isSolved(node) {
// 	// 	return FOUND
// 	// }
// 	// min := 0x77777777
// 	// root := newNode(/*0, */node/*, nil*/)
// // 	move := []string{
// // 		"U2",
// // 		"D2",
// // 		"R2",
// // 		"L2",
// // 		"F2",
// // 		"B2",
// // 	}
// // 	for i:= 0; i < 6; i++ {
// // 		new := root
// // 		spin(move[i], &new)
// // 		heuristic := heuristicG3(&new)
// // 	}
// }

// 	for succ in successors(node) do
// 	if succ not in path then
// 		path.push(succ)
// 		t := search(path, g + cost(node, succ), bound)
// 		if t = FOUND then return FOUND
// 		if t < min then min := t
// 		path.pop()
// 	end if
// 	end for
// 	return min
// 	end function
// }