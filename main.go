package solver

func main() {
	facelets := [48]int{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	performAlgorithm(&facelets, "F R U' R' U' R U R' F' R U R' U' R' F R F'")

	moves := moveSubset([]string{"U1", "U2", "U3", "F1", "F2", "F3", "R1", "R2", "R3"}) // RUF

	solve(facelets, moves)
}
