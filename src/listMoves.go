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
			"L",
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
	// if subgroup != 0 {
	if cube.move != "" {
		dry := []string{}
		for _, move := range moves {
			// fmt.Printf("move: %v\n", move)//
			// fmt.Printf("move[0]: %v\n", move[0])//
			// fmt.Printf("cube.move: %v\n", cube.move)//
			// fmt.Printf("cube.move[0]: %v\n", cube.move[0])//
			// if move[0] == cube.move[0] {
			// 	fmt.Printf("remove!: %v\n", move)//
			// 	// moves = append(moves[:i], moves[i+1:]...)
			// 	moves = moves[:i+copy(moves[i:], moves[i+1:])]
			// 	fmt.Printf("moves!: %v\n", moves)//
			// }
			// if move[0] == cube.move[0] {
			// 	fmt.Printf("remove!: %v\n", move)//
			// 	// moves = append(moves[:i], moves[i+1:]...)
			// 	moves = moves[:i+copy(moves[i:], moves[i+1:])]
			// 	fmt.Printf("moves!: %v\n", moves)//
			// }
			if move[0] != cube.move[0] {
				// fmt.Printf("Add!: %v\n", move)//
				dry = append(dry, move)
				// fmt.Printf("moves!: %v\n", dry)//
			}
		}// remove opposite face move, not just last move?? i.e. avoid G0 R L R L????!!!
		moves = dry
	}
	// fmt.Printf("moves: %v\n", moves)//
	return moves
}
