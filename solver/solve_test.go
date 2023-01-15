package solver

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	facelets := [48]int{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	PerformAlgorithm(&facelets, "R U R' F' R U R' U' R' F R2 U' R' U'")

	moves := MoveSubset([]string{"U1", "U2", "U3", "F1", "F2", "F3", "R1", "R2", "R3"}) // RUF

	solutions := Solve(facelets, moves, 10, false)
	for _, solution := range solutions {
		fmt.Println(solution)
	}
}
