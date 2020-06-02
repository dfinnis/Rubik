package rubik

import (
	"fmt"
	"strings"
)

func spinFace(face *face) {
	tmpCorner := face.pieces[0][0]
	face.pieces[0][0] = face.pieces[2][0]
	face.pieces[2][0] = face.pieces[2][2]
	face.pieces[2][2] = face.pieces[0][2]
	face.pieces[0][2] = tmpCorner
	
	tmpMid := face.pieces[0][1]
	face.pieces[0][1] = face.pieces[1][0]
	face.pieces[1][0] = face.pieces[2][1]
	face.pieces[2][1] = face.pieces[1][2]
	face.pieces[1][2] = tmpMid
}

func spinFaceAnti(face *face) {
	tmpCorner := face.pieces[0][0]
	face.pieces[0][0] = face.pieces[0][2]
	face.pieces[0][2] = face.pieces[2][2]
	face.pieces[2][2] = face.pieces[2][0]
	face.pieces[2][0] = tmpCorner

	tmpMid := face.pieces[0][1]
	face.pieces[0][1] = face.pieces[1][2]
	face.pieces[1][2] = face.pieces[2][1]
	face.pieces[2][1] = face.pieces[1][0]
	face.pieces[1][0] = tmpMid
}

func spinF(cube *[6]face) {
	spinFace(&cube[2])
	// spin edges
	tmp0 := cube[0].pieces[2][0]
	tmp1 := cube[0].pieces[2][1]
	tmp2 := cube[0].pieces[2][2]

	cube[0].pieces[2][0] = cube[1].pieces[0][2]
	cube[0].pieces[2][1] = cube[1].pieces[1][2]
	cube[0].pieces[2][2] = cube[1].pieces[2][2]

	cube[1].pieces[0][2] = cube[5].pieces[0][0]
	cube[1].pieces[1][2] = cube[5].pieces[0][1]
	cube[1].pieces[2][2] = cube[5].pieces[0][2]

	cube[5].pieces[0][0] = cube[3].pieces[2][0]
	cube[5].pieces[0][1] = cube[3].pieces[1][0]
	cube[5].pieces[0][2] = cube[3].pieces[0][0]

	cube[3].pieces[0][0] = tmp0
	cube[3].pieces[1][0] = tmp1
	cube[3].pieces[2][0] = tmp2
}

func spinFa(cube *[6]face) {
	spinFaceAnti(&cube[2])
	// spin edges
	tmp0 := cube[0].pieces[2][0]
	tmp1 := cube[0].pieces[2][1]
	tmp2 := cube[0].pieces[2][2]

	cube[0].pieces[2][0] = cube[3].pieces[0][0]
	cube[0].pieces[2][1] = cube[3].pieces[1][0]
	cube[0].pieces[2][2] = cube[3].pieces[2][0]

	cube[3].pieces[0][0] = cube[5].pieces[0][2]
	cube[3].pieces[1][0] = cube[5].pieces[0][1]
	cube[3].pieces[2][0] = cube[5].pieces[0][0]

	cube[5].pieces[0][0] = cube[1].pieces[0][2]
	cube[5].pieces[0][1] = cube[1].pieces[1][2]
	cube[5].pieces[0][2] = cube[1].pieces[2][2]

	cube[1].pieces[0][2] = tmp2
	cube[1].pieces[1][2] = tmp1
	cube[1].pieces[2][2] = tmp0
}

func spinR(cube *[6]face) {
	spinFace(&cube[3])
	// spin edges
	tmp0 := cube[0].pieces[0][2]
	tmp1 := cube[0].pieces[1][2]
	tmp2 := cube[0].pieces[2][2]

	cube[0].pieces[0][2] = cube[2].pieces[0][2]
	cube[0].pieces[1][2] = cube[2].pieces[1][2]
	cube[0].pieces[2][2] = cube[2].pieces[2][2]

	cube[2].pieces[0][2] = cube[5].pieces[0][2]
	cube[2].pieces[1][2] = cube[5].pieces[1][2]
	cube[2].pieces[2][2] = cube[5].pieces[2][2]

	cube[5].pieces[0][2] = cube[4].pieces[2][0]
	cube[5].pieces[1][2] = cube[4].pieces[1][0]
	cube[5].pieces[2][2] = cube[4].pieces[0][0]

	cube[4].pieces[0][0] = tmp0
	cube[4].pieces[1][0] = tmp1
	cube[4].pieces[2][0] = tmp2
}

func spinRa(cube *[6]face) {
	spinFace(&cube[3])
	// spin edges
	tmp0 := cube[0].pieces[0][2]
	tmp1 := cube[0].pieces[1][2]
	tmp2 := cube[0].pieces[2][2]

	cube[0].pieces[0][2] = cube[4].pieces[2][0]
	cube[0].pieces[1][2] = cube[4].pieces[1][0]
	cube[0].pieces[2][2] = cube[4].pieces[0][0]

	cube[4].pieces[0][0] = cube[5].pieces[0][2]
	cube[4].pieces[1][0] = cube[5].pieces[1][2]
	cube[4].pieces[2][0] = cube[5].pieces[2][2]

	cube[5].pieces[0][2] = cube[2].pieces[0][2]
	cube[5].pieces[1][2] = cube[2].pieces[1][2]
	cube[5].pieces[2][2] = cube[2].pieces[2][2]

	cube[2].pieces[0][2] = tmp0
	cube[2].pieces[1][2] = tmp1
	cube[2].pieces[2][2] = tmp2
}

func spin(mix string, r *rubik) {
	// checkSpinError(mix)
	sequence := strings.Fields(mix)
	fmt.Printf("\nsequence: %v, len: %d\n", sequence, len(sequence)) //
	for spin := 0; spin < len(sequence); spin++ {
		// fmt.Printf("\nspin: %v\n", spin) //
		fmt.Printf("\nspin: %v\n", sequence[spin]) //
		if sequence[spin] == "U" {
			fmt.Printf("\nU!!!!\n") //
		} else if sequence[spin] == "U'" {
			fmt.Printf("\nU'\n") //
		} else if sequence[spin] == "U2" {
			fmt.Printf("\nU2\n") //
		} else if sequence[spin] == "D" {
			fmt.Printf("\nD\n") //
		} else if sequence[spin] == "D'" {
			fmt.Printf("\nD'\n") //
		} else if sequence[spin] == "D2" {
			fmt.Printf("\nD2\n") //
		} else if sequence[spin] == "R" {
			spinR(&r.cube)
		} else if sequence[spin] == "R'" {
			spinRa(&r.cube)
		} else if sequence[spin] == "R2" {
			spinR(&r.cube)
			spinR(&r.cube)
		} else if sequence[spin] == "L" {
			fmt.Printf("\nL\n") //
		} else if sequence[spin] == "L'" {
			fmt.Printf("\nL'\n") //
		} else if sequence[spin] == "L2" {
			fmt.Printf("\nL2\n") //
		} else if sequence[spin] == "F" {
			spinF(&r.cube)
		} else if sequence[spin] == "F'" {
			spinFa(&r.cube)
		} else if sequence[spin] == "F2" {
			spinF(&r.cube)
			spinF(&r.cube)
		} else if sequence[spin] == "B" {
			fmt.Printf("\nB\n") //
		} else if sequence[spin] == "B'" {
			fmt.Printf("\nB'\n") //
		} else if sequence[spin] == "B2" {
			fmt.Printf("\nB2\n") //
		} else {
			errorExit("bad input")
		}
		dumpCube(&r.cube)////
	}
	// fmt.Println(sequence, len(sequence)) //
}
