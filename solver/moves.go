package solver

import (
	"bytes"
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

func algString(forward []byte, inverse []byte) string {
	buff := bytes.NewBufferString("")
	inverseHasMoves := len(inverse) > 0
	for i := 0; i < len(forward); i++ {
		move := forward[i]
		name := moveNames[move]
		buff.WriteString(name)

		// Don't add a space after the last move if there are no inverse moves
		if i < len(forward)-1 || inverseHasMoves {
			buff.WriteByte(' ')
		}
	}
	for i := len(inverse) - 1; i >= 0; i-- {
		move := inverse[i]
		name := inverseMoveNames[move]
		buff.WriteString(name)

		if i > 0 {
			buff.WriteByte(' ')
		}
	}
	return buff.String()
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
