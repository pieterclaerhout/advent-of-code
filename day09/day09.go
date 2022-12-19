package day09

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"

	"golang.org/x/exp/slog"
)

//go:embed input.txt
var input string

type Command struct {
}

func (c *Command) Execute() {
	c.part1()
	c.part2()
}

func (c *Command) part1() {
	visitedByTail := make(map[Point]bool)

	head := Point{X: 0, Y: 0}
	tail := Point{X: 0, Y: 0}
	visitedByTail[tail] = true

	for _, motion := range c.parse() {
		moves := motion.Moves
		for moves > 0 {
			switch motion.Direction {
			case 'U':
				head.Y++
			case 'R':
				head.X++
			case 'D':
				head.Y--
			case 'L':
				head.X--
			}
			moves--
			tail = tail.AdjustTail1(head)
			visitedByTail[tail] = true
		}
	}

	slog.Info("Part 1", slog.Any("visitedByTail", len(visitedByTail)))
}

func (c *Command) part2() {
	visitedByTail := make(map[Point]bool)
	knots := make([]Point, 10)

	visitedByTail[knots[9]] = true

	for _, motion := range c.parse() {
		moves := motion.Moves
		for moves > 0 {
			switch motion.Direction {
			case 'U':
				knots[0].Y++
			case 'R':
				knots[0].X++
			case 'D':
				knots[0].Y--
			case 'L':
				knots[0].X--
			}
			for i := range knots[:len(knots)-1] {
				knots[i+1] = knots[i+1].AdjustTail2(knots[i])
			}
			moves--
			visitedByTail[knots[9]] = true
		}
	}

	slog.Info("Part 2", slog.Any("visitedByTail", len(visitedByTail)))
}

func (c *Command) parse() []Motion {
	var list []Motion

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		direction := rune(sc.Text()[0])
		moves, _ := strconv.Atoi(sc.Text()[2:])
		list = append(list, Motion{
			Direction: direction,
			Moves:     moves,
		})
	}

	return list
}
