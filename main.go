package main

import (
	"flag"

	"github.com/pieterclaerhout/advent-of-code/day01"
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

	var command Command
	switch *day {
	case 1:
		command = day01.Command{}
	}

	if command == nil {
		slog.Error("Command not found", nil, slog.Any("day", *day))
		return
	}

	command.Execute()
}
