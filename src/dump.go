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

func dumpTop(cube *[6]uint32, face uint8) {
	// A	// top left corner
	if cube[face]&0x50000000 == 0x50000000 {			// 0101 0000 0000 0000 0000 0000 0000 0000
		fmt.Printf("%v5%v ", Yellow, Reset)
	} else if cube[face]&0x40000000 == 0x40000000 {		// 0100 0000 0000 0000 0000 0000 0000 0000
		fmt.Printf("%v4%v ", Blue, Reset)
	} else if cube[face]&0x30000000 == 0x30000000 {		// 0011 0000 0000 0000 0000 0000 0000 0000
		fmt.Printf("%v3%v ", Red, Reset)
	} else if cube[face]&0x20000000 == 0x20000000 {		// 0010 0000 0000 0000 0000 0000 0000 0000
		fmt.Printf("%v2%v ", Green, Reset)
	} else if cube[face]&0x10000000 == 0x10000000 {		// 0001 0000 0000 0000 0000 0000 0000 0000
		fmt.Printf("%v1%v ", Orange, Reset)
	} else {
		fmt.Printf("0 ")
	}

	// B	// top edge
	if cube[face]&0x05000000 == 0x05000000 {
		fmt.Printf("%v5%v ", Yellow, Reset)
	} else if cube[face]&0x04000000 == 0x04000000 {
		fmt.Printf("%v4%v ", Blue, Reset)
	} else if cube[face]&0x03000000 == 0x03000000 {
		fmt.Printf("%v3%v ", Red, Reset)
	} else if cube[face]&0x02000000 == 0x02000000 {
		fmt.Printf("%v2%v ", Green, Reset)
	} else if cube[face]&0x01000000 == 0x01000000 {
		fmt.Printf("%v1%v ", Orange, Reset)
	} else {
		fmt.Printf("0 ")
	}

	// C	// top right corner
	if cube[face]&0x00500000 == 0x00500000 {
		fmt.Printf("%v5%v ", Yellow, Reset)
	} else if cube[face]&0x00400000 == 0x00400000 {
		fmt.Printf("%v4%v ", Blue, Reset)
	} else if cube[face]&0x00300000 == 0x00300000 {
		fmt.Printf("%v3%v ", Red, Reset)
	} else if cube[face]&0x00200000 == 0x00200000 {
		fmt.Printf("%v2%v ", Green, Reset)
	} else if cube[face]&0x00100000 == 0x00100000 {
		fmt.Printf("%v1%v ", Orange, Reset)
	} else {
		fmt.Printf("0 ")
	}
}

func dumpMiddle(cube *[6]uint32, face uint8) {
	// H	// left edge
	if cube[face]&0x00000005 == 0x00000005 {
		fmt.Printf("%v5%v ", Yellow, Reset)
	} else if cube[face]&0x00000004 == 0x00000004 {
		fmt.Printf("%v4%v ", Blue, Reset)
	} else if cube[face]&0x00000003 == 0x00000003 {
		fmt.Printf("%v3%v ", Red, Reset)
	} else if cube[face]&0x00000002 == 0x00000002 {
		fmt.Printf("%v2%v ", Green, Reset)
	} else if cube[face]&0x00000001 == 0x00000001 {
		fmt.Printf("%v1%v ", Orange, Reset)
	} else {
		fmt.Printf("0 ")
	}

	// Center
	if face == 5 {
		fmt.Printf("%v5%v ", Yellow, Reset)
	} else if face == 4 {
		fmt.Printf("%v4%v ", Blue, Reset)
	} else if face == 3 {
		fmt.Printf("%v3%v ", Red, Reset)
	} else if face == 2 {
		fmt.Printf("%v2%v ", Green, Reset)
	} else if face == 1 {
		fmt.Printf("%v1%v ", Orange, Reset)
	} else {
		fmt.Printf("0 ")
	}

	// D	// right edge
	if cube[face]&0x00050000 == 0x00050000 {
		fmt.Printf("%v5%v ", Yellow, Reset)
	} else if cube[face]&0x00040000 == 0x00040000 {
		fmt.Printf("%v4%v ", Blue, Reset)
	} else if cube[face]&0x00030000 == 0x00030000 {
		fmt.Printf("%v3%v ", Red, Reset)
	} else if cube[face]&0x00020000 == 0x00020000 {
		fmt.Printf("%v2%v ", Green, Reset)
	} else if cube[face]&0x00010000 == 0x00010000 {
		fmt.Printf("%v1%v ", Orange, Reset)
	} else {
		fmt.Printf("0 ")
	}
}

