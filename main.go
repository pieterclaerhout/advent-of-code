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
	"github.com/pieterclaerhout/advent-of-code/day10"
	"github.com/pieterclaerhout/advent-of-code/day11"
	"github.com/pieterclaerhout/advent-of-code/day12"
	"github.com/pieterclaerhout/advent-of-code/day13"
	"github.com/pieterclaerhout/advent-of-code/day14"
	"github.com/pieterclaerhout/advent-of-code/day15"
	"github.com/pieterclaerhout/advent-of-code/day16"
	"github.com/pieterclaerhout/advent-of-code/day17"
	"github.com/pieterclaerhout/advent-of-code/day18"
	"github.com/pieterclaerhout/advent-of-code/day19"
	"github.com/pieterclaerhout/advent-of-code/day20"
	"github.com/pieterclaerhout/advent-of-code/day21"
	"github.com/pieterclaerhout/advent-of-code/day22"
	"github.com/pieterclaerhout/advent-of-code/day23"
	"github.com/pieterclaerhout/advent-of-code/day24"
	"github.com/pieterclaerhout/advent-of-code/day25"
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
		&day01.Command{},
		&day02.Command{},
		&day03.Command{},
		&day04.Command{},
		&day05.Command{},
		&day06.Command{},
		&day07.Command{},
		&day08.Command{},
		&day09.Command{},
		&day10.Command{},
		&day11.Command{},
		&day12.Command{},
		&day13.Command{},
		&day14.Command{},
		&day15.Command{},
		&day16.Command{},
		&day17.Command{},
		&day18.Command{},
		&day19.Command{},
		&day20.Command{},
		&day21.Command{},
		&day22.Command{},
		&day23.Command{},
		&day24.Command{},
		&day25.Command{},
	}

	if *day > len(commands) {
		slog.Error("Command not found", nil, slog.Any("day", *day))
		return
	}

	command := commands[*day-1]
	command.Execute()
}
