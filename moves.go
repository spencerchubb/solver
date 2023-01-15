package solver

import (
	"bytes"
	"strings"
)

//          00 03 05
//		    01    06
//          02 04 07
// 32 35 37 08 11 13 40 43 45 31 28 26
// 33    38 09    14 41    46 30	25
// 34 36 39 10 12 15 42 44 47 29 27 24
//          16 19 21
//          17    22
//          18 20 23

type Move struct {
	name string
	proc func(*[48]int)
}

var moves = []Move{
	{"U1", U1},
	{"U2", U2},
	{"U3", U3},
	{"F1", F1},
	{"F2", F2},
	{"F3", F3},
	{"D1", D1},
	{"D2", D2},
	{"D3", D3},
	{"B1", B1},
	{"B2", B2},
	{"B3", B3},
	{"L1", L1},
	{"L2", L2},
	{"L3", L3},
	{"R1", R1},
	{"R2", R2},
	{"R3", R3},
}

func moveSubset(moveNames []string) []Move {
	output := make([]Move, len(moveNames))
	idx := 0
	for _, moveName := range moveNames {
		for _, move := range moves {
			if moveName == move.name {
				output[idx] = move
				idx++
			}
		}
	}
	return output
}

func swap2(arr *[48]int, a, b int) {
	temp := arr[b]
	arr[b] = arr[a]
	arr[a] = temp
}

func swap4(arr *[48]int, a, b, c, d int) {
	temp := arr[d]
	arr[d] = arr[c]
	arr[c] = arr[b]
	arr[b] = arr[a]
	arr[a] = temp
}

func performMultipleMoves(facelets *[48]int, moveNames []string) {
	for _, moveName := range moveNames {
		switch moveName {
		case "U1":
			U1(facelets)
		case "U2":
			U2(facelets)
		case "U3":
			U3(facelets)
		case "F1":
			F1(facelets)
		case "F2":
			F2(facelets)
		case "F3":
			F3(facelets)
		case "D1":
			D1(facelets)
		case "D2":
			D2(facelets)
		case "D3":
			D3(facelets)
		case "B1":
			B1(facelets)
		case "B2":
			B2(facelets)
		case "B3":
			B3(facelets)
		case "L1":
			L1(facelets)
		case "L2":
			L2(facelets)
		case "L3":
			L3(facelets)
		case "R1":
			R1(facelets)
		case "R2":
			R2(facelets)
		case "R3":
			R3(facelets)
		}
	}
}

func algString(forward []string, inverse []string) string {
	buff := bytes.NewBufferString("")
	inverseHasMoves := len(inverse) > 0
	for i := 0; i < len(forward); i++ {
		move := forward[i]
		buff.WriteByte(move[0])
		switch move[1] {
		case '1':
			break
		case '2':
			buff.WriteByte('2')
		case '3':
			buff.WriteByte('\'')
		}

		// Don't add a space after the last move if there are no inverse moves
		if i < len(forward)-1 || inverseHasMoves {
			buff.WriteByte(' ')
		}
	}
	for i := len(inverse) - 1; i >= 0; i-- {
		move := inverse[i]
		buff.WriteByte(move[0])
		switch move[1] {
		case '1':
			buff.WriteByte('\'')
		case '2':
			buff.WriteByte('2')
		case '3':
			break
		}

		if i > 0 {
			buff.WriteByte(' ')
		}
	}
	return buff.String()
}

func performAlgorithm(facelets *[48]int, algorithm string) {
	for _, move := range strings.Split(algorithm, " ") {
		switch move {
		case "U":
			U1(facelets)
		case "U2":
			U2(facelets)
		case "U'":
			U3(facelets)
		case "F":
			F1(facelets)
		case "F2":
			F2(facelets)
		case "F'":
			F3(facelets)
		case "D":
			D1(facelets)
		case "D2":
			D2(facelets)
		case "D'":
			D3(facelets)
		case "B":
			B1(facelets)
		case "B2":
			B2(facelets)
		case "B'":
			B3(facelets)
		case "L":
			L1(facelets)
		case "L2":
			L2(facelets)
		case "L'":
			L3(facelets)
		case "R":
			R1(facelets)
		case "R2":
			R2(facelets)
		case "R'":
			R3(facelets)
		}
	}
}

func U1(facelets *[48]int) {
	swap4(facelets, 00, 05, 07, 02)
	swap4(facelets, 01, 03, 06, 04)
	swap4(facelets, 8, 32, 31, 40)
	swap4(facelets, 11, 35, 28, 43)
	swap4(facelets, 13, 37, 26, 45)
}

func U2(facelets *[48]int) {
	swap2(facelets, 00, 07)
	swap2(facelets, 05, 02)
	swap2(facelets, 01, 06)
	swap2(facelets, 03, 04)
	swap2(facelets, 8, 31)
	swap2(facelets, 32, 40)
	swap2(facelets, 11, 28)
	swap2(facelets, 35, 43)
	swap2(facelets, 13, 26)
	swap2(facelets, 37, 45)
}

func U3(facelets *[48]int) {
	swap4(facelets, 00, 02, 07, 05)
	swap4(facelets, 01, 04, 06, 03)
	swap4(facelets, 8, 40, 31, 32)
	swap4(facelets, 11, 43, 28, 35)
	swap4(facelets, 13, 45, 26, 37)
}

