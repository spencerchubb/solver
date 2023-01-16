package solver

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	facelets := SolvedFacelets()
	PerformAlgorithm(&facelets, "R U R' F' R U R' U' R' F R2 U' R' U'")

	moves := MoveSubset([]string{"U1", "U2", "U3", "F1", "F2", "F3", "R1", "R2", "R3"}) // RUF

	solutions := Solve(facelets, moves, 10, false)
	for _, solution := range solutions {
		fmt.Println(solution)
	}
}
