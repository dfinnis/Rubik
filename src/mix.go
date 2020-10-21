package rubik

import (
	"fmt"
	"strings"
	"io/ioutil"
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
	file, err := ioutil.ReadFile(mix)
	if err != nil {
		errorExit("failed to read mix file")
	}
	if len(file) > 200 {
		errorExit("file too long")
	}
	filepath := mix
	mix = string(file)
	fmt.Printf("\nMix read from filepath \"%v\":\n%v\n\n", filepath, mix)
	return mix
}
