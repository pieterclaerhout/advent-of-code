package day24

import (
	"bufio"
	_ "embed"
	"strings"

	"golang.org/x/exp/slog"
)

//go:embed input.txt
var input string

type Command struct {
}

func (c *Command) Execute() {

	lines := c.parse()

	state := newState()
	state.valley[0] = make([][]BlizzardSet, len(lines))
	for i, line := range lines {
		state.valley[0][i] = make([]BlizzardSet, len(line))
		for j := range line {
			if line[j] == '.' {
				continue
			}
			state.valley[0][i][j] = BlizzardSet(parseBlizzard(line[j]))
		}
	}
	for i := 0; i < maxTime; i++ {
		state.computeNext(i)
	}

	start := SpaceTime{0, 0, 1}
	goal := Pair{26, 120}
	t1 := state.solve(start, goal)

	slog.Info("Part 1", slog.Any("result", t1))

	start, goal = SpaceTime{t1, goal.i, goal.j}, Pair{start.i, start.j}
	t2 := state.solve(start, goal)

	start, goal = SpaceTime{t2, goal.i, goal.j}, Pair{start.i, start.j}
	t3 := state.solve(start, goal)

	slog.Info("Part 1", slog.Any("result", t3))
}

func (c *Command) parse() []string {
	lines := []string{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}
