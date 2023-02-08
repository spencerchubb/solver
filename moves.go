package solver

import (
	"fmt"
	"strings"
)

//          00 08 01
//          09 20 10
//          02 11 03
// 00 09 02 02 11 03 03 10 02 02 08 00
// 18 24 12 12 21 13 13 25 19 19 23 18
// 06 15 04 04 14 05 05 16 07 07 17 06
//          04 14 05
//          15 22 16
//          06 17 07

const U1Num, U2Num, U3Num byte = 0, 1, 2
const F1Num, F2Num, F3Num byte = 3, 4, 5
const D1Num, D2Num, D3Num byte = 6, 7, 8
const B1Num, B2Num, B3Num byte = 9, 10, 11
const L1Num, L2Num, L3Num byte = 12, 13, 14
const R1Num, R2Num, R3Num byte = 15, 16, 17
const M1Num, M2Num, M3Num byte = 18, 19, 20
const E1Num, E2Num, E3Num byte = 21, 22, 23
const S1Num, S2Num, S3Num byte = 24, 25, 26

var moveNames = []string{
	"U", "U2", "U'",
	"F", "F2", "F'",
	"D", "D2", "D'",
	"B", "B2", "B'",
	"L", "L2", "L'",
	"R", "R2", "R'",
	"M", "M2", "M'",
	"E", "E2", "E'",
	"S", "S2", "S'",
}

var moveFuncs = []func(*Cube){
	U1, U2, U3,
	F1, F2, F3,
	D1, D2, D3,
	B1, B2, B3,
	L1, L2, L3,
	R1, R2, R3,
	M1, M2, M3,
	E1, E2, E3,
	S1, S2, S3,
}

func invertMove(move byte) byte {
	switch move % 3 {
	case 0:
		return move + 2
	case 1:
		return move
	case 2:
		return move - 2
	}
	panic("unreachable")
}

func InvertAlgorithm(alg Algorithm) Algorithm {
	out := make(Algorithm, len(alg))
	for i, move := range alg {
		out[len(alg)-i-1] = invertMove(move)
	}
	return out
}

const noCancel = 0xFF
const perfectCancel = 0xFE

var equivalences = [][]byte{
	{U2Num, U3Num, perfectCancel},
	{U3Num, perfectCancel, U1Num},
	{perfectCancel, U1Num, U2Num},
	{F2Num, F3Num, perfectCancel},
	{F3Num, perfectCancel, F1Num},
	{perfectCancel, F1Num, F2Num},
	{D2Num, D3Num, perfectCancel},
	{D3Num, perfectCancel, D1Num},
	{perfectCancel, D1Num, D2Num},
	{B2Num, B3Num, perfectCancel},
	{B3Num, perfectCancel, B1Num},
	{perfectCancel, B1Num, B2Num},
	{L2Num, L3Num, perfectCancel},
	{L3Num, perfectCancel, L1Num},
	{perfectCancel, L1Num, L2Num},
	{R2Num, R3Num, perfectCancel},
	{R3Num, perfectCancel, R1Num},
	{perfectCancel, R1Num, R2Num},
}

// Returns noCancel if the moves cannot be canceled.
// Returns perfectCancel if the moves perfectly cancel.
// Returns the move that remains if the moves can be simplified.
// For example, U1 and U2 can simplify to U3.
func cancelPairOfMoves(m1, m2 byte) byte {
	face1 := m1 / 3
	face2 := m2 / 3
	if face1 != face2 {
		return noCancel
	}
	return equivalences[m1][m2%3]
}

func algString(forward Algorithm, inverse Algorithm) string {
	inverted := InvertAlgorithm(inverse)
	combined := appendImmutable(forward, inverted...)
	cleaned := make(Algorithm, 0, len(combined))
	for _, move := range combined {
		if len(cleaned) > 0 {
			cancel := cancelPairOfMoves(cleaned[len(cleaned)-1], move)
			if cancel == noCancel {
				cleaned = append(cleaned, move)
			} else if cancel == perfectCancel {
				cleaned = cleaned[:len(cleaned)-1]
			} else {
				cleaned[len(cleaned)-1] = cancel
			}
		} else {
			cleaned = append(cleaned, move)
		}
	}
	movesAsStrings := make([]string, len(cleaned))
	for i, m := range cleaned {
		movesAsStrings[i] = moveNames[m]
	}
	return strings.Join(movesAsStrings, " ")
}

