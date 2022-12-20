package day16

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
	valves := c.parse()
	nonEmpty := valves.NonEmpty()
	distances := valves.ValvesDistances(nonEmpty)

	cache := Cache{
		TotalMinutes:     30,
		StateMaxPressure: map[string]int{},
		Distances:        distances,
		ValvesFlow:       valves.FlowRates,
	}

	slog.Info("Part 1", slog.Any("result", cache.DFS([]string{"AA"}, "")))
}

func (c *Command) part2() {
	valves := c.parse()
	nonEmpty := valves.NonEmpty()
	distances := valves.ValvesDistances(nonEmpty)

	cache := Cache{
		TotalMinutes:     26,
		StateMaxPressure: map[string]int{},
		Distances:        distances,
		ValvesFlow:       valves.FlowRates,
	}

	cache.DFS([]string{"AA"}, "")

	slog.Info("Part 2", slog.Any("result", cache.MaxExclusivePair()))
}

func (c *Command) parse() Valves {
	valves := Valves{
		FlowRates:   map[string]int{},
		Connections: map[string][]string{},
	}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		var label string
		var flowRate int
		var connections []string

		lineParts := strings.Split(sc.Text(), "; tunnels lead to valves ")
		if len(lineParts) == 1 {
			lineParts = strings.Split(sc.Text(), "; tunnel leads to valve ")
		}

		fmt.Sscanf(lineParts[0], "Valve %s has flow rate=%d", &label, &flowRate)

		connections = strings.Split(lineParts[1], ", ")

		valves.FlowRates[label] = flowRate
		valves.Connections[label] = connections
	}

	return valves
}
