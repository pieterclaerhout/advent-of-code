package day17

import (
	"fmt"
)

type Command struct{}

func (cmd *Command) Execute(input string) (any, any) {
	return partOne(input), partTwo(input)
}

var shapes = [][][]int{
	{
		{1, 1, 1, 1},
	},
	{
		{0, 1, 0},
		{1, 1, 1},
		{0, 1, 0},
	},
	{
		{1, 1, 1},
		{0, 0, 1},
		{0, 0, 1},
	},
	{
		{1},
		{1},
		{1},
		{1},
	},
	{
		{1, 1},
		{1, 1},
	},
}

type state [][]bool

func (s state) FilledRows() int {
	for y := len(s) - 1; y >= 0; y-- {
		for _, f := range s[y] {
			if f {
				return y + 1
			}
		}
	}
	return 0
}

func (s *state) EnsureHeight(width int, height int) {
	for len(*s) < height {
		*s = append(*s, make([]bool, width))
	}
}

func (s *state) HitsShape(shape [][]int, width int, lx int, by int) bool {
	for y, row := range shape {
		for x, f := range row {
			if f == 1 && (by+y < 0 || lx+x < 0 || lx+x >= width || (*s)[by+y][lx+x]) {
				return true
			}
		}
	}
	return false
}

func (s *state) ApplyShape(shape [][]int, lx int, by int) {
	for y, row := range shape {
		for x, f := range row {
			if f == 1 {
				(*s)[by+y][lx+x] = true
			}
		}
	}
}

func (s *state) TryToMoveShape(shape [][]int, width int, lx int, ty int, dx int, dy int) (int, int, bool) {
	if !s.HitsShape(shape, width, lx+dx, ty+dy) {
		return lx + dx, ty + dy, true
	}
	return lx, ty, false
}

func Simulate(width int, dx int, dy int, jets string, cb func(i int, j int, state state) bool) {
	state := state([][]bool{})

	i, j := 0, 0

	for cb(i, j, state) {
		topY := state.FilledRows()
		shape := shapes[i%len(shapes)]

		state.EnsureHeight(width, topY+dy+len(shape))

		cx, cy := dx, topY+dy

		for {
			jet := jets[j%len(jets)]
			j++

			if jet == '<' {
				cx, cy, _ = state.TryToMoveShape(shape, width, cx, cy, -1, 0)
			} else {
				cx, cy, _ = state.TryToMoveShape(shape, width, cx, cy, +1, 0)
			}

			var moved bool
			cx, cy, moved = state.TryToMoveShape(shape, width, cx, cy, 0, -1)

			if !moved {
				state.ApplyShape(shape, cx, cy)
				break
			}
		}

		i++
	}
}

func partOne(insn string) int {
	var ans int
	Simulate(7, 2, 3, insn, func(i, j int, state state) bool {
		if i == 2022 {
			ans = state.FilledRows()
			return false
		}

		return true
	})

	return ans
}

type memoRow struct {
	i      int
	height int
}

func partTwo(insn string) int {
	heights := make([]int, 0, 200)
	var prefix, cycle memoRow

	memo := map[string]memoRow{}
	Simulate(7, 2, 3, insn, func(i, j int, state state) bool {
		topFilled := state.FilledRows()
		heights = append(heights, topFilled)

		key := ""
		for y := topFilled - 1; y >= topFilled-50; y-- {
			if y >= 0 {
				key += fmt.Sprintf("%v;", state[y])
			} else {
				key += "bedrock;"
			}
		}

		key += fmt.Sprintf("###;%d;%d", i%len(shapes), j%len(insn))

		mr, isCycle := memo[key]

		if isCycle {
			prefix = mr
			cycle = memoRow{i - mr.i, topFilled - mr.height}
			return false
		}

		memo[key] = memoRow{i, topFilled}
		return true
	})

	cnt := 1000000000000
	rep := cnt - prefix.i
	div, rem := rep/cycle.i, rep%cycle.i
	ans := prefix.height + div*cycle.height
	ans += heights[len(heights)-cycle.i+rem-1] - heights[len(heights)-cycle.i-1]

	return ans
}
