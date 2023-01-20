package solver

import (
	"fmt"
)

func lastMoveSameFace(moves []byte, move byte) bool {
	// If there are no moves in the array, there is no last move
	if len(moves) == 0 {
		return false
	}

	lastMove := moves[len(moves)-1]

	// If the moves are on the same face, then these should be equal.
	// e.g. U1 is associated with 0x00, and U3 is associated with 0x02.
	// With integer division, 0x00/3 is 0, and 0x02/3 is 0.
	return lastMove/3 == move/3
}

func Solve(cube Cube, moves []int, maxSolutions int, log bool) []string {
	depth := 0
	inverseDepth := 0

	visited := initVisited()
	inverseVisited := initVisited()

	// It is faster with *Node instead of Node
	queue := []*Node{{cube, &[]byte{}}}
	inverseQueue := []*Node{{NewCube(), &[]byte{}}}

	var solutions []string
	for loc := 0; ; loc++ {
		node := queue[0]
		queue = queue[1:]

		inverseNode := inverseQueue[0]
		inverseQueue = inverseQueue[1:]

		algs := get(visited, inverseNode.cube)
		for _, alg := range algs {
			algStr := algString(alg, *inverseNode.moves)
			if log {
				fmt.Println(algStr)
			}
			solutions = append(solutions, algStr)
			if len(solutions) >= maxSolutions {
				return solutions
			}
		}

		inverseAlgs := get(inverseVisited, node.cube)
		for _, inverseAlg := range inverseAlgs {
			algStr := algString(*node.moves, inverseAlg)
			if log {
				fmt.Println(algStr)
			}
			solutions = append(solutions, algStr)
			if len(solutions) >= maxSolutions {
				return solutions
			}
		}

		if log && len(*node.moves) > depth {
			depth = len(*node.moves)
			fmt.Printf("Searching depth: %d\n", depth)
		}

		if log && len(*inverseNode.moves) > inverseDepth {
			inverseDepth = len(*inverseNode.moves)
			fmt.Printf("Searching inverse depth: %d\n", inverseDepth)
		}

		for _, move := range moves {
			if !lastMoveSameFace(*node.moves, moveAliases[move]) {
				cpy := node.cube
				allMoves[move](&cpy)
				newMoves := appendImmutable(*node.moves, moveAliases[move])
				queue = append(queue, &Node{cpy, &newMoves})

				add(visited, cpy, newMoves)
			}
			if !lastMoveSameFace(*inverseNode.moves, moveAliases[move]) {
				cpy := inverseNode.cube
				allMoves[move](&cpy)
				newMoves := appendImmutable(*inverseNode.moves, moveAliases[move])
				inverseQueue = append(inverseQueue, &Node{cpy, &newMoves})

				add(inverseVisited, cpy, newMoves)
			}
		}
	}
}
