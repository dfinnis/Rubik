package rubik

// maxLength = 9999

// this is the framework for kociemba's alogrithm

// func Kociemba(position p) {
// 	for depth d from 0 to maxLength {
// 		Phase1search( p; d )
// 	}
// }

// func Phase1search( position p; depth d ) {
// 	if d=0 {
// 		if subgoal reached and last move was a quarter turn of R, L, F, or B {
// 			Phase2start( p )
// 		}
// 	}
// 	else if d>0 {
// 		if prune1[p]<=d {
// 			for each available move m {
// 				Phase1search( result of m applied to p; d-1 )
// 			}
// 		}
// 	}
// }

// func Phase2start( position p) {
// 	for depth d from 0 to maxLength - currentDepth {
//     	Phase2search( p; d )
// 	}
// }

// func Phase2search( position p; depth d ) {
// 	if d=0 {
//     	if solved {
// 			Found a solution!
// 			maxLength = currentDepth-1
// 		}
// 	}
// 	else if d>0 {
// 		if prune2[p]<=d {
// 			for each available move m {
// 				Phase2search( result of m applied to p; d-1 )
// 			}
// 		}
// 	}
// }

