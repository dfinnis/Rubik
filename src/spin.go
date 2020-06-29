package rubik

import (
	"fmt" //
	"math/bits"
	"strings"
)

func spinFace(cube *[6]uint32, face uint8) {
	// cube[0] |= 0x1 // !@!!!!!!!!!!!
	cube[face] = bits.RotateLeft32(cube[face], -8)
}

func spinU(cube *[6]uint32) {
	spinFace(cube, 0)
	// spin edges
	tmp0 := cube[1]
	tmp0 &= 0x77700000

	tmp1 := cube[2]
	tmp1 &= 0x77700000
	cube[1] &= 0x77777
	cube[1] |= tmp1
	
	tmp1 = cube[3]
	tmp1 &= 0x77700000
	cube[2] &= 0x77777
	cube[2] |= tmp1

	tmp1 = cube[4]
	tmp1 &= 0x77700000
	cube[3] &= 0x77777
	cube[3] |= tmp1

	cube[4] &= 0x77777
	cube[4] |= tmp0
}

func spin(mix string, cube *[6]uint32) {
	sequence := strings.Fields(mix)
	fmt.Printf("\nsequence: %v, len: %d\n", sequence, len(sequence)) //
	for spin := 0; spin < len(sequence); spin++ {
		fmt.Printf("\nspin: %v\n", spin) //
		fmt.Printf("\nspin: %v\n", sequence[spin]) //
		if sequence[spin] == "U" {
			spinU(&r.cube)
	// 	} else if sequence[spin] == "U'" {
	// 		spinUa(&r.cube)
	// 	} else if sequence[spin] == "U2" {
	// 		spinU(&r.cube)
	// 		spinU(&r.cube)
	// 	} else if sequence[spin] == "D" {
	// 		spinD(&r.cube)
	// 	} else if sequence[spin] == "D'" {
	// 		spinDa(&r.cube)
	// 	} else if sequence[spin] == "D2" {
	// 		spinD(&r.cube)
	// 		spinD(&r.cube)
	// 	} else if sequence[spin] == "R" {
	// 		spinR(&r.cube)
	// 	} else if sequence[spin] == "R'" {
	// 		spinRa(&r.cube)
	// 	} else if sequence[spin] == "R2" {
	// 		spinR(&r.cube)
	// 		spinR(&r.cube)
	// 	} else if sequence[spin] == "L" {
	// 		spinL(&r.cube)
	// 	} else if sequence[spin] == "L'" {
	// 		spinLa(&r.cube)
	// 	} else if sequence[spin] == "L2" {
	// 		spinL(&r.cube)
	// 		spinL(&r.cube)
	// 	} else if sequence[spin] == "F" {
	// 		spinF(&r.cube)
	// 	} else if sequence[spin] == "F'" {
	// 		spinFa(&r.cube)
	// 	} else if sequence[spin] == "F2" {
	// 		spinF(&r.cube)
	// 		spinF(&r.cube)
	// 	} else if sequence[spin] == "B" {
	// 		spinB(&r.cube)
	// 	} else if sequence[spin] == "B'" {
	// 		spinBa(&r.cube)
	// 	} else if sequence[spin] == "B2" {
	// 		spinB(&r.cube)
	// 		spinB(&r.cube)
	// 	} else {
	// 		errorExit("bad input")
		}
	}
	dumpCube(cube)////
	// test()//////!!!!!
}
