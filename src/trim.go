package rubik

import (
	"strings"
)

// replaceMove concaternates 2 consecutive moves e.g. "U U" => "U2"
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

// replaceMove2 concaternates 2 non-consecutive moves e.g. "U D U" => "U2 D"
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

func assignMoves(move byte) (string, string, string, byte) {
	var quarter string
	var anti string
	var half string
	var opposite byte
	if move == 'U' {
		quarter = "U"
		anti = "U'"
		half = "U2"
		opposite = 'D'
	} else if move == 'D' {
		quarter = "D"
		anti = "D'"
		half = "D2"
		opposite = 'U'
	} else if move == 'R' {
		quarter = "R"
		anti = "R'"
		half = "R2"
		opposite = 'L'
	} else if move == 'L' {
		quarter = "L"
		anti = "L'"
		half = "L2"
		opposite = 'R'
	} else if move == 'F' {
		quarter = "F"
		anti = "F'"
		half = "F2"
		opposite = 'B'
	} else { // move = 'B'
		quarter = "B"
		anti = "B'"
		half = "B2"
		opposite = 'F'
	}
	return quarter, anti, half, opposite
}

// trimSequence concaternates redundant moves, e.g. "U U" => "U2"
func trimSequence(sequence string) string {
	trimed := strings.Fields(sequence)
	for i, move := range trimed {
		if i + 1 < len(trimed) {
			quarter, anti, half, opposite := assignMoves(move[0])
			if move == quarter {
				if trimed[i + 1] == quarter {
					trimed = replaceMove(trimed, half, i)
				} else if trimed[i + 1] == anti {
					trimed = replaceMove(trimed, "", i)
				} else if trimed[i + 1] == half {
					trimed = replaceMove(trimed, anti, i)
				} else if trimed[i + 1][0] == opposite && i + 2 < len(trimed) {
					if trimed[i + 2] == quarter {
						trimed = replaceMove2(trimed, half, i)
					} else if trimed[i + 2] == anti {
						trimed = replaceMove2(trimed, "", i)
					} else if trimed[i + 2] == half {
						trimed = replaceMove2(trimed, anti, i)
					}
				}
			} else if move == anti {
				if trimed[i + 1] == quarter {
					trimed = replaceMove(trimed, "", i)
				} else if trimed[i + 1] == anti {
					trimed = replaceMove(trimed, half, i)
				} else if trimed[i + 1] == half {
					trimed = replaceMove(trimed, quarter, i)
				} else if trimed[i + 1][0] == opposite && i + 2 < len(trimed) {
					if trimed[i + 2] == quarter {
						trimed = replaceMove2(trimed, "", i)
					} else if trimed[i + 2] == anti {
						trimed = replaceMove2(trimed, half, i)
					} else if trimed[i + 2] == half {
						trimed = replaceMove2(trimed, quarter, i)
					}
				}
			} else if move == half {
				if trimed[i + 1] == quarter {
					trimed = replaceMove(trimed, anti, i)
				} else if trimed[i + 1] == anti {
					trimed = replaceMove(trimed, quarter, i)
				} else if trimed[i + 1] == half {
					trimed = replaceMove(trimed, "", i)
				} else if trimed[i + 1][0] == opposite && i + 2 < len(trimed) {
					if trimed[i + 2] == quarter {
						trimed = replaceMove2(trimed, anti, i)
					} else if trimed[i + 2] == anti {
						trimed = replaceMove2(trimed, quarter, i)
					} else if trimed[i + 2] == half {
						trimed = replaceMove2(trimed, "", i)
					}
				}
			}
		}
	}
	var trimedString string
	for i := 0; i < len(trimed); i++ {
		trimedString += trimed[i] + " "
	}
	return trimedString
}

// trim concaternates redundant moves to minimize Half Turn Metric
func trim(sequence string) string {
	for trimSequence(sequence) != sequence {
		sequence = trimSequence(sequence)
	}
	return sequence
}
