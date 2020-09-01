package rubik

import (
	"fmt"
	"math"
	"os"
)

func binaryToDecimal(binary [12]int8) int {
	var decimal int
	for i := 0; i < 11; i++ {
		decimal += int(binary[i]) * int(math.Pow(2, float64(10-i)))
	}
	return decimal
}

func tableG0() [2048]uint8 {
	fmt.Printf("\nGenerating pruning table for G0")
	var table [2048]uint8
	var parents []cepo
	parents = append(parents, *initCube())
	var depth uint8
	for depth < 6 {
		var children []cepo
		depth++
		for _, parent := range parents {
			for _, move := range listMoves(&parent, 0) {
				child := newNode(&parent, move)
				spin(move, child)
				index := binaryToDecimal(child.eO)
				if index != 0 && table[index] == 0 {
					table[index] = depth
				}
				children = append(children, *child)
			}
		}
		parents = children
		fmt.Printf(".")
	}
	for i, depth := range table {
		if i > 0 && depth == 0 {
			table[i] = 7
		}
	}
	return table
}

func makeTableG0(tables *tables) {
	if _, err := os.Stat("tables/G0.txt"); os.IsNotExist(err) {
		tables.G0 = tableG0()
		file := createFile("tables/G0.txt")
		defer file.Close()
		for i := 0; i < len(tables.G0); i++ {
			_, err = file.WriteString(fmt.Sprintf("%d", tables.G0[i]))
			if err != nil {
				errorExit("failed to write to file")
			}
		}
	} else {
		file := readFile("tables/G0.txt")
		for i, depth := range file {
			tables.G0[i] = depth - 48
		}
	}
}