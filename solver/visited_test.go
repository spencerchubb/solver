package solver

import "testing"

func assertAlgsEqual(t *testing.T, expected, actual Algorithms) {
	if len(expected) != len(actual) {
		t.Errorf("Expected %d algorithms, got %d", len(expected), len(actual))
	}

	for i := 0; i < len(expected); i++ {
		if len(expected[i]) != len(actual[i]) {
			t.Errorf("Expected %d moves, got %d", len(expected[i]), len(actual[i]))
		}
		for j := 0; j < len(expected[i]); j++ {
			if expected[i][j] != actual[i][j] {
				t.Errorf("Expected %v, got %v", expected, actual)

			}
		}
	}
}

func TestAdd(t *testing.T) {
	c1 := Cube{0, 1, 2, 3}
	c2 := Cube{2, 3, 4, 5}

	visited := initVisited()
	add(visited, c1, Algorithm{0})
	add(visited, c1, Algorithm{3})
	add(visited, c2, Algorithm{6})

	movesArray := get(visited, c1)
	if len(movesArray) != 2 {
		t.Errorf("Expected 2 moves, got %d", len(movesArray))
	}
	assertAlgsEqual(t, movesArray, Algorithms{Algorithm{0}, Algorithm{3}})

	movesArray = get(visited, c2)
	assertAlgsEqual(t, movesArray, Algorithms{Algorithm{6}})
}
