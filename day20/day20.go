package day20

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
	input := c.parse()
	front := NewList(input)
	Mix(front, len(input))
	sum := GroveCoordinates(front)

	slog.Info("Part 1", slog.Any("result", sum))
}

func (c *Command) part2() {
	input := c.parse()
	front := NewList(input)
	ApplyDecryptKey(front, 811589153)
	for i := 0; i < 10; i++ {
		Mix(front, len(input))
	}

	sum := GroveCoordinates(front)
	slog.Info("Part 2", slog.Any("result", sum))
}

func (c *Command) parse() []int {
	numbers := []int{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		number, _ := strconv.Atoi(sc.Text())
		numbers = append(numbers, number)

	}

	return numbers
}
