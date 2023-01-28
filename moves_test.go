package solver

import (
	"testing"
)

func TestInvertMove(t *testing.T) {
	assertEqual(t, U1Num, invertMove(U3Num))
	assertEqual(t, U2Num, invertMove(U2Num))
	assertEqual(t, U3Num, invertMove(U1Num))
	assertEqual(t, F1Num, invertMove(F3Num))
	assertEqual(t, F2Num, invertMove(F2Num))
	assertEqual(t, F3Num, invertMove(F1Num))
	assertEqual(t, D1Num, invertMove(D3Num))
	assertEqual(t, D2Num, invertMove(D2Num))
	assertEqual(t, D3Num, invertMove(D1Num))
	assertEqual(t, B1Num, invertMove(B3Num))
	assertEqual(t, B2Num, invertMove(B2Num))
	assertEqual(t, B3Num, invertMove(B1Num))
	assertEqual(t, L1Num, invertMove(L3Num))
	assertEqual(t, L2Num, invertMove(L2Num))
	assertEqual(t, L3Num, invertMove(L1Num))
	assertEqual(t, R1Num, invertMove(R3Num))
	assertEqual(t, R2Num, invertMove(R2Num))
	assertEqual(t, R3Num, invertMove(R1Num))
}

func TestInvertMoves(t *testing.T) {
	moves := []byte{
		U1Num, U2Num, U3Num,
		F1Num, F2Num, F3Num,
		D1Num, D2Num, D3Num,
		B1Num, B2Num, B3Num,
		L1Num, L2Num, L3Num,
		R1Num, R2Num, R3Num,
	}
	moves = invertMoves(moves)
	expected := []byte{
		R1Num, R2Num, R3Num,
		L1Num, L2Num, L3Num,
		B1Num, B2Num, B3Num,
		D1Num, D2Num, D3Num,
		F1Num, F2Num, F3Num,
		U1Num, U2Num, U3Num,
	}
	for i := range moves {
		assertEqual(t, expected[i], moves[i])
	}
}

func TestCancelPairOfMoves(t *testing.T) {
	assertEqual(t, noCancel, cancelPairOfMoves(U1Num, F1Num))
	assertEqual(t, perfectCancel, cancelPairOfMoves(F1Num, F3Num))
	assertEqual(t, perfectCancel, cancelPairOfMoves(F2Num, F2Num))
	assertEqual(t, F1Num, cancelPairOfMoves(F2Num, F3Num))
	assertEqual(t, F2Num, cancelPairOfMoves(F1Num, F1Num))
	assertEqual(t, F3Num, cancelPairOfMoves(F1Num, F2Num))
}

func TestAlgString(t *testing.T) {
	// Basic case with no cancellations
	forward := []byte{R1Num, U1Num, R3Num, U3Num}
	inverse := []byte{U1Num, R1Num, U3Num, R3Num}
	assertEqual(t, "R U R' U' R U R' U'", algString(forward, inverse))

	// One cancellation
	forward = []byte{R1Num, U1Num, R3Num, U3Num}
	inverse = []byte{F3Num, U1Num, F1Num, U3Num}
	assertEqual(t, "R U R' F' U' F", algString(forward, inverse))

	// Several cancellations
	forward = []byte{R1Num, U1Num, R3Num, U1Num, R1Num, U2Num, R3Num}
	inverse = []byte{R1Num, U2Num, R3Num, U3Num, R1Num, U3Num, R3Num}
	assertEqual(t, "R U R' U R U' R' U R U2 R'", algString(forward, inverse))

	// U2 and U' simplify to U
	forward = []byte{U1Num, R1Num, U1Num, F3Num, R3Num, F1Num, U2Num}
	inverse = []byte{R1Num, U2Num, F3Num, R3Num, F1Num, U1Num}
	assertEqual(t, "U R U F' R' F U F' R F U2 R'", algString(forward, inverse))
}

