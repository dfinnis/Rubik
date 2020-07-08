package rubik

import (
	"math/rand"
	"fmt"
	"time"
)

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

//randomMixDry returns a random 19 to 25 spin long mix
func randomMix() string {
	spin := []string{
		"U",
		"U'",
		"U2",
		"D",
		"D'",
		"D2",
		"R",
		"R'",
		"R2",
		"L",
		"L'",
		"L2",
		"F",
		"F'",
		"F2",
		"B",
		"B'",
		"B2",
	}
	dry := spin
	var mix string
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(6) + 19
	for i := 0; i <= n; i++ {
		move := dry[rand.Intn(len(dry))]
		mix += move
		if move == "U" || move == "U'" || move == "U2" {
			if stringInSlice("D", dry) {
				dry = spin[3:]
			} else {
				dry = spin[6:]
			}
		} else if move == "D" || move == "D'" || move == "D2" {
			if stringInSlice("U", dry) {
				dry = append([]string{}, spin[:3]...)
				dry = append(dry, spin[6:]...)
			} else {
				dry = spin[6:]
			}
		} else if move == "R" || move == "R'" || move == "R2" {
			dry = append([]string{}, spin[:6]...)
			if stringInSlice("L", dry) {
				dry = append(dry, spin[9:]...)
			} else {
				dry = append(dry, spin[12:]...)
			}
		} else if move == "L" || move == "L'" || move == "L2" {
			if stringInSlice("R", dry) {
				dry = append([]string{}, spin[:9]...)
			} else {
				dry = append([]string{}, spin[:6]...)
			}
			dry = append(dry, spin[12:]...)
		} else if move == "F" || move == "F'" || move == "F2" {
			dry = append([]string{}, spin[:12]...)
			if stringInSlice("B", dry) {
				dry = append(dry, spin[15:]...)
			}
		} else if move == "B" || move == "B'" || move == "B2" {
			if stringInSlice("F", dry) {
				dry = spin[:15]
			} else {
				dry = spin[:12]
			}
		}
		if i != n {
			mix += " "
		}
	}
	fmt.Printf("\nRandom Mix: %v\n", mix)
	return mix
}