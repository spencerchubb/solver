package solver

import (
	"strings"
)

//          00 08 01
//          09    10
//          02 11 03
// 00 09 02 02 11 03 03 10 02 02 08 00
// 18    12 12    13 13    19 19    18
// 06 15 04 04 14 05 05 16 07 07 17 06
//          04 14 05
//          15    16
//          06 17 07

const U1Num = 0
const U2Num = 1
const U3Num = 2
const F1Num = 3
const F2Num = 4
const F3Num = 5
const D1Num = 6
const D2Num = 7
const D3Num = 8
const B1Num = 9
const B2Num = 10
const B3Num = 11
const L1Num = 12
const L2Num = 13
const L3Num = 14
const R1Num = 15
const R2Num = 16
const R3Num = 17

var inverseMoveNums = []byte{
	U3Num, U2Num, U1Num,
	F3Num, F2Num, F1Num,
	D3Num, D2Num, D1Num,
	B3Num, B2Num, B1Num,
	L3Num, L2Num, L1Num,
	R3Num, R2Num, R1Num,
}

var moveAliases = []byte{
	0x00,
	0x01,
	0x02,
	0x03,
	0x04,
	0x05,
	0x06,
	0x07,
	0x08,
	0x09,
	0x0A,
	0x0B,
	0x0C,
	0x0D,
	0x0E,
	0x0F,
	0x10,
	0x11,
}

var moveNames = []string{
	"U",
	"U2",
	"U'",
	"F",
	"F2",
	"F'",
	"D",
	"D2",
	"D'",
	"B",
	"B2",
	"B'",
	"L",
	"L2",
	"L'",
	"R",
	"R2",
	"R'",
}

var inverseMoveNames = []string{
	"U'",
	"U2",
	"U",
	"F'",
	"F2",
	"F",
	"D'",
	"D2",
	"D",
	"B'",
	"B2",
	"B",
	"L'",
	"L2",
	"L",
	"R'",
	"R2",
	"R",
}

var allMoves = []func(*Cube){
	U1,
	U2,
	U3,
	F1,
	F2,
	F3,
	D1,
	D2,
	D3,
	B1,
	B2,
	B3,
	L1,
	L2,
	L3,
	R1,
	R2,
	R3,
}

func invertMove(move byte) byte {
	return inverseMoveNums[move]
}

func invertMoves(moves []byte) []byte {
	out := make([]byte, len(moves))
	for i, move := range moves {
		out[len(moves)-i-1] = inverseMoveNums[move]
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

// Returns 0xFF if the moves cannot be canceled.
// Returns 0 if the moves perfectly cancel.
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
		if len(out) > 0 && out[len(out)-1] == inverseMoveNums[move] {
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
		}
	}
}

// Twist a corner clockwise
func twistCW(b byte) byte {
	upper := b & 0xF0
	lower := b & 0x0F
	return ((upper+1)%3)<<4 | lower
}

// Twist a corner counter-clockwise
func twistCCW(b byte) byte {
	upper := b & 0xF0
	lower := b & 0x0F
	return ((upper+2)%3)<<4 | lower

}

// Flip an edge
func flip(b byte) byte {
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
