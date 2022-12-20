package day19

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"

	"golang.org/x/exp/slog"
)

//go:embed input.txt
var input string

type Command struct {
}

func (c *Command) Execute() {
	c.part1()
	c.part2()
}

func (c *Command) part1() {
	blueprints := c.parse()

	solution := 0
	for id, bprint := range blueprints {
		st := NewState(bprint)
		geodes := st.Dfs(0, Resources{}, Resources{Ores: 1}, Resources{})

		qualityLevel := (id + 1) * geodes
		solution += qualityLevel
	}

	slog.Info("Part 1", slog.Any("result", solution))
}

func (c *Command) part2() {
	blueprints := c.parse()
	if len(blueprints) > 3 {
		blueprints = blueprints[0:3]
	}

	solution := 1
	for _, bprint := range blueprints {
		st := NewState(bprint)
		st.MaxMinutes = 32
		geodes := st.Dfs(0, Resources{}, Resources{Ores: 1}, Resources{})

		solution *= geodes
	}

	slog.Info("Part 2", slog.Any("result", solution))
}

func (c *Command) parse() []Blueprint {

	blueprints := []Blueprint{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {

		var bID, oreOre, clayOre, obsidianOre, obsidianClay, geodeOre, geodeObsidian int
		fmt.Sscanf(sc.Text(), "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &bID, &oreOre, &clayOre, &obsidianOre, &obsidianClay, &geodeOre, &geodeObsidian)

		blueprints = append(blueprints, Blueprint{
			{oreOre, 0, 0, 0},
			{clayOre, 0, 0, 0},
			{obsidianOre, obsidianClay, 0, 0},
			{geodeOre, 0, geodeObsidian, 0},
		})
	}

	return blueprints
}
