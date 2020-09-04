package rubik

import "fmt"

// dumpCube prints cube state
func dumpCube(cube *cepo) {
	fmt.Printf("\n+-------------+-----------------+-------------------------------------+")
	fmt.Printf("\n|             | Corner          |  Edge                               |\n")
	fmt.Printf("|      Number | ")
	for i := range cube.cP {
		fmt.Printf("%v ", i)
	}
	fmt.Printf("| ")
	for i := range cube.eP {
		fmt.Printf("%2v ", i)
	}
	fmt.Printf("|")
	fmt.Printf("\n+-------------+-----------------+-------------------------------------+\n")
	fmt.Printf("| Permutation | ")
	for _, cP := range cube.cP {
		fmt.Printf("%v ", cP)
	}
	fmt.Printf("| ")
	for _, eP := range cube.eP {
		fmt.Printf("%2v ", eP)
	}
	fmt.Printf("|")

	fmt.Printf("\n| Orientation | ")
	for _, cO := range cube.cO {
		fmt.Printf("%v ", cO)
	}
	fmt.Printf("| ")
	for _, eO := range cube.eO {
		fmt.Printf("%2v ", eO)
	}
	fmt.Printf("|")
	fmt.Printf("\n+-------------+-----------------+-------------------------------------+\n")
}
