package solver

type Algorithm []byte
type Algorithms []Algorithm

type Visited map[Cube]Algorithms

func add(visited Visited, cube Cube, alg Algorithm) {
	visited[cube] = append(visited[cube], alg)
}

func get(visited Visited, cube Cube) Algorithms {
	return visited[cube]
}

func initVisited() Visited {
	return make(Visited, 10_000_000)
}
