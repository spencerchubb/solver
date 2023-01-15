package solver

type Moves []string
type Visited map[[48]int][]Moves

func add(visited Visited, facelets [48]int, moves Moves) {
	visited[facelets] = append(visited[facelets], moves)
}

func get(visited Visited, facelets [48]int) []Moves {
	return visited[facelets]
}

func initVisited() Visited {
	var m Visited
	m = make(Visited)
	return m
}
