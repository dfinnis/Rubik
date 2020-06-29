package rubik

import (
	"fmt"
)

const Reset		= "\x1B[0m"
const White		= "\x1B[0m"					// 0 U
const Orange	= "\x1B[38;2;255;165;0m"	// 1 L
const Green		= "\x1B[32m"				// 2 F
const Red		= "\x1B[31m"				// 3 R
const Blue		= "\x1B[34m"				// 4 B
const Yellow	= "\x1B[33m"				// 5 D

func dumpCube(cube *[6]uint32) {
	fmt.Printf("r: %032b\n\n", cube)
	fmt.Printf("face 0:\t%032b\n", cube[0])
	fmt.Printf("face 1:\t%032b\n", cube[1])
	fmt.Printf("face 2:\t%032b\n", cube[2])
	fmt.Printf("face 3:\t%032b\n", cube[3])
	fmt.Printf("face 4:\t%032b\n", cube[4])
	fmt.Printf("face 5:\t%032b\n", cube[5])
	fmt.Printf("\n")//
	// bit := uint8(cube[1] >> 1)
	// fmt.Printf("bit\t%032b\n", bit)
	fmt.Printf("bite me 0: %v\n", cube[0]&1)
	fmt.Printf("bite me 1: %v\n", cube[1]&1)
	fmt.Printf("bite me 2: %v\n", cube[2]&1)
	fmt.Printf("bite me 3: %v\n", cube[3]&1)
	fmt.Printf("bite me 4: %v\n", cube[4]&1)
	fmt.Printf("bite me 5: %v\n", cube[5]&1)
	fmt.Printf("\n")//

	fmt.Printf("bite me 0: %v\n", cube[0]&2)
	fmt.Printf("bite me 1: %v\n", cube[1]&2)
	fmt.Printf("bite me 2: %v\n", cube[2]&2)
	fmt.Printf("bite me 3: %v\n", cube[3]&2)
	fmt.Printf("bite me 4: %v\n", cube[4]&2)
	fmt.Printf("bite me 5: %v\n", cube[5]&2)
	fmt.Printf("\n")//

	fmt.Printf("bite me 0: %v\n", cube[0]&3)
	fmt.Printf("bite me 1: %v\n", cube[1]&3)
	fmt.Printf("bite me 2: %v\n", cube[2]&3)
	fmt.Printf("bite me 3: %v\n", cube[3]&3)
	fmt.Printf("bite me 4: %v\n", cube[4]&3)
	fmt.Printf("bite me 5: %v\n", cube[5]&3)
	fmt.Printf("\n")//

	fmt.Printf("bite me 0: %v\n", cube[0]&4)
	fmt.Printf("bite me 1: %v\n", cube[1]&4)
	fmt.Printf("bite me 2: %v\n", cube[2]&4)
	fmt.Printf("bite me 3: %v\n", cube[3]&4)
	fmt.Printf("bite me 4: %v\n", cube[4]&4)
	fmt.Printf("bite me 5: %v\n", cube[5]&4)
	fmt.Printf("\n")//

	fmt.Printf("bite me 0: %v\n", cube[0]&5)
	fmt.Printf("bite me 1: %v\n", cube[1]&5)
	fmt.Printf("bite me 2: %v\n", cube[2]&5)
	fmt.Printf("bite me 3: %v\n", cube[3]&5)
	fmt.Printf("bite me 4: %v\n", cube[4]&5)
	fmt.Printf("bite me 5: %v\n", cube[5]&5)
	fmt.Printf("\n")//

	fmt.Printf("bite me 0: %v\n", cube[0]&6)
	fmt.Printf("bite me 1: %v\n", cube[1]&6)
	fmt.Printf("bite me 2: %v\n", cube[2]&6)
	fmt.Printf("bite me 3: %v\n", cube[3]&6)
	fmt.Printf("bite me 4: %v\n", cube[4]&6)
	fmt.Printf("bite me 5: %v\n", cube[5]&6)
	fmt.Printf("\n")//

	if cube[0]&5 == 5 {
		fmt.Printf("face 5\n")//
	} else if cube[0]&4 == 4 {
		fmt.Printf("face 4\n")//
	} else if cube[0]&3 == 3 {
		fmt.Printf("face 3\n")//
	} else if cube[0]&2 == 2 {
		fmt.Printf("face 2\n")//
	} else if cube[0]&1 == 1 {
		fmt.Printf("face 1\n")//
	} else { // is 0
		fmt.Printf("face 0\n")//
	}

	fmt.Printf("\n")//	
	for face := 0; face < 6; face++ {
		fmt.Printf("oh hello %v\n", face)
		fmt.Printf("\n")//	
	}

	fmt.Printf("V: %v\n", cube[5]&1342177280)
	if cube[5]&1342177280 == 1342177280 {
		fmt.Printf("0 ")
	}	

	fmt.Printf("\n        ")
	if cube[0]&0x50000000 == 0x50000000 {				// 0101 0000 0000 0000 0000 0000 0000 0000
		fmt.Printf("%v5%v ", Yellow, Reset)
	} else if cube[0]&0x40000000 == 0x40000000 {		// 0100 0000 0000 0000 0000 0000 0000 0000
		fmt.Printf("%v4%v ", Blue, Reset)
	} else if cube[0]&0x30000000 == 0x30000000 {		// 0011 0000 0000 0000 0000 0000 0000 0000
		fmt.Printf("%v3%v ", Red, Reset)
	} else if cube[0]&0x20000000 == 0x20000000 {		// 0010 0000 0000 0000 0000 0000 0000 0000
		fmt.Printf("%v2%v ", Green, Reset)
	} else if cube[0]&0x10000000 == 0x10000000 {		// 0001 0000 0000 0000 0000 0000 0000 0000
		fmt.Printf("%v1%v ", Orange, Reset)
	} else {
		fmt.Printf("0 ")
	}

	fmt.Printf("\n")
}

//	
//	A B C
//  H @ D
//	G F E

// a & 196	query a value for its set bits
// &=		selectively clearing bits of an integer value to zero
// |=		set arbitrary bits for a given integer value