package day17

type Piece struct {
	width int
	shape [4][4]bool
}

var Pieces = []Piece{
	{
		width: 4,
		shape: [4][4]bool{
			{true, true, true, true},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		},
	},
	{
		width: 3,
		shape: [4][4]bool{
			{false, true, false, false},
			{true, true, true, false},
			{false, true, false, false},
			{false, false, false, false},
		},
	},
	{
		width: 3,
		shape: [4][4]bool{
			{true, true, true, false},
			{false, false, true, false},
			{false, false, true, false},
			{false, false, false, false},
		},
	},
	{
		width: 1,
		shape: [4][4]bool{
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
		},
	},
	{
		width: 2,
		shape: [4][4]bool{
			{true, true, false, false},
			{true, true, false, false},
			{false, false, false, false},
			{false, false, false, false},
		},
	},
}
