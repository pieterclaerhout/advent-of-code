package day19

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Command struct {
}

func (c *Command) Execute(input string) (any, any) {
	return c.part1(input), c.part2(input)
}

func (c *Command) part1(input string) int {
	blueprints := c.parse(input)

	solution := 0
	for id, bprint := range blueprints {
		st := NewState(bprint)
		geodes := st.Dfs(0, Resources{}, Resources{Ores: 1}, Resources{})

		qualityLevel := (id + 1) * geodes
		solution += qualityLevel
	}

	return solution
}

func (c *Command) part2(input string) int {
	blueprints := c.parse(input)
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
	return solution
}

func (c *Command) parse(input string) []Blueprint {

	blueprints := []Blueprint{}

	for _, line := range strings.Split(input, "\n") {

		var bID, oreOre, clayOre, obsidianOre, obsidianClay, geodeOre, geodeObsidian int
		fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &bID, &oreOre, &clayOre, &obsidianOre, &obsidianClay, &geodeOre, &geodeObsidian)

		blueprints = append(blueprints, Blueprint{
			{oreOre, 0, 0, 0},
			{clayOre, 0, 0, 0},
			{obsidianOre, obsidianClay, 0, 0},
			{geodeOre, 0, geodeObsidian, 0},
		})
	}

	return blueprints
}
