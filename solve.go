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

func solve(facelets [48]int, moves []Move) {
	depth := 0
	visited := initVisited()
	queue := []Node{{facelets, []string{}}}

	inverseDepth := 0
	inverseVisited := initVisited()
	inverseQueue := []Node{{solvedFacelets(), []string{}}}

	solutions := 0
	for loc := 0; ; loc++ {
		node := queue[0]
		queue = queue[1:]

		inverseNode := inverseQueue[0]
		inverseQueue = inverseQueue[1:]

		algs := get(visited, inverseNode.facelets)
		for _, alg := range algs {
			solutions++
			// fmt.Printf("Inverse solution #%d at location %d!\n", solutions, loc)
			// fmt.Printf("Forward: %v, Inverse: %v\n", alg, inverseNode.moves)
			fmt.Println(algString(alg, inverseNode.moves))
		}

		inverseAlgs := get(inverseVisited, node.facelets)
		for _, inverseAlg := range inverseAlgs {
			solutions++
			// fmt.Printf("Solution #%d at location %d!\n", solutions, loc)
			// fmt.Printf("Forward: %v, Inverse: %v\n", node.moves, inverseAlg)
			fmt.Println(algString(node.moves, inverseAlg))
		}

		if len(node.moves) > depth {
			depth = len(node.moves)
			fmt.Printf("Searching depth: %d\n", depth)
		}

		if len(inverseNode.moves) > inverseDepth {
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
