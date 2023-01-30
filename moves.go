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
const X1Num, X2Num, X3Num byte = 27, 28, 29
const Y1Num, Y2Num, Y3Num byte = 30, 31, 32
const Z1Num, Z2Num, Z3Num byte = 33, 34, 35

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
	"x", "x2", "x'",
	"y", "y2", "y'",
	"z", "z2", "z'",
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
	X1, X2, X3,
	Y1, Y2, Y3,
	Z1, Z2, Z3,
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

func invertMoves(moves []byte) []byte {
	out := make([]byte, len(moves))
	for i, move := range moves {
		out[len(moves)-i-1] = invertMove(move)
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

func cancelMoves(moves []byte) []byte {
	out := make([]byte, 0, len(moves))
	for _, move := range moves {
		if len(out) > 0 && out[len(out)-1] == invertMove(move) {
			out = out[:len(out)-1]
		} else {
			out = append(out, move)
		}
	}
	return out
}

func algString(forward []byte, inverse []byte) string {
	inverted := invertMoves(inverse)
	combined := appendImmutable(forward, inverted...)
	cleaned := make([]byte, 0, len(combined))
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

func PerformAlgorithm(cube *Cube, algorithm string) {
	for _, move := range strings.Split(algorithm, " ") {
		switch move {
		case "U":
			U1(cube)
		case "U2":
			U2(cube)
		case "U'":
			U3(cube)
		case "F":
			F1(cube)
		case "F2":
			F2(cube)
		case "F'":
			F3(cube)
		case "D":
			D1(cube)
		case "D2":
			D2(cube)
		case "D'":
			D3(cube)
		case "B":
			B1(cube)
		case "B2":
			B2(cube)
		case "B'":
			B3(cube)
		case "L":
			L1(cube)
		case "L2":
			L2(cube)
		case "L'":
			L3(cube)
		case "R":
			R1(cube)
		case "R2":
			R2(cube)
		case "R'":
			R3(cube)
		case "M":
			M1(cube)
		case "M2":
			M2(cube)
		case "M'":
			M3(cube)
		case "E":
			E1(cube)
		case "E2":
			E2(cube)
		case "E'":
			E3(cube)
		case "S":
			S1(cube)
		case "S2":
			S2(cube)
		case "S'":
			S3(cube)
		case "x":
			X1(cube)
		case "x2":
			X2(cube)
		case "x'":
			X3(cube)
		case "y":
			Y1(cube)
		case "y2":
			Y2(cube)
		case "y'":
			Y3(cube)
		case "z":
			Z1(cube)
		case "z2":
			Z2(cube)
		case "z'":
			Z3(cube)
		default:
			fmt.Printf("Unknown move: %s\n", move)
		}
	}
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

func X1(c *Cube) {
	R1(c)
	M3(c)
	L3(c)
}

func X2(c *Cube) {
	R2(c)
	M2(c)
	L2(c)
}

func X3(c *Cube) {
	R3(c)
	M1(c)
	L1(c)
}

func Y1(c *Cube) {
	U1(c)
	E3(c)
	D3(c)
}

func Y2(c *Cube) {
	U2(c)
	E2(c)
	D2(c)
}

func Y3(c *Cube) {
	U3(c)
	E1(c)
	D1(c)
}

func Z1(c *Cube) {
	F1(c)
	S1(c)
	B3(c)
}

func Z2(c *Cube) {
	F2(c)
	S2(c)
	B2(c)
}

func Z3(c *Cube) {
	F3(c)
	S3(c)
	B1(c)
}
