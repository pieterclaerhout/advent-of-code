package day15

import (
	"fmt"
	"image"
	"strings"
)

type Command struct {
	Row1 int
	Row2 int
}

func (cmd *Command) Execute(input string) (any, any) {
	sensors := map[image.Point]int{}
	line := map[int]struct{}{}

	for _, s := range strings.Split(input, "\n") {
		var a, b image.Point
		fmt.Sscanf(s, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &a.X, &a.Y, &b.X, &b.Y)

		sensors[a] = abs(a.X-b.X) + abs(a.Y-b.Y)
		d := sensors[a] - abs(cmd.Row1-a.Y)

		for x := a.X - d; x <= a.X+d; x++ {
			if !(b.X == x && b.Y == cmd.Row1) {
				line[x] = struct{}{}
			}
		}
	}

	part1 := len(line)

	for y := 0; y <= cmd.Row2; y++ {
	loop:
		for x := 0; x <= cmd.Row2; x++ {
			for s, d := range sensors {
				if dx, dy := s.X-x, s.Y-y; abs(dx)+abs(dy) <= d {
					x += d - abs(dy) + dx
					continue loop
				}
			}
			return part1, x*4000000 + y
		}
	}

	return part1, 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
