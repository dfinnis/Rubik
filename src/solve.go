package rubik

import (
	"fmt"
)

type node struct {
	id			int
	cube		[6]uint32
	parent 		*node
	children	[]*node
}

func newNode(id int, newCube *[6]uint32, parent *node) *node {
	return &node{
		id:     id,
		cube:   *newCube,
		parent: parent,
	}
}


// 17 moves max
func heuristicG3(cube *[6]uint32) uint8 {
	var correct uint8
	// dumpCube(cube)/////
	var face uint32
	for face = 0; face < 6; face++ {
		// fmt.Printf("face: %v\n", face)///
		var cubie uint32
		var mask uint32 = 0x10000000
		for cubie = 0x70000000; cubie > 0; cubie /= 16 {
			// fmt.Printf("cubie: %x\n", cubie)//
			if cube[face]&cubie == mask * face {
				// fmt.Printf("CORRECT!!!!!!!!!!!\n")//
				correct++
			}
			mask /= 16
			// fmt.Printf("mask: %x\n", mask)//
		}
	}
	return 48 - correct
}

// generates a cube for each of the 6 possible moves
func generateMovesG3(root *node) {
	move := []string{
		"U2",
		"D2",
		"R2",
		"L2",
		"F2",
		"B2",
	}
	for i:= 0; i < 6; i++ {
		new := root.cube
		spin(move[i], &new)
		heuristic := heuristicG3(&new)
		fmt.Printf("heuristic = %v\n", heuristic)///
		newNode := newNode(i+1, &new, root)
		root.children = append(root.children, newNode)
	}
}

func printTree(root *node) {
	current := root
	fmt.Printf("id = %v\n", current.id)
	dumpCube(&current.cube)
	for i := range root.children {
		current := current.children[i]
		fmt.Printf("\n------------\n")
		fmt.Printf("parent = %v, id = %v\n", current.parent.id, current.id)
		dumpCube(&current.cube)
	}
	// dumpCube(&root.cube)//
}

func tree(cube *[6]uint32) {
	root := newNode(0, cube, nil)
	generateMovesG3(root)
	printTree(root)
	fmt.Printf("Oh finished!\n")
}

func solve(cube *[6]uint32) string {
	g4 := heuristicG3(cube)
	fmt.Printf("G4: %v\n", g4)
	tree(cube)
	solution := randomMix()/////////
	return solution
}