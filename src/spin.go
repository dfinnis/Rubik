package rubik

import (
	"strings"
	// "fmt"
)

func modulo3(n int8) int8 {
	if n == -1 {
		return 2
	} else {
		return n % 3
	}
}

func spinU(cube *cepo) {
	// corner permutation
	tmp := cube.cP[0]
	cube.cP[0] = cube.cP[4]
	cube.cP[4] = cube.cP[3]
	cube.cP[3] = cube.cP[7]
	cube.cP[7] = tmp
	// edge permutation
	tmp = cube.eP[0]
	cube.eP[0] = cube.eP[8]
	cube.eP[8] = cube.eP[3]
	cube.eP[3] = cube.eP[11]
	cube.eP[11] = tmp
	// corner orientation
	cube.cO[cube.cP[0]] = modulo3(cube.cO[cube.cP[0]] - 1)
	cube.cO[cube.cP[3]] = modulo3(cube.cO[cube.cP[3]] - 1)
	cube.cO[cube.cP[4]] = modulo3(cube.cO[cube.cP[4]] + 1)
	cube.cO[cube.cP[7]] = modulo3(cube.cO[cube.cP[7]] + 1)
	// edge orientation
	cube.eO[cube.eP[0]] = (cube.eO[cube.eP[0]] + 1) % 2
	cube.eO[cube.eP[3]] = (cube.eO[cube.eP[3]] + 1) % 2
	cube.eO[cube.eP[8]] = (cube.eO[cube.eP[8]] + 1) % 2
	cube.eO[cube.eP[11]] = (cube.eO[cube.eP[11]] + 1) % 2
}

func spinD(cube *cepo) {
	// corner permutation
	tmp := cube.cP[1]
	cube.cP[1] = cube.cP[5]
	cube.cP[5] = cube.cP[2]
	cube.cP[2] = cube.cP[6]
	cube.cP[6] = tmp
	// edge permutation
	tmp = cube.eP[1]
	cube.eP[1] = cube.eP[10]
	cube.eP[10] = cube.eP[2]
	cube.eP[2] = cube.eP[9]
	cube.eP[9] = tmp
	// corner orientation
	cube.cO[cube.cP[1]] = modulo3(cube.cO[cube.cP[1]] - 1)
	cube.cO[cube.cP[2]] = modulo3(cube.cO[cube.cP[2]] - 1)
	cube.cO[cube.cP[5]] = modulo3(cube.cO[cube.cP[5]] + 1)
	cube.cO[cube.cP[6]] = modulo3(cube.cO[cube.cP[6]] + 1)
	// edge orientation
	cube.eO[cube.eP[1]] = (cube.eO[cube.eP[1]] + 1) % 2
	cube.eO[cube.eP[10]] = (cube.eO[cube.eP[10]] + 1) % 2
	cube.eO[cube.eP[2]] = (cube.eO[cube.eP[2]] + 1) % 2
	cube.eO[cube.eP[9]] = (cube.eO[cube.eP[9]] + 1) % 2
}

func spinF(cube *cepo) {
	// corner permutation
	tmp := cube.cP[4]
	cube.cP[4] = cube.cP[1]
	cube.cP[1] = cube.cP[6]
	cube.cP[6] = cube.cP[3]
	cube.cP[3] = tmp
	// edge permutation
	tmp = cube.eP[5]
	cube.eP[5] = cube.eP[9]
	cube.eP[9] = cube.eP[6]
	cube.eP[6] = cube.eP[8]
	cube.eP[8] = tmp
	// corner orientation
	cube.cO[cube.cP[1]] = modulo3(cube.cO[cube.cP[1]] + 1)
	cube.cO[cube.cP[3]] = modulo3(cube.cO[cube.cP[3]] + 1)
	cube.cO[cube.cP[4]] = modulo3(cube.cO[cube.cP[4]] - 1)
	cube.cO[cube.cP[6]] = modulo3(cube.cO[cube.cP[6]] - 1)
}

