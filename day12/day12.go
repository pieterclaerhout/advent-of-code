package day12

import (
	"bufio"
	_ "embed"
	"strings"

	"golang.org/x/exp/slog"
)

//go:embed input.txt
var input string

type Command struct {
	RegisterX   int
	CycleNumber int
	FinalValue  int
}

func (c *Command) Execute() {
	c.part1()
	c.part2()
}

func (c *Command) part1() {
	heightmap := c.parse()

	visited := make([][]int, len(heightmap.Data))
	for i := range visited {
		visited[i] = make([]int, len(heightmap.Data[0]))
		for j := range visited[i] {
			visited[i][j] = 1 << 32
		}
	}

	visited[heightmap.Start.Y][heightmap.Start.X] = 0
	fifo := []Point{heightmap.Start}

	for {
		curr := fifo[0]
		y := curr.Y
		x := curr.X

		if y == heightmap.End.Y && x == heightmap.End.X {
			slog.Info("Part 1", slog.Any("result", visited[y][x]))
			return
		}

		// right
		if x+1 < len(heightmap.Data[y]) && visited[y][x]+1 < visited[y][x+1] && heightmap.Data[y][x]+1 >= heightmap.Data[y][x+1] {
			fifo = append(fifo, Point{Y: y, X: x + 1})
			visited[y][x+1] = visited[y][x] + 1
		}
		// left
		if x-1 >= 0 && visited[y][x]+1 < visited[y][x-1] && heightmap.Data[y][x]+1 >= heightmap.Data[y][x-1] {
			fifo = append(fifo, Point{Y: y, X: x - 1})
			visited[y][x-1] = visited[y][x] + 1
		}

		// top
		if y-1 >= 0 && visited[y][x]+1 < visited[y-1][x] && heightmap.Data[y][x]+1 >= heightmap.Data[y-1][x] {
			fifo = append(fifo, Point{Y: y - 1, X: x})
			visited[y-1][x] = visited[y][x] + 1
		}

		// down
		if y+1 < len(visited) && visited[y][x]+1 < visited[y+1][x] && heightmap.Data[y][x]+1 >= heightmap.Data[y+1][x] {
			fifo = append(fifo, Point{Y: y + 1, X: x})
			visited[y+1][x] = visited[y][x] + 1
		}

		fifo = fifo[1:]
	}
}

func (c *Command) part2() {
	heightmap := c.parse()

	visited := make([][]int, len(heightmap.Data))
	for i := range visited {
		visited[i] = make([]int, len(heightmap.Data[0]))
		for j := range visited[i] {
			visited[i][j] = 1 << 32
		}
	}

	visited[heightmap.End.Y][heightmap.End.X] = 0
	fifo := []Point{heightmap.End}

	for {
		curr := fifo[0]
		y := curr.Y
		x := curr.X

		if heightmap.Data[y][x] == 'a' {
			slog.Info("Part 2", slog.Any("result", visited[y][x]))
			return
		}

		// right
		if x+1 < len(heightmap.Data[y]) && visited[y][x]+1 < visited[y][x+1] && heightmap.Data[y][x]-1 <= heightmap.Data[y][x+1] {
			fifo = append(fifo, Point{Y: y, X: x + 1})
			visited[y][x+1] = visited[y][x] + 1
		}
		// left
		if x-1 >= 0 && visited[y][x]+1 < visited[y][x-1] && heightmap.Data[y][x]-1 <= heightmap.Data[y][x-1] {
			fifo = append(fifo, Point{Y: y, X: x - 1})
			visited[y][x-1] = visited[y][x] + 1
		}

		// top
		if y-1 >= 0 && visited[y][x]+1 < visited[y-1][x] && heightmap.Data[y][x]-1 <= heightmap.Data[y-1][x] {
			fifo = append(fifo, Point{Y: y - 1, X: x})
			visited[y-1][x] = visited[y][x] + 1
		}

		// down
		if y+1 < len(visited) && visited[y][x]+1 < visited[y+1][x] && heightmap.Data[y][x]-1 <= heightmap.Data[y+1][x] {
			fifo = append(fifo, Point{Y: y + 1, X: x})
			visited[y+1][x] = visited[y][x] + 1
		}

		fifo = fifo[1:]
	}
}

func (c *Command) parse() *HeightMap {
	heightmap := &HeightMap{
		Data: make([][]rune, 0),
	}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		heightmap.Data = append(heightmap.Data, []rune(sc.Text()))
	}

	for y, line := range heightmap.Data {
		for x, elevation := range line {
			if elevation == 'S' {
				heightmap.Start = Point{x, y}
				elevation = 'a'
			}
			if elevation == 'E' {
				heightmap.End = Point{x, y}
				elevation = 'z'
			}
			heightmap.Data[y][x] = elevation
		}
	}

	return heightmap
}
