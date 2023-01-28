package solver

type Cube [26]byte

func NewCube() Cube {
	return [26]byte{
		0, 1, 2, 3, 4, 5, 6, 7, // corners
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, //edges
		0, 1, 2, 3, 4, 5, // centers
	}
}
