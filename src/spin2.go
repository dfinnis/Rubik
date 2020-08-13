package rubik

import (
	"strings"
	// "fmt"
)

// func modulo2(n int8) int8 {
// 	if n == -1 {
// 		return 1
// 	} else {
// 		return n % 2
// 	}
// }

func modulo3(n int8) int8 {
	if n == -1 {
		return 2
	} else {
		return n % 3
	}
}

func spinCepoU(cube *cepo) {
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

func spinCepoD(cube *cepo) {
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

func spinCepoF(cube *cepo) {
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

func spinCepoB(cube *cepo) {
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

func spinCepoL(cube *cepo) {
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

func spinCepoR(cube *cepo) {
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

func spinCepo(mix string, cube *cepo) {
	sequence := strings.Fields(mix)
	// fmt.Printf("\nsequence: %v, len: %d\n", sequence, len(sequence))	//	debug tool
	for spin := 0; spin < len(sequence); spin++ {
		// fmt.Printf("\nspin %v: %v\n", spin, sequence[spin])	//	debug tool
		if sequence[spin] == "U" {
			spinCepoU(cube)
		} else if sequence[spin] == "U'" {
			spinCepoU(cube)
			spinCepoU(cube)
			spinCepoU(cube)
		} else if sequence[spin] == "U2" {
			spinCepoU(cube)
			spinCepoU(cube)
		} else if sequence[spin] == "D" {
			spinCepoD(cube)
		} else if sequence[spin] == "D'" {
			spinCepoD(cube)
			spinCepoD(cube)
			spinCepoD(cube)
		} else if sequence[spin] == "D2" {
			spinCepoD(cube)
			spinCepoD(cube)
		} else if sequence[spin] == "R" {
			spinCepoR(cube)
		} else if sequence[spin] == "R'" {
			spinCepoR(cube)
			spinCepoR(cube)
			spinCepoR(cube)
		} else if sequence[spin] == "R2" {
			spinCepoR(cube)
			spinCepoR(cube)
		} else if sequence[spin] == "L" {
			spinCepoL(cube)
		} else if sequence[spin] == "L'" {
			spinCepoL(cube)
			spinCepoL(cube)
			spinCepoL(cube)
		} else if sequence[spin] == "L2" {
			spinCepoL(cube)
			spinCepoL(cube)
		} else if sequence[spin] == "F" {
			spinCepoF(cube)
		} else if sequence[spin] == "F'" {
			spinCepoF(cube)
			spinCepoF(cube)
			spinCepoF(cube)
		} else if sequence[spin] == "F2" {
			spinCepoF(cube)
			spinCepoF(cube)
		} else if sequence[spin] == "B" {
			spinCepoB(cube)
		} else if sequence[spin] == "B'" {
			spinCepoB(cube)
			spinCepoB(cube)
			spinCepoB(cube)
		} else if sequence[spin] == "B2" {
			spinCepoB(cube)
			spinCepoB(cube)
		} else {
			errorExit("bad input")
		}
		// dumpCube(cube)	//	debug tool
	}
}
