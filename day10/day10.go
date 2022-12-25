package day10

import (
	_ "embed"
	"fmt"
	"strings"
)

type Command struct{}

func (cmd *Command) Execute(input string) (interface{}, interface{}) {

	w := 40

	c := 0
	x := 1
	part1 := 0
	part2 := ""

	tick := func() {
		part2 += map[bool]string{
			true:  "██",
			false: "  ",
		}[c%w >= x-1 && c%w <= x+1]

		part2 += map[bool]string{true: "\n"}[c%w == w-1]
		if c++; (c+w/2)%w == 0 {
			part1 += c * x
		}
	}

	for _, s := range strings.Split(input, "\n") {
		var ins string
		var v int
		fmt.Sscanf(s, "%s %d", &ins, &v)

		tick()
		if ins == "addx" {
			tick()
			x += v
		}
	}

	return part1, "\n" + strings.TrimSpace(part2)
}