func StringToAlg(algStr string) Algorithm {
	moveStrings := strings.Split(algStr, " ")
	moves := make(Algorithm, 0, len(moveStrings))
	for _, move := range moveStrings {
		switch move {
		case "U":
			moves = append(moves, U1Num)
		case "U2":
			moves = append(moves, U2Num)
		case "U'":
			moves = append(moves, U3Num)
		case "F":
			moves = append(moves, F1Num)
		case "F2":
			moves = append(moves, F2Num)
		case "F'":
			moves = append(moves, F3Num)
		case "D":
			moves = append(moves, D1Num)
		case "D2":
			moves = append(moves, D2Num)
		case "D'":
			moves = append(moves, D3Num)
		case "B":
			moves = append(moves, B1Num)
		case "B2":
			moves = append(moves, B2Num)
		case "B'":
			moves = append(moves, B3Num)
		case "L":
			moves = append(moves, L1Num)
		case "L2":
			moves = append(moves, L2Num)
		case "L'":
			moves = append(moves, L3Num)
		case "R":
			moves = append(moves, R1Num)
		case "R2":
			moves = append(moves, R2Num)
		case "R'":
			moves = append(moves, R3Num)
		case "M":
			moves = append(moves, M1Num)
		case "M2":
			moves = append(moves, M2Num)
		case "M'":
			moves = append(moves, M3Num)
		case "E":
			moves = append(moves, E1Num)
		case "E2":
			moves = append(moves, E2Num)
		case "E'":
			moves = append(moves, E3Num)
		case "S":
			moves = append(moves, S1Num)
		case "S2":
			moves = append(moves, S2Num)
		case "S'":
			moves = append(moves, S3Num)
		case "x":
			moves = append(moves, R1Num, M3Num, L3Num)
		case "x2":
			moves = append(moves, R2Num, M2Num, L2Num)
		case "x'":
			moves = append(moves, R3Num, M1Num, L1Num)
		case "y":
			moves = append(moves, U1Num, E3Num, D3Num)
		case "y2":
			moves = append(moves, U2Num, E2Num, D2Num)
		case "y'":
			moves = append(moves, U3Num, E1Num, D1Num)
		case "z":
			moves = append(moves, F1Num, S1Num, B3Num)
		case "z2":
			moves = append(moves, F2Num, S2Num, B2Num)
		case "z'":
			moves = append(moves, F3Num, S3Num, B1Num)
		case "l":
			moves = append(moves, L1Num, M1Num)
		case "l2":
			moves = append(moves, L2Num, M2Num)
		case "l'":
			moves = append(moves, L3Num, M3Num)
		case "r":
			moves = append(moves, R1Num, M3Num)
		case "r2":
			moves = append(moves, R2Num, M2Num)
		case "r'":
			moves = append(moves, R3Num, M1Num)
		default:
			fmt.Printf("Unknown move: %s\n", move)
		}
	}
	return moves
}

func PerformAlgorithm(cube *Cube, alg Algorithm) {
	for _, move := range alg {
		moveFunc := moveFuncs[move]
		moveFunc(cube)
	}
}

func PerformAlgString(cube *Cube, algStr string) {
	alg := StringToAlg(algStr)
	PerformAlgorithm(cube, alg)
}

// Twist a corner clockwise
func twistCW(b byte) byte {
	if b == disregard {
		return b
	}
	upper := b & 0xF0
	lower := b & 0x0F
	return ((upper+1)%3)<<4 | lower
}

// Twist a corner counter-clockwise
func twistCCW(b byte) byte {
	if b == disregard {
		return b
	}
	upper := b & 0xF0
	lower := b & 0x0F
	return ((upper+2)%3)<<4 | lower

}

// Flip an edge
func flip(b byte) byte {
	if b == disregard {
		return b
	}
	return b ^ 0b00010000
}

func U1(c *Cube) {
	temp := c[0]
	c[0] = c[2]
	c[2] = c[3]
	c[3] = c[1]
	c[1] = temp

	temp = c[8]
	c[8] = c[9]
	c[9] = c[11]
	c[11] = c[10]
	c[10] = temp
}

