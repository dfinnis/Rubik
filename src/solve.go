package rubik

import (
	"fmt"
)

type node struct {
	// id			int
	cube		[6]uint32
	// parent 		*node
	// children	[]*node
}

func newNode(/*id int, */newCube *[6]uint32/*, parent *node*/) *node {
	return &node{
		// id:     id,
		cube:   *newCube,
		// parent: parent,
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
	return 48 - correct
}

// generates a cube for each of the 6 possible moves
func generateMovesG3(cube *[6]uint32) {
	root := newNode(/*0, */cube/*, nil*/)
	move := []string{
		"U2",
		"D2",
		"R2",
		"L2",
		"F2",
		"B2",
	}
	var bestHeuristic uint8 = 255
	var bestCube [6]uint32
	var bestMove string
	dumpCube(&root.cube)//
	for i:= 0; i < 6; i++ {
		new := root.cube
		spin(move[i], &new)
		heuristic := heuristicG3(&new)
		if heuristic < bestHeuristic {
			bestHeuristic = heuristic
			bestCube = new
			bestMove = move[i]
			fmt.Printf("\n\n---- best = %v ----", bestHeuristic)//

		}
		fmt.Printf("\nMove: %v\n", move[i])//
		fmt.Printf("heuristic = %v\n", heuristic)///
		dumpCube(&new)//
		fmt.Printf("Best:")///
		fmt.Printf("Best Move: %v\n", bestMove)///
		dumpCube(&bestCube)//
		fmt.Printf("#######################################")///
		// newNode := newNode(i+1, &new, root)
		// root.children = append(root.children, newNode)
	}
}

// func printTree(root *node) {
// 	current := root
// 	fmt.Printf("id = %v\n", current.id)
// 	dumpCube(&current.cube)
// 	for i := range root.children {
// 		current := current.children[i]
// 		fmt.Printf("\n------------\n")
// 		fmt.Printf("parent = %v, id = %v\n", current.parent.id, current.id)
// 		dumpCube(&current.cube)
// 	}
// 	// dumpCube(&root.cube)//
// }

// func tree(cube *[6]uint32) {
// 	// root := newNode(/*0, */cube/*, nil*/)
// 	generateMovesG3(cube)
// 	// printTree(root)
// 	fmt.Printf("Oh finished!\n")
// }

func solve(cube *[6]uint32) string {
	// g4 := heuristicG3(cube)
	// fmt.Printf("G4: %v\n", g4)
	// tree(cube)
	// idaStar(cube)
	// recursiveG3(cube)
	generateMovesG3(cube)
	solution := randomMix()/////////
	return solution
}

// func solve(r *rubik) string {
// 	// g4 := heuristicG3(cube)
// 	// fmt.Printf("G4: %v\n", g4)
// 	// tree(cube)
// 	// idaStar(cube)
// 	recursiveG3(r)
// 	solution := randomMix()/////////
// 	return solution
// }

// func recursiveG3(r *rubik) {
// 	move := []string{
// 		"U2",
// 		"D2",
// 		"R2",
// 		"L2",
// 		"F2",
// 		"B2",
// 	}
// 	dumpCube(&r.cube)//
// 	fmt.Printf("------------------------------------------------------------\n\n")//
// 	for i:= 0; i < 6; i++ {
// 		new := &r.cube
// 		spin(move[i], new)
// 		fmt.Printf("\n\nMove: %v\n", move[i])//
// 		heuristic := heuristicG3(new)
// 		fmt.Printf("heuristic = %v\n", heuristic)///
// 		dumpCube(new)//
// 		// newNode := newNode(i+1, &new, root)
// 		// root.children = append(root.children, newNode)
// 	}
// }





// func idaStar(root *[6]uint32) {
// 	var bound uint8 = 255
// 	path := root
// 	// path = append(path, root)
// 	for {
// 		cost := search()
// 	}
// 	// fmt.Println("Oh HIIII!!!s")//
// 	dumpCube(path)
// }


// path              current search path (acts like a stack)
//  node              current node (last node in current path)
//  g                 the cost to reach current node
//  f                 estimated cost of the cheapest path (root..node..goal)
//  h(node)           estimated cost of the cheapest path (node..goal)
//  cost(node, succ)  step cost function
//  is_goal(node)     goal test
//  successors(node)  node expanding function, expand nodes ordered by g + h(node)
//  ida_star(root)    return either NOT_FOUND or a pair with the best path and its cost
 
//  procedure ida_star(root)
//    bound := h(root)
//    path := [root]
//    loop
//      t := search(path, 0, bound)
//      if t = FOUND then return (path, bound)
//      if t = ∞ then return NOT_FOUND
//      bound := t
//    end loop
//  end procedure
 
//  function search(path, g, bound)
//    node := path.last
//    f := g + h(node)
//    if f > bound then return f
//    if is_goal(node) then return FOUND
//    min := ∞
//    for succ in successors(node) do
//      if succ not in path then
//        path.push(succ)
//        t := search(path, g + cost(node, succ), bound)
//        if t = FOUND then return FOUND
//        if t < min then min := t
//        path.pop()
//      end if
//    end for
//    return min
//  end function