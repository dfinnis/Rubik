package rubik

import (
	"os"
	"io/ioutil"
)

// tables contains pruning tables, and index converters, for all 4 subgroups
type tables struct {
	G0 [2048]uint8
	G1ePindex [4096]int16
	G1 [495][2187]uint8
	G2ePindex [255]uint8
	G2 [40320][70]uint8
	G3 [96][24][24][24]uint8
	G3cPindex [40320]uint8
}

// newNode copies an existing cube
func newNode(parent *cepo, move string) *cepo {
	return &cepo{
		cP:   	parent.cP,
		cO:   	parent.cO,
		eP:   	parent.eP,
		eO:   	parent.eO,
		move:	parent.move,
		move2:	parent.move2,
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

// makeTables initializes tables, and creates or reads tables from file
func makeTables() *tables {
	tables := &tables{}
	makeTableG0(tables)
	makeTableG1(tables)
	makeTableG2(tables)
	makeTableG3(tables)
	return tables
}
