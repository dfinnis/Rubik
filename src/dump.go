package rubik

import (
	"fmt"
)

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
}

// a & 196	query a value for its set bits
// &=		selectively clearing bits of an integer value to zero
// |=		set arbitrary bits for a given integer value