package day21

import (
	"sort"
	"strconv"
	"strings"
)

type Command struct{}

func (cmd *Command) Execute(input string) (any, any) {
	monkeys := map[string]string{}

	for _, s := range strings.Split(input, "\n") {
		s := strings.Split(s, ": ")
		monkeys[s[0]] = s[1]
	}

	part1 := cmd.solve(monkeys, "root")

	monkeys["humn"] = "0"
	s := strings.Fields(monkeys["root"])
	if cmd.solve(monkeys, s[0]) < cmd.solve(monkeys, s[2]) {
		s[0], s[2] = s[2], s[0]
	}

	part2, _ := sort.Find(1e16, func(v int) int {
		monkeys["humn"] = strconv.Itoa(v)
		return cmd.solve(monkeys, s[0]) - cmd.solve(monkeys, s[2])
	})

	return part1, part2
}

func (cmd *Command) solve(monkeys map[string]string, expr string) int {
	if v, err := strconv.Atoi(monkeys[expr]); err == nil {
		return v
	}

	mapping := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	s := strings.Fields(monkeys[expr])
	left := s[0]
	oper := s[1]
	right := s[2]

	return mapping[oper](cmd.solve(monkeys, left), cmd.solve(monkeys, right))
}
