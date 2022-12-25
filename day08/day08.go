package day08

import (
	_ "embed"
	"image"
	"strings"
)

type Command struct {
}

func (c *Command) Execute(input string) (interface{}, interface{}) {
	trees := map[image.Point]int{}
	for y, s := range strings.Fields(input) {
		for x, r := range s {
			trees[image.Point{x, y}] = int(r - '0')
		}
	}

	points := []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	part1 := 0
	part2 := 0

	for p, t := range trees {
		vis, score := 0, 1

		for _, d := range points {
			for i := 1; ; i++ {
				if nt, ok := trees[p.Add(d.Mul(i))]; !ok {
					vis, score = 1, score*(i-1)
					break
				} else if nt >= t {
					score *= i
					break
				}
			}
		}

		part1 += vis
		if score > part2 {
			part2 = score
		}
	}

	return part1, part2
}
