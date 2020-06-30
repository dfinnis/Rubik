package rubik

import (
	// "fmt"
)

type node struct {
	id     int
	cube   [6]uint32
	parent *node
	// children         []*node
}

func newNode(id int, newCube *[6]uint32, parent *node) *node {
	return &node{
		id:     id,
		cube:   *newCube,
		parent: parent,
	}
}

// generates a cube for each of the 18 possible moves
func generateCubes() {

}

func tree(cube *[6]uint32) {

	// start_cube := initRubik()
	root := newNode(0, cube, nil)
	dumpCube(&root.cube)
	
	generateCubes()
}
