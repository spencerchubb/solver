package solver

type Cube [20]byte

func NewCube() Cube {
	return [20]byte{
		0, 1, 2, 3, 4, 5, 6, 7, // corners
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, //edges
	}
}
