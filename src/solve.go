package rubik

import (
	"fmt"
)

// 17 moves max
func G4heuristic(cube *[6]uint32) uint8 {
	var correct uint8
	dumpCube(cube)/////
	var face uint32
	for face = 0; face < 6; face++ {
		fmt.Printf("face: %v\n", face)///
		var cubie uint32
		var mask uint32 = 0x10000000
		for cubie = 0x70000000; cubie > 0; cubie /= 16 {
			fmt.Printf("cubie: %x\n", cubie)//
			if cube[face]&cubie == mask * face {
				fmt.Printf("CORRECT!!!!!!!!!!!\n")//
				correct++
			}
			mask /= 16
			fmt.Printf("mask: %x\n", mask)//
		}
	}
	return correct
}

func solve(cube *[6]uint32) string {
	g4 := G4heuristic(cube)
	fmt.Printf("G4: %v\n", g4)
	solution := randomMix()/////////
	return solution
}