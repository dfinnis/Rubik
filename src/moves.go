package rubik

// import (
// 	"fmt"//
// )

func listMoves(cube *cepo, subgroup int8) []string {
	moves := []string{}
	if subgroup == 0 {
		moves = []string{
			"U",
			"U'",
			"U2",
			"D",
			"D'",
			"D2",
			"R",
			"R'",
			"R2",
			"L",
			"L'",
			"L2",
			"F",
			"F'",
			"F2",
			"B",
			"B'",
			"B2",
		}
	} else if subgroup == 1 {
		moves = []string{
			"U2",
			"D2",
			"R",
			"R'",
			"R2",
			"L",
			"L'",
			"L2",
			"F",
			"F'",
			"F2",
			"B",
			"B'",
			"B2",
		}
	} else if subgroup == 2 {
		moves = []string{
			"U2",
			"D2",
			"R",
			"R'",
			"R2",
			"L",
			"L'",
			"L2",
			"F2",
			"B2",
		}
	} else { // subgroup = 3
		moves = []string{
			"U2",
			"D2",
			"R2",
			"L2",
			"F2",
			"B2",
		}
	}
	if cube.move != "" {
		dry := []string{}
		for _, move := range moves {
			if move[0] != cube.move[0] {
				if cube.move2 != "" {
					if move[0] != cube.move2[0] {
						dry = append(dry, move)
					}
				} else {
					dry = append(dry, move)
				}
			}
		}
		moves = dry
	}
	// fmt.Printf("moves: %v\n", moves)//
	return moves
}

func listAllMoves(cube *cepo) []string {
	moves := []string{
			"U",
			"U'",
			"U2",
			"D",
			"D'",
			"D2",
			"R",
			"R'",
			"R2",
			"L",
			"L'",
			"L2",
			"F",
			"F'",
			"F2",
			"B",
			"B'",
			"B2",
	}
	// dry := moves
	// // // for i := 0; i < n; i++ {
	// // // 	move := dry[rand.Intn(len(dry))]
	// // // 	mix += move
	// if cube.move == "U" || cube.move == "U'" || cube.move == "U2" {
	// 	if cube.move2 == "D" || cube.move2 == "D'" || cube.move2 == "D2" {
	// 		dry = moves[3:]
	// 	} else {
	// 		dry = moves[6:]
	// 	}
// 	} else if move == "D" || move == "D'" || move == "D2" {
// 		if stringInSlice("U", dry) {
// 			dry = append([]string{}, spin[:3]...)
// 			dry = append(dry, spin[6:]...)
// 		} else {
// 			dry = spin[6:]
// 		}
// 	} else if move == "R" || move == "R'" || move == "R2" {
// 		dry = append([]string{}, spin[:6]...)
// 		if stringInSlice("L", dry) {
// 			dry = append(dry, spin[9:]...)
// 		} else {
// 			dry = append(dry, spin[12:]...)
// 		}
// 	} else if move == "L" || move == "L'" || move == "L2" {
// 		if stringInSlice("R", dry) {
// 			dry = append([]string{}, spin[:9]...)
// 		} else {
// 			dry = append([]string{}, spin[:6]...)
// 		}
// 		dry = append(dry, spin[12:]...)
// 	} else if move == "F" || move == "F'" || move == "F2" {
// 		dry = append([]string{}, spin[:12]...)
// 		if stringInSlice("B", dry) {
// 			dry = append(dry, spin[15:]...)
// 		}
// 	} else if move == "B" || move == "B'" || move == "B2" {
// 		if stringInSlice("F", dry) {
// 			dry = spin[:15]
// 		} else {
// 			dry = spin[:12]
	// }
	return moves
}