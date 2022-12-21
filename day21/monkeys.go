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

	s := strings.Fields(expr)

	if len(s) == 3 {
		return mapping[s[1]](
			monkeys.Solve(s[0]),
			monkeys.Solve(s[2]),
		)
	}

	v, _ := strconv.Atoi(expr)
	return v
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
