package solver

// Each byte has 8 bits, obviously.
// If the byte is 0xFF, it should be disregarded.
// The 3rd and 4th bits indicate the orientation.
// The 5th-8th bits indicate the permutation.
type Cube [26]byte

const disregard = 0xFF

func NewCube() Cube {
	return [26]byte{
		0, 1, 2, 3, 4, 5, 6, 7, // corners
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, //edges
		0, 1, 2, 3, 4, 5, // centers
	}
}

// pieces is an array of pieces indices.
// This function makes it so we only consider the orientation of the pieces.
// The way this works is by settings the 5th-8th bits to 1.
// This means that the permutation is disregarded.
func (c *Cube) SetOnlyOrientation(pieces []int) {
	for _, i := range pieces {
		c[i] |= 0x0F
	}
}

// pieces is an array of pieces indices.
// This function makes it so the pieces are disregarded.
func (c *Cube) SetDisregard(pieces []int) {
	for _, i := range pieces {
		c[i] = disregard
	}
}
