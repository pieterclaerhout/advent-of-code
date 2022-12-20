package day17

type Rock struct {
	Shape  []byte
	Width  int
	Height int
}

var Rocks = []Rock{
	{Shape: []byte{0b_1111_000}, Width: 4, Height: 1},
	{Shape: []byte{
		0b_010_0000,
		0b_111_0000,
		0b_010_0000,
	}, Width: 3, Height: 3},
	{Shape: []byte{
		0b_111_0000,
		0b_001_0000,
		0b_001_0000,
	}, Width: 3, Height: 3},
	{Shape: []byte{
		0b_1_000000,
		0b_1_000000,
		0b_1_000000,
		0b_1_000000,
	}, Width: 1, Height: 4},
	{Shape: []byte{
		0b_11_00000,
		0b_11_00000,
	}, Width: 2, Height: 2},
}
