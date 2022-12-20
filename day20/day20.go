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
	input, indexes := c.parse()

	mixedInput, _ := Mix(input, indexes, 1)

	slog.Info("Part 1", slog.Any("result", Sum(mixedInput, 1)))
}

func (c *Command) part2() {
	const decryptionKey = 811589153

	input, indexes := c.parse()

	for r := 0; r < 10; r++ {
		input, indexes = Mix(input, indexes, decryptionKey)
	}

	slog.Info("Part 2", slog.Any("result", Sum(input, decryptionKey)))

}

func (c *Command) parse() ([]int, []int) {
	numbers := []int{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		number, _ := strconv.Atoi(sc.Text())
		numbers = append(numbers, number)
	}

	indexes := make([]int, len(numbers))
	for i := range numbers {
		indexes[i] = i
	}

	return numbers, indexes
}
