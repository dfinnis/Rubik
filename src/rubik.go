package rubik

import (
	"fmt"
	"math/bits"
)

func test() {
	fmt.Printf("test!\n")/////////
	var face uint64
	face = 1
	fmt.Printf("\nint before: %v\n", face)/////////
	face = bits.RotateLeft64(face, 2)
	fmt.Printf("\nint after: %v\n\n", face)/////////
	face = bits.RotateLeft64(face, -1)
	fmt.Printf("\nint after: %v\n\n", face)/////////
	fmt.Printf("test end!\n")/////////
}

func RunRubik() {
	test()
}