func U2(c *Cube) {
	temp := c[0]
	c[0] = c[3]
	c[3] = temp

	temp = c[1]
	c[1] = c[2]
	c[2] = temp

	temp = c[8]
	c[8] = c[11]
	c[11] = temp

	temp = c[9]
	c[9] = c[10]
	c[10] = temp
}

func U3(c *Cube) {
	temp := c[0]
	c[0] = c[1]
	c[1] = c[3]
	c[3] = c[2]
	c[2] = temp

	temp = c[8]
	c[8] = c[10]
	c[10] = c[11]
	c[11] = c[9]
	c[9] = temp
}

func F1(c *Cube) {
	temp := c[2]
	c[2] = twistCCW(c[4])
	c[4] = twistCW(c[5])
	c[5] = twistCCW(c[3])
	c[3] = twistCW(temp)

	temp = c[11]
	c[11] = flip(c[12])
	c[12] = flip(c[14])
	c[14] = flip(c[13])
	c[13] = flip(temp)
}

func F2(c *Cube) {
	temp := c[2]
	c[2] = c[5]
	c[5] = temp

	temp = c[3]
	c[3] = c[4]
	c[4] = temp

	temp = c[11]
	c[11] = c[14]
	c[14] = temp

	temp = c[12]
	c[12] = c[13]
	c[13] = temp
}

func F3(c *Cube) {
	temp := c[2]
	c[2] = twistCCW(c[3])
	c[3] = twistCW(c[5])
	c[5] = twistCCW(c[4])
	c[4] = twistCW(temp)

	temp = c[11]
	c[11] = flip(c[13])
	c[13] = flip(c[14])
	c[14] = flip(c[12])
	c[12] = flip(temp)
}

func D1(c *Cube) {
	temp := c[4]
	c[4] = c[6]
	c[6] = c[7]
	c[7] = c[5]
	c[5] = temp

	temp = c[14]
	c[14] = c[15]
	c[15] = c[17]
	c[17] = c[16]
	c[16] = temp
}

func D2(c *Cube) {
	temp := c[4]
	c[4] = c[7]
	c[7] = temp

	temp = c[5]
	c[5] = c[6]
	c[6] = temp

	temp = c[14]
	c[14] = c[17]
	c[17] = temp

	temp = c[15]
	c[15] = c[16]
	c[16] = temp
}

func D3(c *Cube) {
	temp := c[4]
	c[4] = c[5]
	c[5] = c[7]
	c[7] = c[6]
	c[6] = temp

	temp = c[14]
	c[14] = c[16]
	c[16] = c[17]
	c[17] = c[15]
	c[15] = temp
}

func B1(c *Cube) {
	temp := c[0]
	c[0] = twistCW(c[1])
	c[1] = twistCCW(c[7])
	c[7] = twistCW(c[6])
	c[6] = twistCCW(temp)

	temp = c[8]
	c[8] = flip(c[19])
	c[19] = flip(c[17])
	c[17] = flip(c[18])
	c[18] = flip(temp)
}

func B2(c *Cube) {
	temp := c[0]
	c[0] = c[7]
	c[7] = temp

	temp = c[1]
	c[1] = c[6]
	c[6] = temp

	temp = c[8]
	c[8] = c[17]
	c[17] = temp

	temp = c[19]
	c[19] = c[18]
	c[18] = temp
}

func B3(c *Cube) {
	temp := c[0]
	c[0] = twistCW(c[6])
	c[6] = twistCCW(c[7])
	c[7] = twistCW(c[1])
	c[1] = twistCCW(temp)

	temp = c[8]
	c[8] = flip(c[18])
	c[18] = flip(c[17])
	c[17] = flip(c[19])
	c[19] = flip(temp)
}

func L1(c *Cube) {
	temp := c[0]
	c[0] = twistCCW(c[6])
	c[6] = twistCW(c[4])
	c[4] = twistCCW(c[2])
	c[2] = twistCW(temp)

	temp = c[9]
	c[9] = c[18]
	c[18] = c[15]
	c[15] = c[12]
	c[12] = temp
}

func L2(c *Cube) {
	temp := c[0]
	c[0] = c[4]
	c[4] = temp

	temp = c[2]
	c[2] = c[6]
	c[6] = temp

	temp = c[9]
	c[9] = c[15]
	c[15] = temp

	temp = c[12]
	c[12] = c[18]
	c[18] = temp
}

