package main

import (
	"fmt"
	"os"
	// "reflect"//
	"strings"
)

const Reset		= "\x1B[0m"
const White		= "\x1B[0m"					// 0 U
const Orange	= "\x1B[38;2;255;165;0m"	// 1 L
const Green		= "\x1B[32m"				// 2 F
const Red		= "\x1B[31m"				// 3 R
const Blue		= "\x1B[34m"				// 4 B
const Yellow	= "\x1B[33m"				// 5 D

// const WhiteBG	= "\x1B[0m"
// const RedBG		= "\x1B[41m"
// const GreenBG	= "\x1B[42m"
// const YellowBG	= "\x1B[43m"
// const BlueBG	= "\x1B[44m"
// const OrangeBG	= "\x1B[48;2;255;165;0m"

type face struct {
	pieces	[3][3]uint8
}

// rubik struct contains all information about current rubik state
type rubik struct {
	cube	[6]face
}

// rubik var r contains all information about current rubik state
var r *rubik

func errorExit(message string) {
	fmt.Printf("Error: %s\n", message)
	os.Exit(1)
}

func initRubik() *rubik {
	r = &rubik{}
	var face uint8
	for face = 0; face < 6; face++ {
		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				r.cube[face].pieces[y][x] = face
			}
		}
	}
	return r
}

// func dumpCube(cube *[6]face) {
// 	for face := 0; face < 6; face++ {
// 		fmt.Printf("Face: %d\n", face)
// 		// fmt.Printf("%d\n", cube[face])
// 		for y := 0; y < 3; y++ {
// 			for x := 0; x < 3; x++ {
// 				// fmt.Printf("%d ", cube[face].pieces[y][x])
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
// 			fmt.Printf("\n")/////
// 		}
// 		fmt.Printf("\n")/////
// 	}
// }

func dumpCube2(cube *[6]face) {	
	fmt.Printf("\n\n#### -- CUBE -- ####\n")/////
	for y := 0; y < 3; y++ {
		fmt.Printf("\n        ")
		for x := 0; x < 3; x++ {
			if cube[0].pieces[y][x] == 0 {
				fmt.Printf("0 ")				
			} else if cube[0].pieces[y][x] == 1 {
				fmt.Printf("%v1%v ", Orange, Reset)
			} else if cube[0].pieces[y][x] == 2 {
				fmt.Printf("%v2%v ", Green, Reset)
			} else if cube[0].pieces[y][x] == 3 {
				fmt.Printf("%v3%v ", Red, Reset)
			} else if cube[0].pieces[y][x] == 4 {
				fmt.Printf("%v4%v ", Blue, Reset)
			} else if cube[0].pieces[y][x] == 5 {
				fmt.Printf("%v5%v ", Yellow, Reset)
			}
		}
	}
	// fmt.Printf("\n")/////
	for y := 0; y < 3; y++ {
		fmt.Printf("\n")
		for face := 1; face < 5; face++ {
			fmt.Printf(" ")
			for x := 0; x < 3; x++ {
				if cube[face].pieces[y][x] == 0 {
					fmt.Printf("0 ")
				} else if cube[face].pieces[y][x] == 1 {
					fmt.Printf("%v1%v ", Orange, Reset)
				} else if cube[face].pieces[y][x] == 2 {
					fmt.Printf("%v2%v ", Green, Reset)
				} else if cube[face].pieces[y][x] == 3 {
					fmt.Printf("%v3%v ", Red, Reset)
				} else if cube[face].pieces[y][x] == 4 {
					fmt.Printf("%v4%v ", Blue, Reset)
				} else if cube[face].pieces[y][x] == 5 {
					fmt.Printf("%v5%v ", Yellow, Reset)
				}
			}
		}
	}
	for y := 0; y < 3; y++ {
		fmt.Printf("\n        ")
		for x := 0; x < 3; x++ {
			if cube[5].pieces[y][x] == 0 {
				fmt.Printf("0 ")				
			} else if cube[5].pieces[y][x] == 1 {
				fmt.Printf("%v1%v ", Orange, Reset)
			} else if cube[5].pieces[y][x] == 2 {
				fmt.Printf("%v2%v ", Green, Reset)
			} else if cube[5].pieces[y][x] == 3 {
				fmt.Printf("%v3%v ", Red, Reset)
			} else if cube[5].pieces[y][x] == 4 {
				fmt.Printf("%v4%v ", Blue, Reset)
			} else if cube[5].pieces[y][x] == 5 {
				fmt.Printf("%v5%v ", Yellow, Reset)
			}
		}
	}
	fmt.Printf("\n")
}

func parseArg() string {
	args := os.Args[1:]
	if len(args) == 0 {
		errorExit("not enough arguments, no mix given")
	} else if len(args) > 1 {
		errorExit("too many arguments")
	}
	mix := args[0]
	// fmt.Println(reflect.TypeOf(moveList))
	return mix
}

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

	cube[5].pieces[0][0] = cube[3].pieces[0][0]
	cube[5].pieces[0][1] = cube[3].pieces[1][0]
	cube[5].pieces[0][2] = cube[3].pieces[2][0]

	cube[3].pieces[0][0] = tmp0
	cube[3].pieces[1][0] = tmp1
	cube[3].pieces[2][0] = tmp2
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
			fmt.Printf("\nR!!!!\n") //
		} else if sequence[spin] == "R'" {
			fmt.Printf("\nR'\n") //
		} else if sequence[spin] == "R2" {
			fmt.Printf("\nR2\n") //
		} else if sequence[spin] == "L" {
			fmt.Printf("\nL\n") //
		} else if sequence[spin] == "L'" {
			fmt.Printf("\nL'\n") //
		} else if sequence[spin] == "L2" {
			fmt.Printf("\nL2\n") //
		} else if sequence[spin] == "F" {
			fmt.Printf("\nFFF!!!!\n") //
			spinF(&r.cube)
		} else if sequence[spin] == "F'" {
			fmt.Printf("\nF'\n") //
		} else if sequence[spin] == "F2" {
			fmt.Printf("\nF2\n") //
		} else if sequence[spin] == "B" {
			fmt.Printf("\nB\n") //
		} else if sequence[spin] == "B'" {
			fmt.Printf("\nB'\n") //
		} else if sequence[spin] == "B2" {
			fmt.Printf("\nB2\n") //
		} else {
			errorExit("bad input")
		}
	}
	// fmt.Println(sequence, len(sequence)) //
}

// func solve(&r.cube) string {
	
	// return solution
// }

// func printSolution(solution) {

// }


func main() {
	fmt.Printf("oh hi!\n")/////////
	r := initRubik()
	parseArg()
	// dumpCube(&r.cube)
	mix := parseArg()
	fmt.Printf("mix: %s\n", mix)
	dumpCube2(&r.cube)////
	spin(mix, r)
	// solution := solve(&r.cube)
	// printSolution(solution)
	// rubik.runGraphic()
	dumpCube2(&r.cube)////
	fmt.Printf("\nEND!!\n")//////////
}

// ## To run enter either command:
// go run main.go "$(< mix/subject.txt)"
// go build; ./Rubik "$(< mix/subject.txt)"