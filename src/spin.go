package rubik

import (
	"fmt" //
	"math/bits"
	"strings"
)

func spinFace(cube *[6]uint32, face uint8) {
	cube[face] = bits.RotateLeft32(cube[face], -8)
}

func spinFaceAnti(cube *[6]uint32, face uint8) {
	cube[face] = bits.RotateLeft32(cube[face], 8)
}

func spinFace2(cube *[6]uint32, face uint8) {
	cube[face] = bits.RotateLeft32(cube[face], 16)
}

func spinU(cube *[6]uint32) {
	spinFace(cube, 0)
	// spin edges
	swap := cube[1]
	swap &= 0x77700000

	tmp := cube[2]
	tmp &= 0x77700000
	cube[1] &= 0x77777
	cube[1] |= tmp
	
	tmp = cube[3]
	tmp &= 0x77700000
	cube[2] &= 0x77777
	cube[2] |= tmp

	tmp = cube[4]
	tmp &= 0x77700000
	cube[3] &= 0x77777
	cube[3] |= tmp

	cube[4] &= 0x77777
	cube[4] |= swap
}

func spinUa(cube *[6]uint32) {
	spinFaceAnti(cube, 0)
	// spin edges
	swap := cube[1]
	swap &= 0x77700000

	tmp := cube[4]
	tmp &= 0x77700000
	cube[1] &= 0x77777
	cube[1] |= tmp
	
	tmp = cube[3]
	tmp &= 0x77700000
	cube[4] &= 0x77777
	cube[4] |= tmp

	tmp = cube[2]
	tmp &= 0x77700000
	cube[3] &= 0x77777
	cube[3] |= tmp

	cube[2] &= 0x77777
	cube[2] |= swap
}

func spinU2(cube *[6]uint32) {
	spinFace2(cube, 0)
	// spin edges
	swap := cube[1]
	swap &= 0x77700000

	tmp := cube[3]
	tmp &= 0x77700000
	cube[1] &= 0x77777
	cube[1] |= tmp
	
	cube[3] &= 0x77777
	cube[3] |= swap

	swap = cube[2]
	swap &= 0x77700000

	tmp = cube[4]
	tmp &= 0x77700000
	cube[2] &= 0x77777
	cube[2] |= tmp
	
	cube[4] &= 0x77777
	cube[4] |= swap
}

func spinD(cube *[6]uint32) {
	spinFace(cube, 5)
	// spin edges
	swap := cube[1]
	swap &= 0x00007770

	tmp := cube[2]
	tmp &= 0x00007770
	cube[1] &= 0x77770007
	cube[1] |= tmp
	
	tmp = cube[3]
	tmp &= 0x00007770
	cube[2] &= 0x77770007
	cube[2] |= tmp

	tmp = cube[4]
	tmp &= 0x00007770
	cube[3] &= 0x77770007
	cube[3] |= tmp

	cube[4] &= 0x77770007
	cube[4] |= swap
}

func spinDa(cube *[6]uint32) {
	spinFaceAnti(cube, 5)
	// spin edges
	swap := cube[1]
	swap &= 0x00007770

	tmp := cube[4]
	tmp &= 0x00007770
	cube[1] &= 0x77770007
	cube[1] |= tmp
	
	tmp = cube[3]
	tmp &= 0x00007770
	cube[4] &= 0x77770007
	cube[4] |= tmp

	tmp = cube[2]
	tmp &= 0x00007770
	cube[3] &= 0x77770007
	cube[3] |= tmp

	cube[2] &= 0x77770007
	cube[2] |= swap
}

func spinD2(cube *[6]uint32) {
	spinFace2(cube, 5)
	// spin edges
	swap := cube[1]
	swap &= 0x00007770

	tmp := cube[3]
	tmp &= 0x00007770
	cube[1] &= 0x77770007
	cube[1] |= tmp
	
	cube[3] &= 0x77770007
	cube[3] |= swap

	swap = cube[2]
	swap &= 0x00007770

	tmp = cube[4]
	tmp &= 0x00007770
	cube[2] &= 0x77770007
	cube[2] |= tmp
	
	cube[4] &= 0x77770007
	cube[4] |= swap
}

func spin(mix string, cube *[6]uint32) {
	sequence := strings.Fields(mix)
	fmt.Printf("\nsequence: %v, len: %d\n", sequence, len(sequence)) //
	for spin := 0; spin < len(sequence); spin++ {
		fmt.Printf("\nspin: %v\n", spin) //
		fmt.Printf("\nspin: %v\n", sequence[spin]) //
		if sequence[spin] == "U" {
			spinU(cube)
		} else if sequence[spin] == "U'" {
			spinUa(cube)
		} else if sequence[spin] == "U2" {
			spinU2(cube)
		} else if sequence[spin] == "D" {
			spinD(cube)
		} else if sequence[spin] == "D'" {
			spinDa(cube)
		} else if sequence[spin] == "D2" {
			spinD2(cube)
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
		dumpCube(cube)////
	}
	// test()//////!!!!!
}
