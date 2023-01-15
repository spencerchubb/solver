package solver

func SolvedFacelets() [48]int {
	return [48]int{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
}

func faceletsEq(a, b [48]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func faceletsToStr(arr [48]int) string {
	bytes := make([]byte, 48)
	for i, facelet := range arr {
		switch facelet {
		case U:
			bytes[i] = 'U'
		case F:
			bytes[i] = 'F'
		case D:
			bytes[i] = 'D'
		case B:
			bytes[i] = 'B'
		case L:
			bytes[i] = 'L'
		case R:
			bytes[i] = 'R'
		}
	}
	return string(bytes)
}

func faceletsFromStr(str string) []int {
	arr := make([]int, len(str))
	for i, c := range str {
		switch c {
		case 'U':
			arr[i] = U
		case 'F':
			arr[i] = F
		case 'D':
			arr[i] = D
		case 'B':
			arr[i] = B
		case 'L':
			arr[i] = L
		case 'R':
			arr[i] = R
		}
	}
	return arr
}
