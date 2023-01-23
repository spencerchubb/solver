package solver

import (
	"fmt"
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

func Solve(cube Cube, moves []int, maxSolutions int, log bool) []string {
	depth := 0
	inverseDepth := 0

	visited := initVisited()
	inverseVisited := initVisited()

	// It is faster with *Node instead of Node
	queue := []*Node{{cube, &[]byte{}}}
	inverseQueue := []*Node{{NewCube(), &[]byte{}}}

	solutionExists := make(map[string]bool)
	var solutions []string
	for loc := 0; ; loc++ {
		node := queue[0]
		queue = queue[1:]

		inverseNode := inverseQueue[0]
		inverseQueue = inverseQueue[1:]

		results := check(inverseNode, visited, &solutions, solutionExists, maxSolutions, log)
		if results != nil {
			return *results
		}

		results = check(node, inverseVisited, &solutions, solutionExists, maxSolutions, log)
		if results != nil {
			return *results
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

func check(node *Node, visited Visited, solutions *[]string, solutionExists map[string]bool, maxSolutions int, log bool) *[]string {
	algs := get(visited, node.cube)
	for _, alg := range algs {
		algStr := algString(alg, *node.moves)
		if solutionExists[algStr] {
			continue
		}
		if log {
			fmt.Println(algStr)
		}
		solutionExists[algStr] = true
		*solutions = append(*solutions, algStr)
		if len(*solutions) >= maxSolutions {
			return solutions
		}
	}
	return nil
}

func goToChild(queue *[]*Node, node *Node, visited Visited, move int) {
	if !sameFace(*node.moves, moveAliases[move]) {
		cpy := node.cube
		allMoves[move](&cpy)
		newMoves := appendImmutable(*node.moves, moveAliases[move])
		*queue = append(*queue, &Node{cpy, &newMoves})

		add(visited, cpy, newMoves)
	}
}
