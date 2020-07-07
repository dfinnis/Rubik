package rubik

import (
	"math/rand"
	"fmt"
	"time"
)

// func randomSpin() string {
// 	spin := []string{
// 		"U",
// 		"U'",
// 		"U2",
// 		"D",
// 		"D'",
// 		"D2",
// 		"R",
// 		"R'",
// 		"R2",
// 		"L",
// 		"L'",
// 		"L2",
// 		"F",
// 		"F'",
// 		"F2",
// 		"B",
// 		"B'",
// 		"B2",
// 	}
// 	return spin[rand.Intn(len(spin))]
// }

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func randomMixDry() string {
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
	// var repeat0 string
	// var repeat1 string
	fmt.Printf("spin: %v\n", spin[:])//
	fmt.Printf("dry: %v\n", dry[:])//
	// upper := spin[:]
	// lower := spin[:]
	// upper = append(upper, lower...)
	// fmt.Printf("upper: %v\n", upper)
	// fmt.Printf("lower: %v\n", lower)

	var mix string
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(4) + 20
	for i := 0; i <= n; i++ {
		move := dry[rand.Intn(len(dry))]
		mix += move
		fmt.Printf("\nmove: %v\n", move)
		// dry = updateDry(dry, move)//
		if move == "U" || move == "U'" || move == "U2" {
			if stringInSlice("D", dry) {
				dry = spin[3:]
			} else {
				dry = spin[6:]
			}
		} else if move == "D" || move == "D'" || move == "D2" {
			if stringInSlice("U", dry) {
				dry = spin[6:]
				dry = append(dry, spin[0])
				dry = append(dry, spin[1])
				dry = append(dry, spin[2])
			} else {
				dry = spin[6:]
			}
		// } else if move == "R" || move == "R'" || move == "R2" {
		// 	// dry = remove3(dry, "R", "R'", "R2")
		// } else if move == "L" || move == "L'" || move == "L2" {
		// 	// dry = remove3(dry, "L", "L'", "L2")
		// } else if move == "F" || move == "F'" || move == "F2" {
		// 	// dry = remove3(dry, "F", "F'", "F2")
		// } else if move == "B" || move == "B'" || move == "B2" {
		// 	// dry = remove3(dry, "B", "B'", "B2")
		// }
		}
		fmt.Printf("\ndry: %v\n", dry)////
		fmt.Printf("spin: %v\n", spin[:])//
//////////////////////////////////////////////////////////////
		if i != n {
			mix += " "
		}
	}
	fmt.Printf("\nRandom Mix: %v\n", mix)
	return mix
}


//randomMix returns a random 20 to 24 spin long mix
func randomMix() string {
	return randomMixDry()////
	// var mix string
	// rand.Seed(time.Now().UnixNano())
	// n := rand.Intn(4) + 20
	// for i := 0; i <= n; i++ {
	// 	mix += randomSpin()
	// 	if i != n {
	// 		mix += " "
	// 	}
	// }
	// fmt.Printf("\nRandom Mix: %v\n", mix)
	// return mix
}