func F1(facelets *[48]int) {
	swap4(facelets, 8, 13, 15, 10)
	swap4(facelets, 9, 11, 14, 12)
	swap4(facelets, 02, 40, 21, 39)
	swap4(facelets, 04, 41, 19, 38)
	swap4(facelets, 07, 42, 16, 37)
}

func F2(facelets *[48]int) {
	swap2(facelets, 8, 15)
	swap2(facelets, 13, 10)
	swap2(facelets, 9, 14)
	swap2(facelets, 11, 12)
	swap2(facelets, 02, 21)
	swap2(facelets, 40, 39)
	swap2(facelets, 04, 19)
	swap2(facelets, 41, 38)
	swap2(facelets, 07, 16)
	swap2(facelets, 42, 37)
}

func F3(facelets *[48]int) {
	swap4(facelets, 8, 10, 15, 13)
	swap4(facelets, 9, 12, 14, 11)
	swap4(facelets, 02, 39, 21, 40)
	swap4(facelets, 04, 38, 19, 41)
	swap4(facelets, 07, 37, 16, 42)
}

func D1(facelets *[48]int) {
	swap4(facelets, 16, 21, 23, 18)
	swap4(facelets, 17, 19, 22, 20)
	swap4(facelets, 10, 42, 29, 34)
	swap4(facelets, 12, 44, 27, 36)
	swap4(facelets, 15, 47, 24, 39)
}

func D2(facelets *[48]int) {
	swap2(facelets, 16, 23)
	swap2(facelets, 21, 18)
	swap2(facelets, 17, 22)
	swap2(facelets, 19, 20)
	swap2(facelets, 10, 29)
	swap2(facelets, 42, 34)
	swap2(facelets, 12, 27)
	swap2(facelets, 44, 36)
	swap2(facelets, 15, 24)
	swap2(facelets, 47, 39)
}

func D3(facelets *[48]int) {
	swap4(facelets, 16, 18, 23, 21)
	swap4(facelets, 17, 20, 22, 19)
	swap4(facelets, 10, 34, 29, 42)
	swap4(facelets, 12, 36, 27, 44)
	swap4(facelets, 15, 39, 24, 47)
}

func B1(facelets *[48]int) {
	swap4(facelets, 24, 29, 31, 26)
	swap4(facelets, 25, 27, 30, 28)
	swap4(facelets, 00, 34, 23, 45)
	swap4(facelets, 03, 33, 20, 46)
	swap4(facelets, 05, 32, 18, 47)
}

func B2(facelets *[48]int) {
	swap2(facelets, 24, 31)
	swap2(facelets, 29, 26)
	swap2(facelets, 25, 30)
	swap2(facelets, 27, 28)
	swap2(facelets, 00, 23)
	swap2(facelets, 34, 45)
	swap2(facelets, 03, 20)
	swap2(facelets, 33, 46)
	swap2(facelets, 05, 18)
	swap2(facelets, 32, 47)
}

func B3(facelets *[48]int) {
	swap4(facelets, 24, 26, 31, 29)
	swap4(facelets, 25, 28, 30, 27)
	swap4(facelets, 00, 45, 23, 34)
	swap4(facelets, 03, 46, 20, 33)
	swap4(facelets, 05, 47, 18, 32)
}

func L1(facelets *[48]int) {
	swap4(facelets, 32, 37, 39, 34)
	swap4(facelets, 33, 35, 38, 36)
	swap4(facelets, 26, 02, 10, 18)
	swap4(facelets, 25, 01, 9, 17)
	swap4(facelets, 24, 00, 8, 16)
}

func L2(facelets *[48]int) {
	swap2(facelets, 32, 39)
	swap2(facelets, 37, 34)
	swap2(facelets, 33, 38)
	swap2(facelets, 35, 36)
	swap2(facelets, 26, 10)
	swap2(facelets, 02, 18)
	swap2(facelets, 25, 9)
	swap2(facelets, 01, 17)
	swap2(facelets, 24, 8)
	swap2(facelets, 00, 16)
}

func L3(facelets *[48]int) {
	swap4(facelets, 32, 34, 39, 37)
	swap4(facelets, 33, 36, 38, 35)
	swap4(facelets, 26, 18, 10, 02)
	swap4(facelets, 25, 17, 9, 01)
	swap4(facelets, 24, 16, 8, 00)
}

func R1(facelets *[48]int) {
	swap4(facelets, 40, 45, 47, 42)
	swap4(facelets, 41, 43, 46, 44)
	swap4(facelets, 05, 29, 21, 13)
	swap4(facelets, 06, 30, 22, 14)
	swap4(facelets, 07, 31, 23, 15)
}

func R2(facelets *[48]int) {
	swap2(facelets, 40, 47)
	swap2(facelets, 45, 42)
	swap2(facelets, 41, 46)
	swap2(facelets, 43, 44)
	swap2(facelets, 05, 21)
	swap2(facelets, 29, 13)
	swap2(facelets, 06, 22)
	swap2(facelets, 30, 14)
	swap2(facelets, 07, 23)
	swap2(facelets, 31, 15)
}

func R3(facelets *[48]int) {
	swap4(facelets, 40, 42, 47, 45)
	swap4(facelets, 41, 44, 46, 43)
	swap4(facelets, 05, 13, 21, 29)
	swap4(facelets, 06, 14, 22, 30)
	swap4(facelets, 07, 15, 23, 31)
}