func dumpBottom(cube *[6]uint32, face uint8) {
	// G	// bottom left corner
	if cube[face]&0x00000050 == 0x00000050 {
		fmt.Printf("%v5%v ", Yellow, Reset)
	} else if cube[face]&0x00000040 == 0x00000040 { 
		fmt.Printf("%v4%v ", Blue, Reset)
	} else if cube[face]&0x00000030 == 0x00000030 { 
		fmt.Printf("%v3%v ", Red, Reset)
	} else if cube[face]&0x00000020 == 0x00000020 { 
		fmt.Printf("%v2%v ", Green, Reset)
	} else if cube[face]&0x00000010 == 0x00000010 { 
		fmt.Printf("%v1%v ", Orange, Reset)
	} else {
		fmt.Printf("0 ")
	}

	// F	// bottom edge
	if cube[face]&0x00000500 == 0x00000500 {
		fmt.Printf("%v5%v ", Yellow, Reset)
	} else if cube[face]&0x00000400 == 0x00000400 {
		fmt.Printf("%v4%v ", Blue, Reset)
	} else if cube[face]&0x00000300 == 0x00000300 {
		fmt.Printf("%v3%v ", Red, Reset)
	} else if cube[face]&0x00000200 == 0x00000200 {
		fmt.Printf("%v2%v ", Green, Reset)
	} else if cube[face]&0x00000100 == 0x00000100 {
		fmt.Printf("%v1%v ", Orange, Reset)
	} else {
		fmt.Printf("0 ")
	}

	// E	// bottom right corner
	if cube[face]&0x00005000 == 0x00005000 {
		fmt.Printf("%v5%v ", Yellow, Reset)
	} else if cube[face]&0x00004000 == 0x00004000 {
		fmt.Printf("%v4%v ", Blue, Reset)
	} else if cube[face]&0x00003000 == 0x00003000 {
		fmt.Printf("%v3%v ", Red, Reset)
	} else if cube[face]&0x00002000 == 0x00002000 {
		fmt.Printf("%v2%v ", Green, Reset)
	} else if cube[face]&0x00001000 == 0x00001000 {
		fmt.Printf("%v1%v ", Orange, Reset)
	} else {
		fmt.Printf("0 ")
	}
}

func dumpFace(cube *[6]uint32, face uint8) {
	//	A B C	//	top
	//  H @ D	//	middle
	//	G F E	//	bottom
	fmt.Printf("\n        ")
	dumpTop(cube, face)
	fmt.Printf("\n        ")
	dumpMiddle(cube, face)
	fmt.Printf("\n        ")
	dumpBottom(cube, face)
	fmt.Printf("\n")
}

func dumpLFRB(cube *[6]uint32) {
	fmt.Printf(" ")
	var face uint8
	for face = 1; face < 5; face++ {
		dumpTop(cube, face)
		fmt.Printf(" ")
	}
	fmt.Printf("\n ")
	for face = 1; face < 5; face++ {
		dumpMiddle(cube, face)
		fmt.Printf(" ")
	}
	fmt.Printf("\n ")
	for face = 1; face < 5; face++ {
		dumpBottom(cube, face)
		fmt.Printf(" ")
	}
}

func dumpBinary(cube *[6]uint32) {
	// fmt.Printf("\ncube binary: %032b\n\n", cube)
	fmt.Printf("cube %x\n", cube)
	fmt.Printf("\n	A   B   C   D   E   F   G   H\n")
	fmt.Printf("face 0:\t%032b\n", cube[0])
	fmt.Printf("	A   B   C   D   E   F   G   H\n")
	fmt.Printf("face 1:\t%032b\n", cube[1])
	fmt.Printf("	A   B   C   D   E   F   G   H\n")
	fmt.Printf("face 2:\t%032b\n", cube[2])
	fmt.Printf("	A   B   C   D   E   F   G   H\n")
	fmt.Printf("face 3:\t%032b\n", cube[3])
	fmt.Printf("	A   B   C   D   E   F   G   H\n")
	fmt.Printf("face 4:\t%032b\n", cube[4])
	fmt.Printf("	A   B   C   D   E   F   G   H\n")
	fmt.Printf("face 5:\t%032b\n", cube[5])
}

func dumpCube(cube *[6]uint32) {
	fmt.Printf("\n\n#### -- CUBE -- ####\n")
	// dumpBinary(cube)	//	binary debug
	dumpFace(cube, 0)
	dumpLFRB(cube)
	dumpFace(cube, 5)
}