func spinB(cube *cepo) {
	// corner permutation
	tmp := cube.cP[7]
	cube.cP[7] = cube.cP[2]
	cube.cP[2] = cube.cP[5]
	cube.cP[5] = cube.cP[0]
	cube.cP[0] = tmp
	// edge permutation
	tmp = cube.eP[7]
	cube.eP[7] = cube.eP[10]
	cube.eP[10] = cube.eP[4]
	cube.eP[4] = cube.eP[11]
	cube.eP[11] = tmp
	// corner orientation
	cube.cO[cube.cP[0]] = modulo3(cube.cO[cube.cP[0]] + 1)
	cube.cO[cube.cP[2]] = modulo3(cube.cO[cube.cP[2]] + 1)
	cube.cO[cube.cP[5]] = modulo3(cube.cO[cube.cP[5]] - 1)
	cube.cO[cube.cP[7]] = modulo3(cube.cO[cube.cP[7]] - 1)
}

func spinL(cube *cepo) {
	// corner permutation
	tmp := cube.cP[0]
	cube.cP[0] = cube.cP[5]
	cube.cP[5] = cube.cP[1]
	cube.cP[1] = cube.cP[4]
	cube.cP[4] = tmp
	// edge permutation
	tmp = cube.eP[4]
	cube.eP[4] = cube.eP[1]
	cube.eP[1] = cube.eP[5]
	cube.eP[5] = cube.eP[0]
	cube.eP[0] = tmp
}

func spinR(cube *cepo) {
	// corner permutation
	tmp := cube.cP[3]
	cube.cP[3] = cube.cP[6]
	cube.cP[6] = cube.cP[2]
	cube.cP[2] = cube.cP[7]
	cube.cP[7] = tmp
	// edge permutation
	tmp = cube.eP[6]
	cube.eP[6] = cube.eP[2]
	cube.eP[2] = cube.eP[7]
	cube.eP[7] = cube.eP[3]
	cube.eP[3] = tmp
}

func spin(mix string, cube *cepo) {
	sequence := strings.Fields(mix)
	// fmt.Printf("\nsequence: %v, len: %d\n", sequence, len(sequence))	//	debug tool
	for spin := 0; spin < len(sequence); spin++ {
		// fmt.Printf("\nspin %v: %v\n", spin, sequence[spin])	//	debug tool
		if sequence[spin] == "U" {
			spinU(cube)
			cube.move = "U"
		} else if sequence[spin] == "U'" {
			spinU(cube)
			spinU(cube)
			spinU(cube)
			cube.move = "U'"
		} else if sequence[spin] == "U2" {
			spinU(cube)
			spinU(cube)
			cube.move = "U2"
		} else if sequence[spin] == "D" {
			spinD(cube)
			cube.move = "D"
		} else if sequence[spin] == "D'" {
			spinD(cube)
			spinD(cube)
			spinD(cube)
			cube.move = "D'"
		} else if sequence[spin] == "D2" {
			spinD(cube)
			spinD(cube)
			cube.move = "D2"
		} else if sequence[spin] == "R" {
			spinR(cube)
			cube.move = "R"
		} else if sequence[spin] == "R'" {
			spinR(cube)
			spinR(cube)
			spinR(cube)
			cube.move = "R'"
		} else if sequence[spin] == "R2" {
			spinR(cube)
			spinR(cube)
			cube.move = "R2"
		} else if sequence[spin] == "L" {
			spinL(cube)
			cube.move = "L"
		} else if sequence[spin] == "L'" {
			spinL(cube)
			spinL(cube)
			spinL(cube)
			cube.move = "L'"
		} else if sequence[spin] == "L2" {
			spinL(cube)
			spinL(cube)
			cube.move = "L2"
		} else if sequence[spin] == "F" {
			spinF(cube)
			cube.move = "F"
		} else if sequence[spin] == "F'" {
			spinF(cube)
			spinF(cube)
			spinF(cube)
			cube.move = "F'"
		} else if sequence[spin] == "F2" {
			spinF(cube)
			spinF(cube)
			cube.move = "F2"
		} else if sequence[spin] == "B" {
			spinB(cube)
			cube.move = "B"
		} else if sequence[spin] == "B'" {
			spinB(cube)
			spinB(cube)
			spinB(cube)
			cube.move = "B'"
		} else if sequence[spin] == "B2" {
			spinB(cube)
			spinB(cube)
			cube.move = "B2"
		} else {
			errorExit("bad input")
		}
		// dumpCube(cube)	//	debug tool
	}
}
