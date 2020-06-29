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


func spinR(cube *[6]uint32) {
	spinFace(cube, 3)
	// spin edges
	swap := cube[2]
	swap &= 0x00777000

	tmp := cube[5]
	tmp &= 0x00777000
	cube[2] &= 0x77000777
	cube[2] |= tmp
	
	tmp0 := cube[4]
	// tmp1 := cube[4]
	// tmp2 := cube[4]	
	tmp0 &= 0x00000070
	fmt.Printf("tmp0: %032b\n", tmp0)//!!!!
	tmp0 &= 0x00000070/////////////rotate
	fmt.Printf("tmp0: %032b\n", tmp0)//!!!!
	cube[5] &= 0x77000777
	fmt.Printf("cube[5]: %032b\n", cube[5])//!!!!

	// cube[2] |= tmp

	// tmp = cube[4]
	// tmp &= 0x00007770
	// cube[3] &= 0x77770007
	// cube[3] |= tmp

	// cube[4] &= 0x77770007
	// cube[4] |= swap
}

func spinRa(cube *[6]uint32) {
	spinFaceAnti(cube, 3)
	// spin edges
	// swap := cube[1]
	// swap &= 0x00007770

	// tmp := cube[4]
	// tmp &= 0x00007770
	// cube[1] &= 0x77770007
	// cube[1] |= tmp
	
	// tmp = cube[3]
	// tmp &= 0x00007770
	// cube[4] &= 0x77770007
	// cube[4] |= tmp

	// tmp = cube[2]
	// tmp &= 0x00007770
	// cube[3] &= 0x77770007
	// cube[3] |= tmp

	// cube[2] &= 0x77770007
	// cube[2] |= swap
}

func spinR2(cube *[6]uint32) {
	spinFace2(cube, 3)
	// spin edges
	// swap := cube[1]
	// swap &= 0x00007770

	// tmp := cube[3]
	// tmp &= 0x00007770
	// cube[1] &= 0x77770007
	// cube[1] |= tmp
	
	// cube[3] &= 0x77770007
	// cube[3] |= swap

	// swap = cube[2]
	// swap &= 0x00007770

	// tmp = cube[4]
	// tmp &= 0x00007770
	// cube[2] &= 0x77770007
	// cube[2] |= tmp
	
	// cube[4] &= 0x77770007
	// cube[4] |= swap
}

func spin(mix string, cube *[6]uint32) {
	sequence := strings.Fields(mix)
	fmt.Printf("\nsequence: %v, len: %d\n", sequence, len(sequence)) //
	for spin := 0; spin < len(sequence); spin++ {
		fmt.Printf("\nspin %v: %v\n", spin, sequence[spin]) //
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
		} else if sequence[spin] == "R" {
			spinR(&r.cube)
		} else if sequence[spin] == "R'" {
			spinRa(&r.cube)
		} else if sequence[spin] == "R2" {
			spinR2(&r.cube)
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
