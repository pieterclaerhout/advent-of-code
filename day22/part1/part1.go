package part1

import (
	"strconv"
	"strings"
)

type Player struct {
	Row    int
	Col    int
	Facing [2]int
}

var right = [2]int{0, 1}
var left = [2]int{0, -1}
var up = [2]int{-1, 0}
var down = [2]int{1, 0}

func (p *Player) turn(dir byte) {
	switch dir {
	case 'R':
		switch p.Facing {
		case up:
			p.Facing = right
		case down:
			p.Facing = left
		case right:
			p.Facing = down
		case left:
			p.Facing = up
		}
	case 'L':
		switch p.Facing {
		case up:
			p.Facing = left
		case down:
			p.Facing = right
		case right:
			p.Facing = up
		case left:
			p.Facing = down
		}
	}
}

func (player *Player) turnScore() int {
	switch player.Facing {
	case up:
		return 3
	case down:
		return 1
	case right:
		return 0
	case left:
		return 2
	default:
		panic("Invalid direction")
	}
}

func (player *Player) Password() int {
	return 1000*(player.Row+1) + 4*(player.Col+1) + player.turnScore()
}

type Board struct {
	height int
	width  int
	Map    [][]int
	player *Player
}

func (b *Board) run(moves []int, turns []byte) {
	for i, m := range moves {
		b.movePlayer(m)
		if i < len(turns) {
			b.player.turn(turns[i])
		} else {
			break
		}
	}
}

func (b *Board) movePlayer(square int) {
	for i := 0; i < square; i++ {
		newRow := b.player.Row + b.player.Facing[0]
		for {
			if newRow < 0 {
				newRow += b.height
			} else if newRow >= b.height {
				newRow -= b.height
			}
			if b.Map[newRow][b.player.Col] != 0 {
				break
			}
			newRow += b.player.Facing[0]
		}
		if b.Map[newRow][b.player.Col] == 2 { // hit a wall
			break
		} else {
			b.player.Row = newRow
		}

		newCol := b.player.Col + b.player.Facing[1]
		for {
			if newCol < 0 {
				newCol += b.width
			} else if newCol >= b.width {
				newCol -= b.width
			}
			if b.Map[b.player.Row][newCol] != 0 {
				break
			}
			newCol += b.player.Facing[1]
		}
		if b.Map[b.player.Row][newCol] == 2 { // hit a wall
			break
		} else {
			b.player.Col = newCol
		}
	}
}

func parseInput(input []string) (*Board, []int, []byte) {
	var line string
	var lineNum, boardWidth int
	for lineNum, line = range input {
		if len(line) > boardWidth {
			boardWidth = len(line)
		}
		if line == "" {
			break
		}
	}
	boardHeight := lineNum

	b := &Board{}
	b.height = boardHeight
	b.width = boardWidth
	b.Map = make([][]int, 0, boardHeight)
	for lineNum, line = range input[:boardHeight] {
		b.Map = append(b.Map, make([]int, boardWidth))
		for colNum, r := range line {
			switch r {
			case '.':
				b.Map[lineNum][colNum] = 1
			case '#':
				b.Map[lineNum][colNum] = 2
			}
		}
	}
	for colNum, r := range b.Map[0] {
		if r == 1 {
			b.player = &Player{Row: 0, Col: colNum, Facing: right}
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
			} else {
				panic(err)
			}
		} else {
			if tiles, err := strconv.Atoi(instruction[:nextTurn]); err == nil {
				moves = append(moves, tiles)
			} else {
				panic(err)
			}
			turns = append(turns, instruction[nextTurn])
			instruction = instruction[nextTurn+1:]
		}
	}

	return b, moves, turns
}

func Solve(parsed []string) int {
	b, moves, turns := parseInput(parsed)
	b.run(moves, turns)
	return b.player.Password()
}
