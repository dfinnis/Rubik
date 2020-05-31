package main

import (
	"fmt"
)

const Reset		= "\x1B[0m"
const White		= "\x1B[0m"
const Red		= "\x1B[31m"
const Green		= "\x1B[32m"
const Yellow	= "\x1B[33m"
const Blue		= "\x1B[34m"
const Orange	= "\x1B[38;2;255;165;0m"

// const WhiteBG	= "\x1B[0m"
// const RedBG		= "\x1B[31m"
// const GreenBG	= "\x1B[32m"
// const YellowBG	= "\x1B[33m"
// const BlueBG	= "\x1B[34m"
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
	// fmt.Printf(Orange)//////
	for face := 0; face < 6; face++ {
		fmt.Printf("Face: %d\n", face)
		// fmt.Printf("%d\n", cube[face])
		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				// fmt.Printf("%d ", cube[face].pieces[y][x])
				if cube[face].pieces[y][x] == 0 {
					fmt.Printf("1 ")				
				} else if cube[face].pieces[y][x] == 1 {
					fmt.Printf("%v1%v ", Red, Reset)
				} else if cube[face].pieces[y][x] == 2 {
					fmt.Printf("%v2%v ", Green, Reset)
				} else if cube[face].pieces[y][x] == 3 {
					fmt.Printf("%v3%v ", Yellow, Reset)
				} else if cube[face].pieces[y][x] == 4 {
					fmt.Printf("%v4%v ", Blue, Reset)
				} else if cube[face].pieces[y][x] == 5 {
					fmt.Printf("%v5%v ", Orange, Reset)
				}
			}
			fmt.Printf("\n")/////
		}
		fmt.Printf("\n")/////
	}
}

func main() {
	fmt.Printf("oh hi!\n")/////////
	r := initRubik()
	dumpCube(&r.cube)
	fmt.Printf("\nEND!!\n")//////////
}

// ## To run enter either command:
// go run main.go
// go build; ./Rubik