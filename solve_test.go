package solver

import (
	"testing"
)

func TestSolve(t *testing.T) {
	cube := NewCube()
	PerformAlgorithm(&cube, "R U R' F' R U R' U' R' F R2 U' R' U'")

	moves := []byte{0, 1, 2, 3, 4, 5, 15, 16, 17}

	solutions := Solve(cube, moves, 10, 10_000, false)
	expected := []string{
		"R U2 F' R' F U' F' R F U' R' U'",
		"R U' F U' R' U' R U F' U2 R' U'",
		"U R U R' F U F' R F U' F' U' R'",
		"U R U F' R' F U F' R F U2 R'",
		"U R U2 F U' R' U R U F' U R'",
		"U R2 U R2 F U F' R2 F U' F' U' R2",
		"R U F2 U F' R' F U' F' R F' U' R' U'",
		"R U F' U F' R' F U' F' R F2 U' R' U'",
		"R U R' F' R U R' U' R' F R2 U' R' U'",
		"R U F U F' R' F U' F' R U' R' U'",
	}
	assertArraysEqual(t, solutions, expected)
}
