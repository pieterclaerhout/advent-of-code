package main

import (
	"flag"

	"github.com/pieterclaerhout/advent-of-code/day01"
	"github.com/pieterclaerhout/advent-of-code/day02"
	"github.com/pieterclaerhout/advent-of-code/day03"
	"github.com/pieterclaerhout/advent-of-code/day04"
	"github.com/pieterclaerhout/advent-of-code/day05"
	"github.com/pieterclaerhout/advent-of-code/day06"
	"github.com/pieterclaerhout/advent-of-code/day07"
	"github.com/pieterclaerhout/advent-of-code/day08"
	"github.com/pieterclaerhout/advent-of-code/day09"
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

	commands := []Command{
		day01.Command{},
		day02.Command{},
		day03.Command{},
		day04.Command{},
		day05.Command{},
		day06.Command{},
		day07.Command{},
		day08.Command{},
		day09.Command{},
	}

	if *day > len(commands) {
		slog.Error("Command not found", nil, slog.Any("day", *day))
		return
	}

	command := commands[*day-1]
	command.Execute()
}
