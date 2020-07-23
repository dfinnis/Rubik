package rubik

import (
	"fmt"
	// "reflect"
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

// G0 = U, D, L, R, F, B
// G1 = U, D, L2, R2, F2, B2 
// G2 = solved

func generateG0Moves(current *node) {
	move := []string{
		"U",
		"D",
		"R",
		"L",
		"F",
		"B",
	}
	for i:= 0; i < 6; i++ {
		mix := move[i] // tested, works
		new := current.cube
		spin(mix, &new)
		newNode := newNode(i+1, &new, current)
		current.children = append(current.children, newNode)
	}
}

// G1 uses a set of the 18 possible moves that, regardless of order or number of moves, 
// cannot change the orientation of edge and corner pieces
func generateG1Moves(current *node) {
	move := []string{
		"U",
		"D",
		"R2",
		"L2",
		"F2",
		"B2",
	}
	for i:= 0; i < 6; i++ {
		mix := move[i] // tested, works
		new := current.cube
		spin(mix, &new)
		newNode := newNode(i+1, &new, current)
		current.children = append(current.children, newNode)
	}
}

// generates a cube for each of the 18 possible moves
func generateMoves(root *node) {
	move := []string{
		"U",
		"U'",
		"U2",
		"D",
		"D'",
		"D2",
		"R",
		"R'",
		"R2",
		"L",
		"L'",
		"L2",
		"F",
		"F'",
		"F2",
		"B",
		"B'",
		"B2",
	}
	for i:= 0; i < 18; i++ {
		mix := move[i] // tested, works
		new := root.cube
		spin(mix, &new)
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
	dumpCube(&root.cube)
}

func tree(cube *[6]uint32) {

	root := newNode(0, cube, nil)
	generateMoves(root)
	printTree(root)


}
