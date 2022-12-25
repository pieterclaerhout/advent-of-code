package day25

import (
	_ "embed"
	"strings"

	"golang.org/x/exp/slog"
)

//go:embed input.txt
var input string

type Command struct {
}

func (c *Command) Execute() {

	mapping := map[rune]int{'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}

	sum := 0
	for _, s := range strings.Fields(input) {
		n := 0
		for _, r := range s {
			n = 5*n + mapping[r]
		}
		sum += n
	}

	slog.Info("Part 1", slog.Any("result", sum))

	snafu := ""
	for sum > 0 {
		snafu = string("=-012"[(sum+2)%5]) + snafu
		sum = (sum + 2) / 5
	}

	slog.Info("Part 2", slog.Any("result", snafu))
}
