package day08

type Point struct {
	X int
	Y int
}

type Forest [][]rune

func (forest Forest) CalcViewRight(i int, j int, house rune, firstRec bool) int {
	if j == len(forest[0])-1 || !firstRec && forest[i][j] >= house {
		return 0
	}
	return 1 + forest.CalcViewRight(i, j+1, house, false)
}

func (forest Forest) CalcViewLeft(i int, j int, house rune, firstRec bool) int {
	if j == 0 || !firstRec && forest[i][j] >= house {
		return 0
	}
	return 1 + forest.CalcViewLeft(i, j-1, house, false)
}

func (forest Forest) CalcViewDown(i int, j int, house rune, firstRec bool) int {
	if i == len(forest)-1 || !firstRec && forest[i][j] >= house {
		return 0
	}
	return 1 + forest.CalcViewDown(i+1, j, house, false)
}

func (forest Forest) CalcViewTop(i int, j int, house rune, firstRec bool) int {
	if i == 0 || !firstRec && forest[i][j] >= house {
		return 0
	}
	return 1 + forest.CalcViewTop(i-1, j, house, false)
}
