package main

import (
	"flag"

	"github.com/pieterclaerhout/advent-of-code/day01"
	"github.com/pieterclaerhout/advent-of-code/day02"
	"github.com/pieterclaerhout/advent-of-code/day03"
	"github.com/pieterclaerhout/advent-of-code/day04"
	"golang.org/x/exp/slog"
)

var (
	day = flag.Int("day", 0, "day to execute")
)

type Command interface {
	Execute()
}

func main() {
	flag.Parse()

	slog.Info("Executing", slog.Any("day", *day))

	commands := map[int]Command{
		1: day01.Command{},
		2: day02.Command{},
		3: day03.Command{},
		4: day04.Command{},
	}

	command, exists := commands[*day]
	if !exists {
		slog.Error("Command not found", nil, slog.Any("day", *day))
		return
	}

	command.Execute()
}
