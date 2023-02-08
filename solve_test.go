package solver

import (
	"testing"
)

func TestSolve(t *testing.T) {
	start := NewCube()
	end := NewCube()
	PerformAlgString(&start, "R U R' F' R U R' U' R' F R2 U' R' U'")

	moves := []byte{U1Num, U2Num, U3Num, F1Num, F2Num, F3Num, R1Num, R2Num, R3Num}

	solutions := Solve(start, end, moves, 10, 10_000, false)
	expected := []string{
		"R U2 F' R' F U' F' R F U' R' U'",
		"U R U F' R' F U F' R F U2 R'",
		"R U' F U' R' U' R U F' U2 R' U'",
		"U R U2 F U' R' U R U F' U R'",
		"U R U R' F U F' R F U' F' U' R'",
		"R U F U F' R' F U' F' R U' R' U'",
		"U R2 U R2 F U F' R2 F U' F' U' R2",
		"R2 U F U F' R2 F U' F' R2 U' R2 U'",
		"R U F2 U F' R' F U' F' R F' U' R' U'",
		"U R U F R' F U F' R F U' F2 U' R'",
	}
	assertArraysEqual(t, expected, solutions)
}

func TestOnlyOrientation(t *testing.T) {
	start := NewCube()
	end := NewCube()
	start.SetOnlyOrientation([]int{0, 1, 2, 3, 8, 9, 10, 11})
	end.SetOnlyOrientation([]int{0, 1, 2, 3, 8, 9, 10, 11})
	PerformAlgString(&start, "M' U M U2 M' U M")

	moves := []byte{U1Num, U2Num, U3Num, F1Num, F2Num, F3Num, R1Num, R2Num, R3Num}

	solutions := Solve(start, end, moves, 10, 10_000, false)
	expected := []string{
		"F R' F R2 U R2 F R F' U' F2",
		"F' R' F R2 U R2 F R F U F2",
		"R2 U R F R F2 U F2 R F' R'",
		"F R' F R2 U R2 F R F' U' F2 U'",
		"U F2 U F R' F' R2 U' R2 F' R F'",
		"U F2 U F R' F' R2 U' R2 F' R F' U'",
		"F' R' F R2 U R2 F R F U F2 U'",
		"R2 U R F R F2 U F2 R F' R' U'",
		"U R2 U' R' F R F2 U F2 R F' R",
		"U R2 U' R' F R F2 U F2 R F' R U'",
	}
	assertArraysEqual(t, expected, solutions)
}

func TestDisregard(t *testing.T) {
	start := NewCube()
	end := NewCube()
	start.SetDisregard([]int{0, 1, 2, 3})
	end.SetDisregard([]int{0, 1, 2, 3})
	PerformAlgString(&start, "M' U M U2 M' U M")

	moves := []byte{U1Num, U2Num, U3Num, F1Num, F2Num, F3Num, R1Num, R2Num, R3Num}

	solutions := Solve(start, end, moves, 10, 10_000, false)
	expected := []string{
		"F' U2 F U2 F R' F' R U2",
		"R U2 R' F R' F' R2 U2 R'",
		"U F U F R' F' R2 U' R' F' U'",
		"U F R' F' R U2 R U2 R' U",
		"U R' F R F' U2 F' U2 F U'",
		"U R' F R2 F' U2 F' U2 F R' U'",
		"F R U R2 F R F' U' F' U2",
		"U2 R U2 R' U2 R' F R F' U2",
		"U2 F' U2 F R' F R F2 U2 F",
		"U2 F' R U2 R' U2 R' F2 R F' U2",
	}
	assertArraysEqual(t, expected, solutions)
}
