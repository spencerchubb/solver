package solver

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	cube := NewCube()
	PerformAlgorithm(&cube, "R U R' F' R U R' U' R' F R2 U' R' U'")

	moves := []int{0, 1, 2, 3, 4, 5, 15, 16, 17}

	solutions := Solve(cube, moves, 10, 10_000, false)
	for _, solution := range solutions {
		fmt.Println(solution)
	}
}
