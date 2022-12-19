package day03

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
	c.part1()
	c.part2()
}

func (c *Command) part1() {
	ruckSacks := c.parse()

	var score int

	for _, ruckSack := range ruckSacks {
		score += ruckSack.Priorities()
	}

	slog.Info("Part 1", slog.Any("score", score))
}

func (c *Command) part2() {

	groups := [][]RuckSack{}

	subgroup := []RuckSack{}
	for _, ruckSack := range c.parse() {
		subgroup = append(subgroup, ruckSack)
		if len(subgroup) == 3 {
			groups = append(groups, subgroup)
			subgroup = []RuckSack{}
		}
	}

	var score int
	for _, group := range groups {
		common := intersect(group[0].All(), group[1].All())
		common = intersect(common, group[2].All())
		score += scoreForCharacter(common[0])
	}

	slog.Info("Part 2", slog.Any("score", score))
}

func (c *Command) parse() []RuckSack {
	result := []RuckSack{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		result = append(result, NewRuckSack(sc.Text()))
	}

	return result
}
