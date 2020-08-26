package rubik

// import (
// 	"fmt"//
// )

func listMovesCepo(cube *cepo, subgroup int8) []string {
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
				dry = append(dry, move)
			}
		}// remove opposite face move, not just last move?? i.e. avoid G0 R L R L????!!!
		moves = dry
	}
	// fmt.Printf("moves: %v\n", moves)//
	return moves
}
