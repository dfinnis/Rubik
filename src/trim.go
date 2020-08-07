package rubik

import (
	"fmt"
	"strings"
)

// e.g. "U U" => "U2"
func replaceMove(sequence []string, move string, i int) []string {
	var trimed []string
	if move == "" {
		trimed = sequence[:i]
	} else {
		trimed = append(sequence[:i], move)
	}
	if i + 2 < len(sequence) {
		trimed = append(trimed, sequence[i+2:]...)
	}
	return trimed
}

// e.g. "U D U" => "U2 D"
func replaceMove2(sequence []string, move string, i int) []string {
	var trimed []string
	trimed = sequence[:i]
	if move != "" {
		trimed = append(trimed, move)
	}
	trimed = append(trimed, sequence[i + 1])
	if i + 3 < len(sequence) {
		trimed = append(trimed, sequence[i+3:]...)
	}
	return trimed
}



// trimSequence concaternates redundant moves to minimize HTM, e.g. "U U" => "U2"
func trimSequence(sequence string) string {
	fmt.Printf("\nsequence: %v\n", sequence)
	trimed := strings.Fields(sequence)
	fmt.Printf("trimed before: %v\n", trimed)
	fmt.Printf("len(trimed): %v\n", len(trimed))//
	for i, move := range trimed {
		fmt.Printf("\nmove: %v\n", move)//
		// fmt.Printf("move[0]: %v\n", move[0])//
		if i + 1 < len(trimed) {
			fmt.Printf("trimed[i + 1]: %v\n", trimed[i + 1])//

			if move == "U" {
				if trimed[i + 1] == "U" {
					trimed = replaceMove(trimed, "U2", i)
				} else if trimed[i + 1] == "U'" {
					trimed = replaceMove(trimed, "", i)
				} else if trimed[i + 1] == "U2" {
					trimed = replaceMove(trimed, "U'", i)
				} else if trimed[i + 1][0] == 'D' && i + 2 < len(trimed) {
					if trimed[i + 2] == "U" {
						trimed = replaceMove2(trimed, "U2", i)
					} else if trimed[i + 2] == "U'" {
						trimed = replaceMove2(trimed, "", i)
					} else if trimed[i + 2] == "U2" {
						trimed = replaceMove2(trimed, "U'", i)
					}
				}
			} else if move == "U'" {
				if trimed[i + 1] == "U" {
					trimed = replaceMove(trimed, "", i)
				} else if trimed[i + 1] == "U'" {
					trimed = replaceMove(trimed, "U2", i)
				} else if trimed[i + 1] == "U2" {
					trimed = replaceMove(trimed, "U", i)
				} else if trimed[i + 1][0] == 'D' && i + 2 < len(trimed) {
					if trimed[i + 2] == "U" {
						trimed = replaceMove2(trimed, "", i)
					} else if trimed[i + 2] == "U'" {
						trimed = replaceMove2(trimed, "U2", i)
					} else if trimed[i + 2] == "U2" {
						trimed = replaceMove2(trimed, "U", i)
					}
				}
			} else if move == "U2" {
				if trimed[i + 1] == "U" {
					trimed = replaceMove(trimed, "U'", i)
				} else if trimed[i + 1] == "U'" {
					trimed = replaceMove(trimed, "U", i)
				} else if trimed[i + 1] == "U2" {
					trimed = replaceMove(trimed, "", i)
				} else if trimed[i + 1][0] == 'D' && i + 2 < len(trimed) {
					if trimed[i + 2] == "U" {
						trimed = replaceMove2(trimed, "U'", i)
					} else if trimed[i + 2] == "U'" {
						trimed = replaceMove2(trimed, "U", i)
					} else if trimed[i + 2] == "U2" {
						trimed = replaceMove2(trimed, "", i)
					}
				}
			// } else if move == "D" {
			// 	if trimed[i + 1] == "D" {
			// 		trimed = replaceMove(trimed, "D2", i)
			// 	} else if trimed[i + 1] == "D'" {
			// 		trimed = replaceMove(trimed, "", i)
			// 	} else if trimed[i + 1] == "D2" {
			// 		trimed = replaceMove(trimed, "D'", i)
			// 	} else if trimed[i + 1][0] == 'U' && i + 2 < len(trimed) {
			// 		if trimed[i + 2] == "D" {
			// 			trimed = replaceMove2(trimed, "D2", i)
			// 		} else if trimed[i + 2] == "D'" {
			// 			trimed = replaceMove2(trimed, "", i)
			// 		} else if trimed[i + 2] == "D2" {
			// 			trimed = replaceMove2(trimed, "D'", i)
			// 		}
			// 	}
			// } else if move == "D'" {
			// 	if trimed[i + 1] == "D" {
			// 		trimed = replaceMove(trimed, "", i)
			// 	} else if trimed[i + 1] == "D'" {
			// 		trimed = replaceMove(trimed, "D2", i)
			// 	} else if trimed[i + 1] == "D2" {
			// 		trimed = replaceMove(trimed, "D", i)
			// 	} else if trimed[i + 1][0] == 'D' && i + 2 < len(trimed) {
			// 		if trimed[i + 2] == "D" {
			// 			trimed = replaceMove2(trimed, "", i)
			// 		} else if trimed[i + 2] == "D'" {
			// 			trimed = replaceMove2(trimed, "D2", i)
			// 		} else if trimed[i + 2] == "D2" {
			// 			trimed = replaceMove2(trimed, "D", i)
			// 		}
			// 	}
			// } else if move == "D2" {
			// 	if trimed[i + 1] == "D" {
			// 		trimed = replaceMove(trimed, "D'", i)
			// 	} else if trimed[i + 1] == "D'" {
			// 		trimed = replaceMove(trimed, "D", i)
			// 	} else if trimed[i + 1] == "D2" {
			// 		trimed = replaceMove(trimed, "", i)
			// 	} else if trimed[i + 1][0] == 'U' && i + 2 < len(trimed) {
			// 		if trimed[i + 2] == "D" {
			// 			trimed = replaceMove2(trimed, "D'", i)
			// 		} else if trimed[i + 2] == "D'" {
			// 			trimed = replaceMove2(trimed, "D", i)
			// 		} else if trimed[i + 2] == "D2" {
			// 			trimed = replaceMove2(trimed, "", i)
			// 		}
			// 	}
			}
		}
	}
	fmt.Printf("trimed after:  %v\n", trimed)
	fmt.Printf("len(trimed): %v\n", len(trimed))//
	var trimedString string
	for i := 0; i < len(trimed); i++ {
		trimedString += trimed[i] + " "
	}
	return trimedString
}