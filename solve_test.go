package solver

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	start := NewCube()
	end := NewCube()
	PerformAlgorithm(&start, "R U R' F' R U R' U' R' F R2 U' R' U'")

	moves := []byte{U1Num, U2Num, U3Num, F1Num, F2Num, F3Num, R1Num, R2Num, R3Num}

	solutions := Solve(start, end, moves, 10, 10_000, false)
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
	assertArraysEqual(t, expected, solutions)
}

func TestOnlyOrientation(t *testing.T) {
	start := NewCube()
	end := NewCube()
	PerformAlgorithm(&start, "M' U M U2 M' U M")
	start.SetOnlyOrientation([]int{0, 1, 2, 3, 8, 9, 10, 11})
	end.SetOnlyOrientation([]int{0, 1, 2, 3, 8, 9, 10, 11})

	moves := []byte{U1Num, U2Num, U3Num, F1Num, F2Num, F3Num, R1Num, R2Num, R3Num}

	solutions := Solve(start, end, moves, 10, 10_000, false)
	expected := []string{
		"F R' F R2 U R2 F R F' U' F2",
		"F2 U' F' R' F' R2 U' R2 F' R F",
		"R F R' F2 U' F2 R' F' R' U' R2",
		"F R' F R2 U R2 F R F' U' F2 U'",
		"U F R' F R2 U R2 F R F' U' F2 U'",
		"F' R' F R2 U R2 F R F U F2 U'",
		"R2 U R F R F2 U F2 R F' R' U'",
		"R' F R' F2 U' F2 R' F' R U R2 U'",
		"U R2 U' R' F R F2 U F2 R F' R U'",
		"U F2 U F R' F' R2 U' R2 F' R F' U2",
	}
	assertArraysEqual(t, expected, solutions)
}

func TestDisregard(t *testing.T) {
	start := NewCube()
	end := NewCube()
	PerformAlgorithm(&start, "M' U M U2 M' U M")
	start.SetDisregard([]int{0, 1, 2, 3})
	end.SetDisregard([]int{0, 1, 2, 3})

	moves := []byte{U1Num, U2Num, U3Num, F1Num, F2Num, F3Num, R1Num, R2Num, R3Num}

	solutions := Solve(start, end, moves, 10, 10_000, false)
	for _, solution := range solutions {
		fmt.Println(solution)
	}
	// expected := []string{
	// 	"F R' F R2 U R2 F R F' U' F2",
	// 	"F2 U' F' R' F' R2 U' R2 F' R F",
	// 	"R F R' F2 U' F2 R' F' R' U' R2",
	// 	"F R' F R2 U R2 F R F' U' F2 U'",
	// 	"U F R' F R2 U R2 F R F' U' F2 U'",
	// 	"F' R' F R2 U R2 F R F U F2 U'",
	// 	"R2 U R F R F2 U F2 R F' R' U'",
	// 	"R' F R' F2 U' F2 R' F' R U R2 U'",
	// 	"U R2 U' R' F R F2 U F2 R F' R U'",
	// 	"U F2 U F R' F' R2 U' R2 F' R F' U2",
	// }
	// assertArraysEqual(t, expected, solutions)
}
