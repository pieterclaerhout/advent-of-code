package day07

import (
	"fmt"
	"path"
	"strings"
)

type Command struct{}

func (cmd *Command) Execute(input string) (any, any) {
	fs, cd := map[string]int{}, ""
	for _, s := range strings.Split(input, "\n") {
		var size int
		var name string

		if strings.HasPrefix(s, "$ cd") {
			cd = path.Join(cd, strings.Fields(s)[2])
		} else if _, err := fmt.Sscanf(s, "%d %s", &size, &name); err == nil {
			for d := cd; d != "/"; d = path.Dir(d) {
				fs[d] += size
			}
			fs["/"] += size
		}
	}

	part1 := 0
	for _, s := range fs {
		if s <= 100000 {
			part1 += s
		}
	}

	part2 := fs["/"]
	for _, s := range fs {
		if s+70000000-fs["/"] >= 30000000 && s < part2 {
			part2 = s
		}
	}

	return part1, part2
}