func L3(c *Cube) {
	temp := c[0]
	c[0] = twistCCW(c[2])
	c[2] = twistCW(c[4])
	c[4] = twistCCW(c[6])
	c[6] = twistCW(temp)

	temp = c[9]
	c[9] = c[12]
	c[12] = c[15]
	c[15] = c[18]
	c[18] = temp
}

func R1(c *Cube) {
	temp := c[1]
	c[1] = twistCW(c[3])
	c[3] = twistCCW(c[5])
	c[5] = twistCW(c[7])
	c[7] = twistCCW(temp)

	temp = c[10]
	c[10] = c[13]
	c[13] = c[16]
	c[16] = c[19]
	c[19] = temp
}

func R2(c *Cube) {
	temp := c[1]
	c[1] = c[5]
	c[5] = temp

	temp = c[3]
	c[3] = c[7]
	c[7] = temp

	temp = c[10]
	c[10] = c[16]
	c[16] = temp

	temp = c[13]
	c[13] = c[19]
	c[19] = temp
}

func R3(c *Cube) {
	temp := c[1]
	c[1] = twistCW(c[7])
	c[7] = twistCCW(c[5])
	c[5] = twistCW(c[3])
	c[3] = twistCCW(temp)

	temp = c[10]
	c[10] = c[19]
	c[19] = c[16]
	c[16] = c[13]
	c[13] = temp
}

func M1(c *Cube) {
	temp := c[8]
	c[8] = flip(c[17])
	c[17] = flip(c[14])
	c[14] = flip(c[11])
	c[11] = flip(temp)

	temp = c[20]
	c[20] = c[23]
	c[23] = c[22]
	c[22] = c[21]
	c[21] = temp
}

func M2(c *Cube) {
	temp := c[8]
	c[8] = c[14]
	c[14] = temp

	temp = c[11]
	c[11] = c[17]
	c[17] = temp

	temp = c[20]
	c[20] = c[22]
	c[22] = temp

	temp = c[21]
	c[21] = c[23]
	c[23] = temp
}

func M3(c *Cube) {
	temp := c[8]
	c[8] = flip(c[11])
	c[11] = flip(c[14])
	c[14] = flip(c[17])
	c[17] = flip(temp)

	temp = c[20]
	c[20] = c[21]
	c[21] = c[22]
	c[22] = c[23]
	c[23] = temp
}

func E1(c *Cube) {
	temp := c[12]
	c[12] = flip(c[18])
	c[18] = flip(c[19])
	c[19] = flip(c[13])
	c[13] = flip(temp)

	temp = c[21]
	c[21] = c[24]
	c[24] = c[23]
	c[23] = c[25]
	c[25] = temp
}

func E2(c *Cube) {
	temp := c[12]
	c[12] = c[19]
	c[19] = temp

	temp = c[13]
	c[13] = c[18]
	c[18] = temp

	temp = c[21]
	c[21] = c[23]
	c[23] = temp

	temp = c[24]
	c[24] = c[25]
	c[25] = temp
}

func E3(c *Cube) {
	temp := c[12]
	c[12] = flip(c[13])
	c[13] = flip(c[19])
	c[19] = flip(c[18])
	c[18] = flip(temp)

	temp = c[21]
	c[21] = c[25]
	c[25] = c[23]
	c[23] = c[24]
	c[24] = temp
}

func S1(c *Cube) {
	temp := c[9]
	c[9] = flip(c[15])
	c[15] = flip(c[16])
	c[16] = flip(c[10])
	c[10] = flip(temp)

	temp = c[20]
	c[20] = c[24]
	c[24] = c[22]
	c[22] = c[25]
	c[25] = temp
}

func S2(c *Cube) {
	temp := c[9]
	c[9] = c[16]
	c[16] = temp

	temp = c[10]
	c[10] = c[15]
	c[15] = temp

	temp = c[20]
	c[20] = c[22]
	c[22] = temp

	temp = c[24]
	c[24] = c[25]
	c[25] = temp
}

func S3(c *Cube) {
	temp := c[9]
	c[9] = flip(c[10])
	c[10] = flip(c[16])
	c[16] = flip(c[15])
	c[15] = flip(temp)

	temp = c[20]
	c[20] = c[25]
	c[25] = c[22]
	c[22] = c[24]
	c[24] = temp
}
