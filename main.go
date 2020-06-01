package main

import (
	"fmt"
	"os"
	// "reflect"//
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

func dumpCube(cube *[6]face) {
	for face := 0; face < 6; face++ {
		fmt.Printf("Face: %d\n", face)
		// fmt.Printf("%d\n", cube[face])
		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				// fmt.Printf("%d ", cube[face].pieces[y][x])
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
			fmt.Printf("\n")/////
		}
		fmt.Printf("\n")/////
	}
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

// func runMix(mix) {

// }

// func solve(&r.cube) string {
	
	// return solution
// }

// func printSolution(solution) {

// }

func main() {
	fmt.Printf("oh hi!\n")/////////
	r := initRubik()
	parseArg()
	dumpCube(&r.cube)
	mix := parseArg()
	fmt.Printf("mix: %s", mix)
	// runMix(mix)
	// solution := solve(&r.cube)
	// printSolution(solution)
	fmt.Printf("\nEND!!\n")//////////
}

// ## To run enter either command:
// go run main.go
// go build; ./Rubik