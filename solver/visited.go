package solver

type Moves []byte
type Visited map[Facelets][]Moves

func add(visited Visited, facelets Facelets, moves Moves) {
	visited[facelets] = append(visited[facelets], moves)
}

func get(visited Visited, facelets Facelets) []Moves {
	return visited[facelets]
}

func initVisited() Visited {
	var m Visited
	m = make(Visited)
	return m
}
