package solver

import (
	"fmt"
	"time"
)

var oppositeFaces = []byte{2, 3, 0, 1, 5, 4}

func sameFace(moves []byte, move byte) bool {
	// If there are no moves in the array, there is no last move.
	if len(moves) == 0 {
		return false
	}

	lastMove := moves[len(moves)-1]

	// If the moves are on the same face, then these should be equal.
	// e.g. U1 is associated with 0x00, and U3 is associated with 0x02.
	// With integer division, 0x00/3 is 0, and 0x02/3 is 0.
	if lastMove/3 == move/3 {
		return true
	}

	// If there is only one move in the array, there is no second-to-last move.
	if len(moves) == 1 {
		return false
	}

	secondLastMove := moves[len(moves)-2]

	// Returns true if the last move is on the opposite face AND the second-to-last move is on the same face.
	return move/3 == oppositeFaces[lastMove/3] && move/3 == secondLastMove/3
}

func Solve(start Cube, end Cube, moves []byte, maxSolutions int, maxMs int64, log bool) []string {
	depth := 0
	inverseDepth := 0

	visited := initVisited()
	inverseVisited := initVisited()

	// It is faster with *Node instead of Node
	queue := []*Node{{start, &[]byte{}}}
	inverseQueue := []*Node{{end, &[]byte{}}}

	solutionExists := make(map[string]bool)
	var solutions []string
	var startMs = time.Now().UnixMilli()
	for loc := 0; ; loc++ {
		elapsedMs := time.Now().UnixMilli() - startMs
		if elapsedMs > maxMs {
			return solutions
		}

		node := queue[0]
		queue = queue[1:]

		inverseNode := inverseQueue[0]
		inverseQueue = inverseQueue[1:]

		algs := get(visited, inverseNode.cube)
		for _, alg := range algs {
			algStr := algString(alg, *inverseNode.moves)
			if solutionExists[algStr] {
				continue
			}
			solutionExists[algStr] = true
			if log {
				fmt.Println(algStr)
			}
			solutions = append(solutions, algStr)
		}

		algs = get(inverseVisited, node.cube)
		for _, alg := range algs {
			algStr := algString(*node.moves, alg)
			if solutionExists[algStr] {
				continue
			}
			solutionExists[algStr] = true
			if log {
				fmt.Println(algStr)
			}
			solutions = append(solutions, algStr)
		}

		if len(solutions) >= maxSolutions {
			return solutions
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
			goToChild(&queue, node, visited, move)
			goToChild(&inverseQueue, inverseNode, inverseVisited, move)
		}
	}
}

func goToChild(queue *[]*Node, node *Node, visited Visited, move byte) {
	if !sameFace(*node.moves, move) {
		cpy := node.cube
		moveFuncs[move](&cpy)
		newMoves := appendImmutable(*node.moves, move)
		*queue = append(*queue, &Node{cpy, &newMoves})

		add(visited, cpy, newMoves)
	}
}
