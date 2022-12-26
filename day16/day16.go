package day16

import (
	"fmt"
	"strings"
)

type Command struct{}

func (cmd *Command) Execute(input string) (any, any) {
	valves := cmd.parse(input)
	return cmd.part1(valves), cmd.part2(valves)
}

func (c *Command) part1(valves Valves) int {
	nonEmpty := valves.NonEmpty()
	distances := valves.ValvesDistances(nonEmpty)

	cache := Cache{
		TotalMinutes:     30,
		StateMaxPressure: map[string]int{},
		Distances:        distances,
		ValvesFlow:       valves.FlowRates,
	}

	return cache.Dfs([]string{"AA"}, "")
}

func (c *Command) part2(valves Valves) int {
	nonEmpty := valves.NonEmpty()
	distances := valves.ValvesDistances(nonEmpty)

	cache := Cache{
		TotalMinutes:     26,
		StateMaxPressure: map[string]int{},
		Distances:        distances,
		ValvesFlow:       valves.FlowRates,
	}

	cache.Dfs([]string{"AA"}, "")

	return cache.MaxExclusivePair()
}

func (c *Command) parse(input string) Valves {
	valves := Valves{
		FlowRates:   map[string]int{},
		Connections: map[string][]string{},
	}

	for _, line := range strings.Split(input, "\n") {
		var label string
		var flowRate int
		var connections []string

		lineParts := strings.Split(line, "; tunnels lead to valves ")
		if len(lineParts) == 1 {
			lineParts = strings.Split(line, "; tunnel leads to valve ")
		}

		fmt.Sscanf(lineParts[0], "Valve %s has flow rate=%d", &label, &flowRate)

		connections = strings.Split(lineParts[1], ", ")

		valves.FlowRates[label] = flowRate
		valves.Connections[label] = connections
	}

	return valves
}
