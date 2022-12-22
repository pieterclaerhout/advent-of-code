package part2

import (
	"fmt"
	"strconv"
	"strings"
)

type dir int

const right dir = 0
const down dir = 1
const left dir = 2
const up dir = 3
const invalidDir dir = -1

/*
	Handcrafted from looking at a paper cube

oldFaceNum direction -> newFaceNum newDirection
1 up -> 6 right
1 down -> 3 down
1 right -> 2 right
1 left -> 4 right
2 up -> 6 up
2 down -> 3 left
2 right -> 5 left
2 left -> 1 left
3 up -> 1 up
3 down -> 5 down
3 right -> 2 up
3 left -> 4 down
4 up -> 3 right
4 down -> 6 down
4 right -> 5 right
4 left -> 1 right
5 up -> 3 up
5 down -> 6 left
5 right -> 2 left
5 left -> 4 left
6 up -> 4 up
6 down -> 2 down
6 right -> 5 up
6 left -> 1 down
*/
var mapChange = [6][6]dir{
	{invalidDir, right, down, right, invalidDir, right},
	{left, invalidDir, left, invalidDir, left, up},
	{up, up, invalidDir, down, down, invalidDir},
	{right, invalidDir, right, invalidDir, right, down},
	{invalidDir, left, up, left, invalidDir, left},
	{down, down, invalidDir, up, up, invalidDir},
}

type vec2d [2]int

type Player struct {
	cell   *Cell
	facing dir
}

func (p *Player) move() bool {
	var next *Cell
	switch p.facing {
	case right:
		next = p.cell.right
	case down:
		next = p.cell.down
	case left:
		next = p.cell.left
	case up:
		next = p.cell.up
	}
	if next.val == 2 {
		return false
	}

	if next.face != p.cell.face {
		p.facing = mapChange[p.cell.face][next.face]
	}
	p.cell = next
	return true
}

func (p *Player) password() int {
	return 1000*(p.cell.pos[0]+1) + 4*(p.cell.pos[1]+1) + int(p.facing)
}

type Cell struct {
	face                  int
	pos                   vec2d
	val                   int
	right, down, left, up *Cell
}

func (c *Cell) setDown(downCell *Cell) {
	c.down = downCell
	downCell.up = c
}

func (c *Cell) setRight(rightCell *Cell) {
	c.right = rightCell
	rightCell.left = c
}

type Board struct {
	cells  map[vec2d]*Cell
	player *Player
}

func turn(turnDir byte, facing dir) dir {
	switch turnDir {
	case 'R':
		switch facing {
		case up:
			return right
		case down:
			return left
		case right:
			return down
		case left:
			return up
		}
	case 'L':
		switch facing {
		case up:
			return left
		case down:
			return right
		case right:
			return up
		case left:
			return down
		}
	}
	panic("Invalid turn")
}

func (b *Board) run(moves []int, turns []byte) {
	for i, m := range moves {
		for j := 0; j < m; j++ {
			if !b.player.move() {
				break
			}
		}
		if i < len(turns) {
			b.player.facing = turn(turns[i], b.player.facing)
		} else {
			break
		}
	}
}

