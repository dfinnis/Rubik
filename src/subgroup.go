package rubik

// isSubgroup returns subgroup for given cube
func isSubgroup(cube *cepo) int8 {
	for i := range cube.eO { // edges not oriented -> 0
		if cube.eO[i] != 0 {
			return 0
		}
	}
	for i := range cube.cO { // corners not oriented -> 1
		if cube.eO[i] != 0 {
			return 1
		}
	}
	for i := 8; i <= 11; i++ { // edges LR slice not in slice -> 1
		if cube.eP[i] < 8 {
			return 1
		}
	}
	for i := 0; i <= 3; i++ { // corners not in tetrads -> 2
		if cube.cP[i] > 3 {
			return 2
		}
	}
	for i := 4; i <= 7; i++ { // corners not in tetrads -> 2
		if cube.cP[i] < 4 {
			return 2
		}
	}
	for i := 0; i <= 3; i++ { // edges FB slice not in slice -> 2
		if cube.eP[i] > 3 {
			return 2
		}
	}
	for i := 4; i <= 7; i++ { // edges UD slice not in slice -> 2
		if cube.eP[i] < 4 || cube.eP[i] > 7 {
			return 2
		}
	}
	var parity int8
	for i := range cube.cP { // corners parity odd -> 2
		if cube.cP[i] == int8(i) {
			parity++
		}
	}
	if parity % 2 != 0 {
		return 2
	}
	if isSolved(cube) == false {
		return 3
	}
	return 4
}
