package day13

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type Command struct{}

func (cmd *Command) Execute(input string) (any, any) {

	pkts := []any{}
	part1 := 0
	for i, s := range strings.Split(string(input), "\n\n") {
		s := strings.Split(s, "\n")
		var a, b any
		json.Unmarshal([]byte(s[0]), &a)
		json.Unmarshal([]byte(s[1]), &b)
		pkts = append(pkts, a, b)

		if cmp(a, b) <= 0 {
			part1 += i + 1
		}
	}

	pkts = append(pkts, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(pkts, func(i, j int) bool { return cmp(pkts[i], pkts[j]) < 0 })

	part2 := 1
	for i, p := range pkts {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			part2 *= i + 1
		}
	}

	return part1, part2
}

func cmp(a, b any) int {
	as, aok := a.([]any)
	bs, bok := b.([]any)

	switch {
	case !aok && !bok:
		return int(a.(float64) - b.(float64))
	case !aok:
		as = []any{a}
	case !bok:
		bs = []any{b}
	}

	for i := 0; i < len(as) && i < len(bs); i++ {
		if c := cmp(as[i], bs[i]); c != 0 {
			return c
		}
	}
	return len(as) - len(bs)
}
