package rubik

import (
	"fmt" //
	// "math/bits"
	"strings"
)

// func test() {//////
// 	fmt.Printf("test!\n")/////////
// 	var face uint32
// 	face = 1
// 	fmt.Printf("\nint before: %v\n", face)/////////
// 	face = bits.RotateLeft32(face, 2)
// 	fmt.Printf("\nint after: %v\n\n", face)/////////
// 	face = bits.RotateLeft32(face, -3)
// 	fmt.Printf("\nint after: %v\n\n", face)/////////
// 	fmt.Printf("test end!\n")/////////
// }

// func spinFace(face *face) {
// 	tmpCorner := face.pieces[0][0]
// 	face.pieces[0][0] = face.pieces[2][0]
// 	face.pieces[2][0] = face.pieces[2][2]
// 	face.pieces[2][2] = face.pieces[0][2]
// 	face.pieces[0][2] = tmpCorner
	
// 	tmpMid := face.pieces[0][1]
// 	face.pieces[0][1] = face.pieces[1][0]
// 	face.pieces[1][0] = face.pieces[2][1]
// 	face.pieces[2][1] = face.pieces[1][2]
// 	face.pieces[1][2] = tmpMid
// }

// func spinFaceAnti(face *face) {
// 	tmpCorner := face.pieces[0][0]
// 	face.pieces[0][0] = face.pieces[0][2]
// 	face.pieces[0][2] = face.pieces[2][2]
// 	face.pieces[2][2] = face.pieces[2][0]
// 	face.pieces[2][0] = tmpCorner

// 	tmpMid := face.pieces[0][1]
// 	face.pieces[0][1] = face.pieces[1][2]
// 	face.pieces[1][2] = face.pieces[2][1]
// 	face.pieces[2][1] = face.pieces[1][0]
// 	face.pieces[1][0] = tmpMid
// }

func spinU(cube *[6]uint32) {
	// spinFace(&cube[0])
	// spin edges
	tmp0 := cube[1]
	fmt.Printf("\ntmp0			:= %032b\n", tmp0)
	//// a & 196	query a value for its set bits
	// &=		selectively clearing bits of an integer value to zero
	// |=		set arbitrary bits for a given integer value

	fmt.Printf("cube[1] before	:= %032b\n", cube[1])
	// &=		selectively clear bits of an int to zero
	cube[1] &= 0x77777
	fmt.Printf("cube[1] after	:= %032b\n\n", cube[1])
	tmp1 := cube[2]
	fmt.Printf("tmp1 			:= %032b\n\n", tmp1)
	tmp1 &= 0x77700000
	fmt.Printf("tmp1			:= %032b\n\n", tmp1)
	cube[1] |= tmp1	// |=		set arbitrary bits for a given integer value
	fmt.Printf("cube[1] again	:= %032b\n\n", cube[1])

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
