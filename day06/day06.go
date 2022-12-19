package day06

import (
	_ "embed"

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
	slog.Info("Part 1", slog.Any("result", c.firstStartOfPackage(4)))
}

func (c *Command) part2() {
	slog.Info("Part 2", slog.Any("result", c.firstStartOfPackage(14)))
}

func (c *Command) firstStartOfPackage(differentCharactersNeeded int) int {
	for i := range input {
		charactersSet := make(map[byte]bool)
		for j := 0; j < differentCharactersNeeded; j++ {
			charactersSet[input[i+j]] = true
		}
		if len(charactersSet) == differentCharactersNeeded {
			return i + differentCharactersNeeded
		}
	}
	return 0
}
