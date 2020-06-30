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

func spinDa(cube *[6]uint32) {
	spinFaceAnti(cube, 5)
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
	
	tmp = cube[4]
	tmp &= 0x70000077
	tmp = bits.RotateLeft32(tmp, 16)
	cube[5] &= 0x77000777
	cube[5] |= tmp

	tmp = cube[0]
	tmp &= 0x00777000
	tmp = bits.RotateLeft32(tmp, 16)
	cube[4] &= 0x07777700
	cube[4] |= tmp

	cube[0] &= 0x77000777
	cube[0] |= swap
}

func spinRa(cube *[6]uint32) {
	spinFaceAnti(cube, 3)
	// spin edges
	swap := cube[2]
	swap &= 0x00777000

	tmp := cube[0]
	tmp &= 0x00777000
	cube[2] &= 0x77000777
	cube[2] |= tmp
	
	tmp = cube[4]
	tmp &= 0x70000077
	tmp = bits.RotateLeft32(tmp, 16)
	cube[0] &= 0x77000777
	cube[0] |= tmp

	tmp = cube[5]
	tmp &= 0x00777000
	tmp = bits.RotateLeft32(tmp, 16)
	cube[4] &= 0x07777700
	cube[4] |= tmp

	cube[5] &= 0x77000777
	cube[5] |= swap
}

func spinR2(cube *[6]uint32) {
	spinFace2(cube, 3)
	// spin edges
	swap := cube[2]
	swap &= 0x00777000
	swap = bits.RotateLeft32(swap, 16)

	tmp := cube[4]
	tmp &= 0x70000077
	tmp = bits.RotateLeft32(tmp, 16)
	cube[2] &= 0x77000777
	cube[2] |= tmp
	
	cube[4] &= 0x07777700
	cube[4] |= swap

	swap = cube[0]
	swap &= 0x00777000

	tmp = cube[5]
	tmp &= 0x00777000
	cube[0] &= 0x77000777
	cube[0] |= tmp
	
	cube[5] &= 0x77000777
	cube[5] |= swap
}

func spinL(cube *[6]uint32) {
	spinFace(cube, 1)
	// spin edges
	swap := cube[2]
	swap &= 0x70000077

	tmp := cube[0]
	tmp &= 0x70000077
	cube[2] &= 0x07777700
	cube[2] |= tmp

	tmp = cube[4]
	tmp &= 0x00777000
	tmp = bits.RotateLeft32(tmp, 16)
	cube[0] &= 0x07777700
	cube[0] |= tmp

	tmp = cube[5]
	tmp &= 0x70000077
	tmp = bits.RotateLeft32(tmp, 16)
	cube[4] &= 0x77000777
	cube[4] |= tmp

	cube[5] &= 0x07777700
	cube[5] |= swap
}

func spinLa(cube *[6]uint32) {
	spinFaceAnti(cube, 1)
	// spin edges
	swap := cube[2]
	swap &= 0x70000077

	tmp := cube[5]
	tmp &= 0x70000077
	cube[2] &= 0x07777700
	cube[2] |= tmp

	tmp = cube[4]
	tmp &= 0x00777000
	tmp = bits.RotateLeft32(tmp, 16)
	cube[5] &= 0x07777700
	cube[5] |= tmp

	tmp = cube[0]
	tmp &= 0x70000077
	tmp = bits.RotateLeft32(tmp, 16)
	cube[4] &= 0x77000777
	cube[4] |= tmp

	cube[0] &= 0x07777700
	cube[0] |= swap
}

func spinL2(cube *[6]uint32) {
	spinFace2(cube, 1)
	// spin edges
	swap := cube[2]
	swap &= 0x70000077
	swap = bits.RotateLeft32(swap, 16)

	tmp := cube[4]
	tmp &= 0x00777000
	tmp = bits.RotateLeft32(tmp, 16)
	cube[2] &= 0x07777700
	cube[2] |= tmp

	cube[4] &= 0x77000777
	cube[4] |= swap

	swap = cube[0]
	swap &= 0x70000077

	tmp = cube[5]
	tmp &= 0x70000077
	cube[0] &= 0x07777700
	cube[0] |= tmp

	cube[5] &= 0x07777700
	cube[5] |= swap
}

func spinF(cube *[6]uint32) {
	spinFace(cube, 2)
	// spin edges
	swap := cube[1]
	swap &= 0x00777000
	swap = bits.RotateLeft32(swap, -8)

	tmp := cube[5]
	tmp &= 0x77700000
	tmp = bits.RotateLeft32(tmp, -8)
	cube[1] &= 0x77000777
	cube[1] |= tmp

	tmp = cube[3]
	tmp &= 0x70000077
	tmp = bits.RotateLeft32(tmp, -8)
	cube[5] &= 0x00077777
	cube[5] |= tmp

	tmp = cube[0]
	tmp &= 0x00007770
	tmp = bits.RotateLeft32(tmp, -8)
	cube[3] &= 0x07777700
	cube[3] |= tmp

	cube[0] &= 0x77770007
	cube[0] |= swap
}

