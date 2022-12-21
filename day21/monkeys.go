package day21

import (
	"sort"
	"strconv"
	"strings"
)

type Monkeys map[string]string

var mapping = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func (monkeys Monkeys) Solve(key string) int {
	expr := monkeys[key]

	if v, err := strconv.Atoi(expr); err == nil {
		return v
	}

	s := strings.Fields(expr)
	left := monkeys.Solve(s[0])
	oper := s[1]
	right := monkeys.Solve(s[2])

	return mapping[oper](left, right)
}

func (monkeys Monkeys) RootEquality() int {
	monkeys["humn"] = "0"

	s := strings.Fields(monkeys["root"])
	if monkeys.Solve(s[0]) < monkeys.Solve(s[2]) {
		s[0], s[2] = s[2], s[0]
	}

	i, _ := sort.Find(1e16, func(v int) int {
		monkeys["humn"] = strconv.Itoa(v)
		return monkeys.Solve(s[0]) - monkeys.Solve(s[2])
	})

	return i
}
