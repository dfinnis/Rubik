package main

import (
	"fmt"
)

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
	return r
}

func dumpRubik(rubik *[6]face) {
	for face := 0; face < 6; face++ {
		fmt.Printf("Face: %d\n", face)
		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				fmt.Printf("%d%d: \n", y, x)
			}
		}
		fmt.Printf("\n")/////
	}
}

func main() {
	fmt.Printf("oh hi!\n")/////////
	r := initRubik()
	dumpRubik(&r.cube)
	fmt.Printf("\nEND!!\n")//////////
}

// ## To run enter either command:
// go run main.go
// go build; ./Rubik