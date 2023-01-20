package solver

import (
	"testing"
)

func assertByteEqual(t *testing.T, expected, actual byte) {
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func assertCubeEqual(t *testing.T, expected, actual Cube) {
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestTwistCW(t *testing.T) {
	var b byte = 0b00001010
	b = twistCW(b)
	assertByteEqual(t, 0b00011010, b)

	b = twistCW(b)
	assertByteEqual(t, 0b00101010, b)

	b = twistCW(b)
	assertByteEqual(t, 0b00001010, b)
}

func TestTwistCCW(t *testing.T) {
	var b byte = 0b00001010
	b = twistCCW(b)
	assertByteEqual(t, 0b00101010, b)

	b = twistCCW(b)
	assertByteEqual(t, 0b00011010, b)

	b = twistCCW(b)
	assertByteEqual(t, 0b00001010, b)
}

func TestFlip(t *testing.T) {
	var b byte = 0b00001010
	b = flip(b)
	assertByteEqual(t, 0b00011010, b)

	b = flip(b)
	assertByteEqual(t, 0b00001010, b)
}

func TestU(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	U1(&c1)
	U1(&c1)

	U2(&c2)

	assertCubeEqual(t, c1, c2)

	U3(&c1)
	U3(&c1)

	U2(&c2)

	assertCubeEqual(t, c1, c2)
}

func TestF(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	F1(&c1)
	F1(&c1)

	F2(&c2)

	assertCubeEqual(t, c1, c2)

	F3(&c1)
	F3(&c1)

	F2(&c2)

	assertCubeEqual(t, c1, c2)
}

func TestD(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	D1(&c1)
	D1(&c1)

	D2(&c2)

	assertCubeEqual(t, c1, c2)

	D3(&c1)
	D3(&c1)

	D2(&c2)

	assertCubeEqual(t, c1, c2)
}

func TestB(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	B1(&c1)
	B1(&c1)

	B2(&c2)

	assertCubeEqual(t, c1, c2)

	B3(&c1)
	B3(&c1)

	B2(&c2)

	assertCubeEqual(t, c1, c2)
}

func TestL(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	L1(&c1)
	L1(&c1)

	L2(&c2)

	assertCubeEqual(t, c1, c2)

	L3(&c1)
	L3(&c1)

	L2(&c2)

	assertCubeEqual(t, c1, c2)
}

func TestR(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	R1(&c1)
	R1(&c1)

	R2(&c2)

	assertCubeEqual(t, c1, c2)

	R3(&c1)
	R3(&c1)

	R2(&c2)

	assertCubeEqual(t, c1, c2)
}

func TestRU(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	for i := 0; i < 6; i++ {
		PerformAlgorithm(&c1, "R U R' U'")
	}

	assertCubeEqual(t, c1, c2)
}

func TestFL(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	for i := 0; i < 6; i++ {
		PerformAlgorithm(&c1, "F L F' L'")
	}

	assertCubeEqual(t, c1, c2)
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

	assertCubeEqual(t, c1, c2)
}

func TestSome(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	PerformAlgorithm(&c1, "U F D B L R R' L' B' D' F' U'")

	assertCubeEqual(t, c1, c2)
}

func TestMore(t *testing.T) {
	c1 := NewCube()
	c2 := NewCube()

	PerformAlgorithm(&c1, "F R U R' U' F' U2 B U L U' L' B' U2")

	assertCubeEqual(t, c1, c2)
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

	assertCubeEqual(t, c1, c2)
}
