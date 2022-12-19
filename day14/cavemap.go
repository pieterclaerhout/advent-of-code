package day14

import (
	"strconv"
)

type CaveMap [][]bool

func (m CaveMap) draw(p1 []string, p2 []string) {
	xP1, _ := strconv.Atoi(p1[0])
	yP1, _ := strconv.Atoi(p1[1])
	xP2, _ := strconv.Atoi(p2[0])
	yP2, _ := strconv.Atoi(p2[1])

	if xP1 == xP2 {
		for i := min(yP1, yP2); i <= max(yP1, yP2); i++ {
			m[i][xP1] = true
		}
	} else if yP1 == yP2 {
		for i := min(xP1, xP2); i <= max(xP1, xP2); i++ {
			m[yP1][i] = true
		}
	} else {
		panic("not vertica/horizontal line")
	}
}

func (m CaveMap) simulateFallingSand() int {
	counter := 0

	maxY := len(m)
	for l := maxY - 1; l > 0; l-- {
		for x := 0; x < len(m[l]); x++ {
			if m[l][x] {
				maxY = l
				break
			}
		}
		if maxY == l {
			break
		}
	}

	sand := Point{500, 0}
	for {
		if !m[sand.Y+1][sand.X] {
			if sand.Y > maxY+1 {
				break
			}
			sand.Y++
		} else if !m[sand.Y+1][sand.X-1] {
			sand.Y++
			sand.X--
		} else if !m[sand.Y+1][sand.X+1] {
			sand.Y++
			sand.X++
		} else {
			counter++
			m[sand.Y][sand.X] = true
			sand = Point{500, 0}
		}
	}
	return counter
}

func (m CaveMap) simulateBlockSource() int {
	counter := 0

	maxY := len(m)
	for l := maxY - 1; l > 0; l-- {
		for x := 0; x < len(m[l]); x++ {
			if m[l][x] {
				maxY = l
				break
			}
		}
		if maxY == l {
			break
		}
	}

	floorY := maxY + 2

	sand := Point{500, 0}
	for {
		if !m[sand.Y+1][sand.X] && sand.Y+1 < floorY {
			sand.Y++
		} else if !m[sand.Y+1][sand.X-1] && sand.Y+1 < floorY {
			sand.Y++
			sand.X--
		} else if !m[sand.Y+1][sand.X+1] && sand.Y+1 < floorY {
			sand.Y++
			sand.X++
		} else {
			counter++
			m[sand.Y][sand.X] = true
			if sand.X == 500 && sand.Y == 0 {
				break
			}
			sand = Point{500, 0}
		}
	}
	return counter
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
