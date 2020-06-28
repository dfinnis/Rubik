package rubik

// import (
// 	"fmt"
// )

// const Reset		= "\x1B[0m"
// const White		= "\x1B[0m"					// 0 U
// const Orange	= "\x1B[38;2;255;165;0m"	// 1 L
// const Green		= "\x1B[32m"				// 2 F
// const Red		= "\x1B[31m"				// 3 R
// const Blue		= "\x1B[34m"				// 4 B
// const Yellow	= "\x1B[33m"				// 5 D

// // const WhiteBG	= "\x1B[0m"
// // const RedBG		= "\x1B[41m"
// // const GreenBG	= "\x1B[42m"
// // const YellowBG	= "\x1B[43m"
// // const BlueBG	= "\x1B[44m"
// // const OrangeBG	= "\x1B[48;2;255;165;0m"

// func dumpCube(cube *[6]face) {	
// 	fmt.Printf("\n\n#### -- CUBE -- ####\n")/////
// 	for y := 0; y < 3; y++ {
// 		fmt.Printf("\n        ")
// 		for x := 0; x < 3; x++ {
// 			if cube[0].pieces[y][x] == 0 {
// 				fmt.Printf("0 ")				
// 			} else if cube[0].pieces[y][x] == 1 {
// 				fmt.Printf("%v1%v ", Orange, Reset)
// 			} else if cube[0].pieces[y][x] == 2 {
// 				fmt.Printf("%v2%v ", Green, Reset)
// 			} else if cube[0].pieces[y][x] == 3 {
// 				fmt.Printf("%v3%v ", Red, Reset)
// 			} else if cube[0].pieces[y][x] == 4 {
// 				fmt.Printf("%v4%v ", Blue, Reset)
// 			} else if cube[0].pieces[y][x] == 5 {
// 				fmt.Printf("%v5%v ", Yellow, Reset)
// 			}
// 		}
// 	}
// 	// fmt.Printf("\n")/////
// 	for y := 0; y < 3; y++ {
// 		fmt.Printf("\n")
// 		for face := 1; face < 5; face++ {
// 			fmt.Printf(" ")
// 			for x := 0; x < 3; x++ {
// 				if cube[face].pieces[y][x] == 0 {
// 					fmt.Printf("0 ")
// 				} else if cube[face].pieces[y][x] == 1 {
// 					fmt.Printf("%v1%v ", Orange, Reset)
// 				} else if cube[face].pieces[y][x] == 2 {
// 					fmt.Printf("%v2%v ", Green, Reset)
// 				} else if cube[face].pieces[y][x] == 3 {
// 					fmt.Printf("%v3%v ", Red, Reset)
// 				} else if cube[face].pieces[y][x] == 4 {
// 					fmt.Printf("%v4%v ", Blue, Reset)
// 				} else if cube[face].pieces[y][x] == 5 {
// 					fmt.Printf("%v5%v ", Yellow, Reset)
// 				}
// 			}
// 		}
// 	}
// 	for y := 0; y < 3; y++ {
// 		fmt.Printf("\n        ")
// 		for x := 0; x < 3; x++ {
// 			if cube[5].pieces[y][x] == 0 {
// 				fmt.Printf("0 ")				
// 			} else if cube[5].pieces[y][x] == 1 {
// 				fmt.Printf("%v1%v ", Orange, Reset)
// 			} else if cube[5].pieces[y][x] == 2 {
// 				fmt.Printf("%v2%v ", Green, Reset)
// 			} else if cube[5].pieces[y][x] == 3 {
// 				fmt.Printf("%v3%v ", Red, Reset)
// 			} else if cube[5].pieces[y][x] == 4 {
// 				fmt.Printf("%v4%v ", Blue, Reset)
// 			} else if cube[5].pieces[y][x] == 5 {
// 				fmt.Printf("%v5%v ", Yellow, Reset)
// 			}
// 		}
// 	}
// 	fmt.Printf("\n")
// }
