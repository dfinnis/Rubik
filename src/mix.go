package rubik

import (
	"fmt"
	"strings"
	"os"
	"io"
)

// mixIsValid returns true if mix only contains valid moves
func mixIsValid(mix string) (valid bool) {
	allMoves := listMoves(initCube(), 0)
	mixList := strings.Fields(mix)
	for _, mixMove := range mixList {
		var found bool
		for _, move := range allMoves {
			if move == mixMove {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// makeMix checks mix validity, creates random mix, or reads mix file
func makeMix(mix string, length int) string {
	if mixIsValid(mix) {
		return mix
	}
	if mix == "-r" || mix == "--random" {
		return randomMix(length)
	}
	file, err := os.Open(mix)
	if err != nil {
		errorExit("failed to open mix file")
	}
	defer file.Close()
	fileStr := make([]byte, 600)
    len, err := file.Read(fileStr)
	if err != nil && err != io.EOF {
		errorExit("failed to read mix file")
	}
	if len > 599 {
		errorExit("file too long")
	}
	filepath := mix
	if err == io.EOF {
		mix = ""
	} else {
		mix = string(fileStr[:len])
	}
	fmt.Printf("\nMix read from filepath \"%v\":\n%v\n\n", filepath, mix)
	return mix
}
