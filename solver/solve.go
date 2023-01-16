package solver

import (
	"fmt"
)

func lastMoveSameFace(moves []string, move string) bool {
	// If there are no moves in the array, there is no last move
	if len(moves) == 0 {
		return false
	}

	lastMove := moves[len(moves)-1]

	// Compare the first character
	// e.g. if the moves are R1 and R2, then those are on the same face
	return lastMove[0] == move[0]
}

func Solve(facelets Facelets, moves []Move, maxSolutions int, log bool) []string {
	depth := 0
	visited := initVisited()
	queue := []Node{{facelets, []string{}}}

	inverseDepth := 0
	inverseVisited := initVisited()
	inverseQueue := []Node{{SolvedFacelets(), []string{}}}

	var solutions []string
	for loc := 0; ; loc++ {
		node := queue[0]
		queue = queue[1:]

		inverseNode := inverseQueue[0]
		inverseQueue = inverseQueue[1:]

		algs := get(visited, inverseNode.facelets)
		for _, alg := range algs {
			algStr := algString(alg, inverseNode.moves)
			if log {
				fmt.Println(algStr)
			}
			solutions = append(solutions, algStr)
			if len(solutions) >= maxSolutions {
				return solutions
			}
		}

		inverseAlgs := get(inverseVisited, node.facelets)
		for _, inverseAlg := range inverseAlgs {
			algStr := algString(node.moves, inverseAlg)
			if log {
				fmt.Println(algStr)
			}
			solutions = append(solutions, algStr)
			if len(solutions) >= maxSolutions {
				return solutions
			}
		}

		if log && len(node.moves) > depth {
			depth = len(node.moves)
			fmt.Printf("Searching depth: %d\n", depth)
		}

		if log && len(inverseNode.moves) > inverseDepth {
			inverseDepth = len(inverseNode.moves)
			fmt.Printf("Searching inverse depth: %d\n", inverseDepth)
		}

		for _, move := range moves {
			if !lastMoveSameFace(node.moves, move.name) {
				cpy := node.facelets
				move.proc(&cpy)
				newMoves := appendImmutable(node.moves, move.name)
				queue = append(queue, Node{cpy, newMoves})

				add(visited, cpy, newMoves)
			}
			if !lastMoveSameFace(inverseNode.moves, move.name) {
				cpy := inverseNode.facelets
				move.proc(&cpy)
				newMoves := appendImmutable(inverseNode.moves, move.name)
				inverseQueue = append(inverseQueue, Node{cpy, newMoves})

				add(inverseVisited, cpy, newMoves)
			}
		}
	}
}
