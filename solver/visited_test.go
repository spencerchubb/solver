package solver

import "testing"

func TestAdd(t *testing.T) {
	visited := initVisited()
	add(visited, [48]Facelet{0, 1, 2, 3}, []byte{'U'})
	add(visited, [48]Facelet{0, 1, 2, 3}, []byte{'F'})
	add(visited, [48]Facelet{4, 5, 6, 7}, []byte{'D'})

	movesArray := get(visited, [48]Facelet{0, 1, 2, 3})
	if len(movesArray) != 2 {
		t.Errorf("Expected 2 moves, got %d", len(movesArray))
	}
	if movesArray[0][0] != 'U' {
		t.Errorf("Expected U, got %b", movesArray[0][0])
	}
	if movesArray[1][0] != 'F' {
		t.Errorf("Expected F, got %b", movesArray[1][0])
	}

	movesArray = get(visited, [48]Facelet{4, 5, 6, 7})
	if len(movesArray) != 1 {
		t.Errorf("Expected 1 move, got %d", len(movesArray))
	}
	if movesArray[0][0] != 'D' {
		t.Errorf("Expected D, got %b", movesArray[0][0])
	}
}
