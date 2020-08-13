package rubik

import (
	"fmt"
	"os"
)

func recursive(r *rubik, subgroup uint8) {
	// heuristi := heuristic(&r.cube, subgroup)
	// if heuristi == 0 {
	// 	fmt.Printf("END!!!!!!!!")
	// 	os.Exit(1)
	// }
	moves := listMoves(r, subgroup)
	fmt.Printf("moves: %v\n", moves)//

	// var path []rubik
	// path = append(path, *r)

	// for heuristic(&new.cube, subgroup) != 0 {
	for i:= 0; i < len(moves); i++ {
		new := newNode(&r.cube, moves[i])
		spin(moves[i], &new.cube)
		fmt.Printf("moves[i]: %v\n", moves[i])///
		heuristi := heuristic(&new.cube, subgroup)
		if heuristi == 0 {
			fmt.Printf("END!!!!!!!!")
			os.Exit(1)
		}
		// recursive(new, subgroup)
		// heuristi = heuristic(&new.cube, subgroup)
		// if heuristi == 0 {
		// 	fmt.Printf("NEW GROUP!!\n")///
		// 	fmt.Printf("moves[i]: %v\n", moves[i])///
		// 	dumpCube(&new.cube)
		// 	break
		// }
	}
}

func bruteForce(r *rubik) string {
	subgroup := subgroup(&r.cube)
	fmt.Printf("\nsubgroup: %v\n", subgroup)//
	recursive(r, subgroup)
	// moves := listMoves(r, subgroup)
	// fmt.Printf("moves: %v\n", moves)//

	// // var path []rubik
	// // path = append(path, *r)

	// // for heuristic(&new.cube, subgroup) != 0 {
	// heuristi := heuristic(&r.cube, subgroup)
	// for {
	// 	for i:= 0; i < len(moves); i++ {
	// 		new := newNode(&r.cube, moves[i])
	// 		spin(moves[i], &new.cube)
	// 		fmt.Printf("moves[i]: %v\n", moves[i])///

	// 		heuristi = heuristic(&new.cube, subgroup)
	// 		if heuristi == 0 {
	// 			fmt.Printf("NEW GROUP!!\n")///
	// 			fmt.Printf("moves[i]: %v\n", moves[i])///
	// 			dumpCube(&new.cube)
	// 			break
	// 		}
	// 	}
	// 	if heuristi == 0 {
	// 		break
	// 	}
	// }
	return "F U"
}