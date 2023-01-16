package solver

import "testing"

func assertFaceletsEq(t *testing.T, testName string, expected, actual Facelets) {
	if !faceletsEq(expected, actual) {
		t.Errorf("%s failed. Expected %s, got %s", testName, faceletsToStr(expected), faceletsToStr(actual))
	}
}

func TestU1(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	U1(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, R, F, F, R, F, R, F, F, D, D, D, D, D, D, D, D, B, B, L, B, L, B, B, L, F, L, L, F, L, F, L, L, B, R, R, B, R, B, R, R}
	assertFaceletsEq(t, "TestU1", expected, facelets)
}

func TestU2(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	U2(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	U1(&expected)
	U1(&expected)
	assertFaceletsEq(t, "TestU2", expected, facelets)
}

func TestU3(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	U3(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	U1(&expected)
	U1(&expected)
	U1(&expected)
	assertFaceletsEq(t, "TestU3", expected, facelets)

}

func TestF1(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	F1(&facelets)
	expected := Facelets{U, U, L, U, L, U, U, L, F, F, F, F, F, F, F, F, R, D, D, R, D, R, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, D, D, D, U, U, U, R, R, R, R, R}
	assertFaceletsEq(t, "TestF1", expected, facelets)

}

func TestF2(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	F2(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	F1(&expected)
	F1(&expected)
	assertFaceletsEq(t, "TestF2", expected, facelets)

}

func TestF3(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	F3(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	F1(&expected)
	F1(&expected)
	F1(&expected)
	assertFaceletsEq(t, "TestF3", expected, facelets)
}

func TestD1(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	D1(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, F, F, L, F, L, F, F, L, D, D, D, D, D, D, D, D, R, B, B, R, B, R, B, B, L, L, B, L, B, L, L, B, R, R, F, R, F, R, R, F}
	assertFaceletsEq(t, "TestD1", expected, facelets)
}

func TestD2(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	D2(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	D1(&expected)
	D1(&expected)
	assertFaceletsEq(t, "TestD2", expected, facelets)
}

func TestD3(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	D3(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	D1(&expected)
	D1(&expected)
	D1(&expected)
	assertFaceletsEq(t, "TestD3", expected, facelets)
}

func TestB1(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	B1(&facelets)
	expected := Facelets{R, U, U, R, U, R, U, U, F, F, F, F, F, F, F, F, D, D, L, D, L, D, D, L, B, B, B, B, B, B, B, B, U, U, U, L, L, L, L, L, R, R, R, R, R, D, D, D}
	assertFaceletsEq(t, "TestB1", expected, facelets)
}

func TestB2(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	B2(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	B1(&expected)
	B1(&expected)
	assertFaceletsEq(t, "TestB2", expected, facelets)
}

func TestB3(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	B3(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	B1(&expected)
	B1(&expected)
	B1(&expected)
	assertFaceletsEq(t, "TestB3", expected, facelets)
}

func TestL1(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	L1(&facelets)
	expected := Facelets{B, B, B, U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	assertFaceletsEq(t, "TestL1", expected, facelets)
}

func TestL2(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	L2(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	L1(&expected)
	L1(&expected)
	assertFaceletsEq(t, "TestL2", expected, facelets)
}

func TestL3(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	L3(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	L1(&expected)
	L1(&expected)
	L1(&expected)
	assertFaceletsEq(t, "TestL3", expected, facelets)
}

func TestR1(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	R1(&facelets)
	expected := Facelets{U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, U, U, U, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	assertFaceletsEq(t, "TestR1", expected, facelets)
}

func TestR2(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	R2(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	R1(&expected)
	R1(&expected)
	assertFaceletsEq(t, "TestR2", expected, facelets)
}

func TestR3(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	R3(&facelets)
	expected := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	R1(&expected)
	R1(&expected)
	R1(&expected)
	assertFaceletsEq(t, "TestR3", expected, facelets)
}

func TestAll(t *testing.T) {
	facelets := Facelets{U, U, U, U, U, U, U, U, F, F, F, F, F, F, F, F, D, D, D, D, D, D, D, D, B, B, B, B, B, B, B, B, L, L, L, L, L, L, L, L, R, R, R, R, R, R, R, R}
	U1(&facelets)
	F1(&facelets)
	D1(&facelets)
	B1(&facelets)
	L1(&facelets)
	R1(&facelets)
	U2(&facelets)
	F2(&facelets)
	D2(&facelets)
	B2(&facelets)
	L2(&facelets)
	R2(&facelets)
	U3(&facelets)
	F3(&facelets)
	D3(&facelets)
	B3(&facelets)
	L3(&facelets)
	R3(&facelets)
	expected := Facelets{U, D, L, F, B, F, D, B, U, B, B, D, U, D, D, U, D, U, U, L, F, F, B, D, B, R, B, U, R, F, L, R, L, B, R, R, R, F, U, R, L, L, R, F, L, D, F, L}
	assertFaceletsEq(t, "TestAll", expected, facelets)

}