func spinFa(cube *[6]uint32) {
	spinFaceAnti(cube, 2)
	// spin edges
	swap := cube[1]
	swap &= 0x00777000
	swap = bits.RotateLeft32(swap, 8)

	tmp := cube[0]
	tmp &= 0x0007770
	tmp = bits.RotateLeft32(tmp, 8)
	cube[1] &= 0x77000777
	cube[1] |= tmp

	tmp = cube[3]
	tmp &= 0x70000077
	tmp = bits.RotateLeft32(tmp, 8)
	cube[0] &= 0x77770007
	cube[0] |= tmp

	tmp = cube[5]
	tmp &= 0x77700000
	tmp = bits.RotateLeft32(tmp, 8)
	cube[3] &= 0x07777700
	cube[3] |= tmp

	cube[5] &= 0x00077777
	cube[5] |= swap
}

func spinF2(cube *[6]uint32) {
	spinFace2(cube, 2)
	// spin edges
	swap := cube[1]
	swap &= 0x00777000
	swap = bits.RotateLeft32(swap, 16)

	tmp := cube[3]
	tmp &= 0x70000077
	tmp = bits.RotateLeft32(tmp, 16)
	cube[1] &= 0x77000777
	cube[1] |= tmp

	cube[3] &= 0x07777700
	cube[3] |= swap

	swap = cube[0]
	swap &= 0x00007770
	swap = bits.RotateLeft32(swap, 16)

	tmp = cube[5]
	tmp &= 0x77700000
	tmp = bits.RotateLeft32(tmp, 16)
	cube[0] &= 0x77770007
	cube[0] |= tmp

	cube[5] &= 0x00077777
	cube[5] |= swap
}

func spinB(cube *[6]uint32) {
	spinFace(cube, 4)
	// spin edges
	swap := cube[3]
	swap &= 0x00777000
	swap = bits.RotateLeft32(swap, 8)

	tmp := cube[5]
	tmp &= 0x00007770
	tmp = bits.RotateLeft32(tmp, 8)
	cube[3] &= 0x77000777
	cube[3] |= tmp

	tmp = cube[1]
	tmp &= 0x70000077
	tmp = bits.RotateLeft32(tmp, 8)
	cube[5] &= 0x77770007
	cube[5] |= tmp

	tmp = cube[0]
	tmp &= 0x77700000
	tmp = bits.RotateLeft32(tmp, 8)
	cube[1] &= 0x07777700
	cube[1] |= tmp

	cube[0] &= 0x00077777
	cube[0] |= swap
}

func spinBa(cube *[6]uint32) {
	spinFaceAnti(cube, 4)
	// spin edges
	swap := cube[3]
	swap &= 0x00777000
	swap = bits.RotateLeft32(swap, -8)

	tmp := cube[0]
	tmp &= 0x77700000
	tmp = bits.RotateLeft32(tmp, -8)
	cube[3] &= 0x77000777
	cube[3] |= tmp

	tmp = cube[1]
	tmp &= 0x70000077
	tmp = bits.RotateLeft32(tmp, -8)
	cube[0] &= 0x00077777
	cube[0] |= tmp

	tmp = cube[5]
	tmp &= 0x00007770
	tmp = bits.RotateLeft32(tmp, -8)
	cube[1] &= 0x07777700
	cube[1] |= tmp

	cube[5] &= 0x77770007
	cube[5] |= swap
}

func spinB2(cube *[6]uint32) {
	spinFace2(cube, 4)
	// spin edges
	swap := cube[3]
	swap &= 0x00777000
	swap = bits.RotateLeft32(swap, 16)

	tmp := cube[1]
	tmp &= 0x70000077
	tmp = bits.RotateLeft32(tmp, 16)
	cube[3] &= 0x77000777
	cube[3] |= tmp

	cube[1] &= 0x07777700
	cube[1] |= swap

	swap = cube[0]
	swap &= 0x77700000
	swap = bits.RotateLeft32(swap, 16)

	tmp = cube[5]
	tmp &= 0x00007770
	tmp = bits.RotateLeft32(tmp, 16)
	cube[0] &= 0x00077777
	cube[0] |= tmp

	cube[5] &= 0x77770007
	cube[5] |= swap
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
			spinR(cube)
		} else if sequence[spin] == "R'" {
			spinRa(cube)
		} else if sequence[spin] == "R2" {
			spinR2(cube)
		} else if sequence[spin] == "L" {
			spinL(cube)
		} else if sequence[spin] == "L'" {
			spinLa(cube)
		} else if sequence[spin] == "L2" {
			spinL2(cube)
		} else if sequence[spin] == "F" {
			spinF(cube)
		} else if sequence[spin] == "F'" {
			spinFa(cube)
		} else if sequence[spin] == "F2" {
			spinF2(cube)
		} else if sequence[spin] == "B" {
			spinB(cube)
		} else if sequence[spin] == "B'" {
			spinBa(cube)
		} else if sequence[spin] == "B2" {
			spinB2(cube)
		} else {
			errorExit("bad input")
		}
		dumpCube(cube)////
	}
}
