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
		// either init new cube and spin it, or copy root cube after spin
		new := initRubik()
		// fmt.Printf("\n\n## move %s ##\n", mix)
		spin(mix, &new.cube)
		newNode := newNode(i+1, &new.cube, root)
		root.children = append(root.children, newNode)
	}
}

func tree(cube *[6]uint32) {

	// start_cube := initRubik()
	root := newNode(0, cube, nil)
	generateMoves(root)

	current := root
	fmt.Printf("id = %v\n", current.id)
	dumpCube(&current.cube)
	for i := range root.children {
		current := current.children[i]
		fmt.Printf("\n------------\n")
		fmt.Printf("parent = %v, id = %v\n", current.parent.id, current.id)
		dumpCube(&current.cube)
		
	}
}
