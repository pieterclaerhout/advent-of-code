package day01

import (
	_ "embed"

	"golang.org/x/exp/slog"
)

//go:embed input.txt
var input string

type Command struct {
}

func (c Command) Execute() {
	c.part1()
	c.part2()
}

func (c Command) part1() {
	calories := NewIntSlice(input)
	slog.Info("Max calories", slog.Any("max", calories.Max()))
}

func (c Command) part2() {
	calories := NewIntSlice(input)
	slog.Info("Sum", slog.Any("sum", calories.SumTop(3)))
}