func parseCells(input []string, boardWidth int) map[vec2d]*Cell {
	faceSize := boardWidth / 3
	faceCornerRows := []int{0, 0, faceSize, faceSize * 2, faceSize * 2, faceSize * 3}
	faceCornerCols := []int{faceSize, faceSize * 2, faceSize, 0, faceSize, 0}

	var faceCells [6]map[vec2d]*Cell
	cells := make(map[vec2d]*Cell, boardWidth*boardWidth)
	var row, col, val int

	for faceNum := range faceCornerRows {
		faceCells[faceNum] = make(map[vec2d]*Cell)
		cornerRow, cornerCol := faceCornerRows[faceNum], faceCornerCols[faceNum]
		for row := 0; row < faceSize; row++ {
			boardRow := cornerRow + row
			for col := 0; col < faceSize; col++ {
				boardCol := cornerCol + col

				switch input[boardRow][boardCol] {
				case ' ':
					val = 0
				case '.':
					val = 1
				case '#':
					val = 2
				}

				boardPos := [2]int{boardRow, boardCol}
				newCell := Cell{
					face: faceNum,
					val:  val,
					pos:  boardPos,
				}
				cells[boardPos] = &newCell

				facePos := [2]int{row, col}
				faceCells[faceNum][facePos] = &newCell
			}
		}
	}

	// Set regular neighbors
	for faceNum := range faceCornerRows {
		cornerRow, cornerCol := faceCornerRows[faceNum], faceCornerCols[faceNum]
		for r := 0; r < faceSize; r++ {
			row = cornerRow + r
			for c := 0; c < faceSize; c++ {
				col = cornerCol + c
				pos := [2]int{row, col}

				sourceCell := cells[pos]
				rightPos, downPos := [2]int{row, col + 1}, [2]int{row + 1, col}
				if c+1 < faceSize {
					sourceCell.setRight(cells[rightPos])
				}
				if r+1 < faceSize {
					sourceCell.setDown(cells[downPos])
				}
			}
		}
	}

	n := faceSize - 1
	for i := 0; i < faceSize; i++ {
		faceCells[0][vec2d{0, i}].up = faceCells[5][vec2d{i, 0}]
		faceCells[5][vec2d{i, 0}].left = faceCells[0][vec2d{0, i}]

		faceCells[0][vec2d{n, i}].setDown(faceCells[2][vec2d{0, i}])

		faceCells[0][vec2d{i, n}].setRight(faceCells[1][vec2d{i, 0}])

		faceCells[0][vec2d{i, 0}].left = faceCells[3][vec2d{n - i, 0}]
		faceCells[3][vec2d{n - i, 0}].left = faceCells[0][vec2d{i, 0}]

		faceCells[5][vec2d{n, i}].setDown(faceCells[1][vec2d{0, i}])

		faceCells[1][vec2d{n, i}].down = faceCells[2][vec2d{i, n}]
		faceCells[2][vec2d{i, n}].right = faceCells[1][vec2d{n, i}]

		faceCells[1][vec2d{i, n}].right = faceCells[4][vec2d{n - i, n}]
		faceCells[4][vec2d{n - i, n}].right = faceCells[1][vec2d{i, n}]

		faceCells[2][vec2d{n, i}].setDown(faceCells[4][vec2d{0, i}])

		faceCells[2][vec2d{i, 0}].left = faceCells[3][vec2d{0, i}]
		faceCells[3][vec2d{0, i}].up = faceCells[2][vec2d{i, 0}]

		faceCells[3][vec2d{n, i}].setDown(faceCells[5][vec2d{0, i}])
		faceCells[3][vec2d{i, n}].setRight(faceCells[4][vec2d{i, 0}])

		faceCells[4][vec2d{n, i}].down = faceCells[5][vec2d{i, n}]
		faceCells[5][vec2d{i, n}].right = faceCells[4][vec2d{n, i}]
	}

	return cells
}

func parseInput(input []string) (*Board, []int, []byte) {
	var boardWidth int
	var lineNum int
	var line string
	for lineNum, line = range input {
		if len(line) > boardWidth {
			boardWidth = len(line)
		}
		if line == "" {
			break
		}
	}
	boardHeight := lineNum

	b := new(Board)
	b.cells = parseCells(input, boardWidth)

	for colNum := 0; colNum < boardWidth; colNum++ {
		if c, ok := b.cells[vec2d{0, colNum}]; ok && c.val == 1 {
			b.player = &Player{cell: c, facing: right}
			break
		}
	}

	instruction := input[boardHeight+1]
	moves, turns := make([]int, 0), make([]byte, 0)
	for len(instruction) > 0 {
		nextTurn := strings.IndexAny(instruction, "LR")
		if nextTurn == -1 {
			if tiles, err := strconv.Atoi(instruction); err == nil {
				moves = append(moves, tiles)
				break
			}
		} else {
			if tiles, err := strconv.Atoi(instruction[:nextTurn]); err == nil {
				moves = append(moves, tiles)
			}
			turns = append(turns, instruction[nextTurn])
			instruction = instruction[nextTurn+1:]
		}
	}

	return b, moves, turns
}

func Solve(parsed []string) int {
	cube, moves, turns := parseInput(parsed)

	for _, c := range cube.cells {
		if c.up == nil || c.down == nil || c.right == nil || c.left == nil {
			fmt.Println(*c)
		}
	}
	cube.run(moves, turns)

	return cube.player.password()
}