func TestTwistCW(t *testing.T) {
	var b byte = 0b00001010
	b = twistCW(b)
	assertEqual(t, 0b00011010, b)

	b = twistCW(b)
	assertEqual(t, 0b00101010, b)

	b = twistCW(b)
	assertEqual(t, 0b00001010, b)
}

func TestTwistCCW(t *testing.T) {
	var b byte = 0b00001010
	b = twistCCW(b)
	assertEqual(t, 0b00101010, b)

	b = twistCCW(b)
	assertEqual(t, 0b00011010, b)

	b = twistCCW(b)
	assertEqual(t, 0b00001010, b)
}

func TestFlip(t *testing.T) {
	var b byte = 0b00001010
	b = flip(b)
	assertEqual(t, 0b00011010, b)

	b = flip(b)
	assertEqual(t, 0b00001010, b)
}

func TestU(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	U1(&c1)
	U1(&c1)

	U2(&c2)

	assertEqual(t, c1, c2)

	U3(&c1)
	U3(&c1)

	U2(&c2)

	assertEqual(t, c1, c2)
}

func TestF(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	F1(&c1)
	F1(&c1)

	F2(&c2)

	assertEqual(t, c1, c2)

	F3(&c1)
	F3(&c1)

	F2(&c2)

	assertEqual(t, c1, c2)
}

func TestD(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	D1(&c1)
	D1(&c1)

	D2(&c2)

	assertEqual(t, c1, c2)

	D3(&c1)
	D3(&c1)

	D2(&c2)

	assertEqual(t, c1, c2)
}

func TestB(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	B1(&c1)
	B1(&c1)

	B2(&c2)

	assertEqual(t, c1, c2)

	B3(&c1)
	B3(&c1)

	B2(&c2)

	assertEqual(t, c1, c2)
}

func TestL(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	L1(&c1)
	L1(&c1)

	L2(&c2)

	assertEqual(t, c1, c2)

	L3(&c1)
	L3(&c1)

	L2(&c2)

	assertEqual(t, c1, c2)
}

func TestR(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	R1(&c1)
	R1(&c1)

	R2(&c2)

	assertEqual(t, c1, c2)

	R3(&c1)
	R3(&c1)

	R2(&c2)

	assertEqual(t, c1, c2)
}

func TestRU(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	for i := 0; i < 6; i++ {
		PerformAlgorithm(&c1, "R U R' U'")
	}

	assertEqual(t, c1, c2)
}

func TestFL(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	for i := 0; i < 6; i++ {
		PerformAlgorithm(&c1, "F L F' L'")
	}

	assertEqual(t, c1, c2)
}

func TestBD(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	for i := 0; i < 6; i++ {
		B1(&c1)
		D1(&c1)
		B3(&c1)
		D3(&c1)
	}

	assertEqual(t, c1, c2)
}

func TestSome(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	PerformAlgorithm(&c1, "U F D B L R R' L' B' D' F' U'")

	assertEqual(t, c1, c2)
}

func TestMore(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	PerformAlgorithm(&c1, "F R U R' U' F' U2 B U L U' L' B' U2")

	assertEqual(t, c1, c2)
}

func TestRandom(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	// Scramble
	PerformAlgorithm(&c1, "R F B2 U2 B' L' D F' B L2 D2 R2 U R' D' F2 U' L")

	// Cross
	PerformAlgorithm(&c1, "R D2 L2")

	// Pair 1
	PerformAlgorithm(&c1, "D F D' F' D F' D2 F")

	// Pair 2
	PerformAlgorithm(&c1, "D2 R' D R")

	// Pair 3+4
	PerformAlgorithm(&c1, "L' D' L D' B D B2 D' B D' R D R'")

	// OLL
	PerformAlgorithm(&c1, "B' D' L' D L B")

	// PLL
	PerformAlgorithm(&c1, "R D R' D R D R' B' R D R' D' R' B R2 D' R' D2 R D' R'")

	assertEqual(t, c1, c2)
}

func TestSlices(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	PerformAlgorithm(&c1, "M E S M' S2 E' M2 S' E2 S M2 S E2")

	assertEqual(t, c1, c2)
}
