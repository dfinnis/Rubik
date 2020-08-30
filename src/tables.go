package rubik

import (
	"fmt"
	"os"
	"io/ioutil"
)

type tables struct {
	G0 [2048]uint8
	G1ePindex [4096]int16
	G1 [495][2187]uint8
	G2ePindex [255]uint8
	G2 [40320][70]uint8
	G3 [96][24][24][24]uint8
	G3cPindex [40320]uint8
}

func move2nul(move string, move2 string) string {
	if move != "" && move2 != "" {
		if move[0] == 'U' && move2[0] != 'D' || 
			move[0] == 'D' && move2[0] != 'U' ||
			move[0] == 'F' && move2[0] != 'B' || 
			move[0] == 'B' && move2[0] != 'F' ||
			move[0] == 'L' && move2[0] != 'R' || 
			move[0] == 'R' && move2[0] != 'L' {
			move2 = ""
		}
	}
	return move2
}

func newNode(parent *cepo, move string) *cepo {
	return &cepo{
		cP:   	parent.cP,
		cO:   	parent.cO,
		eP:   	parent.eP,
		eO:   	parent.eO,
		move2:	parent.move2,
		move:	parent.move,
	}
}

func createFile(filepath string) *os.File {
	file, err := os.Create(filepath)
	if err != nil {
		errorExit("failed to create file")
	}
	return file
}

func readFile(filepath string) []byte {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		errorExit("failed to read pruning table file")
	}
	return file
}

func makeTables() *tables {
	tables := &tables{}
	makeTableG0(tables)
	makeTableG1(tables)
	makeTableG2(tables)
	makeTableG3(tables)
	fmt.Println()
	return tables
}