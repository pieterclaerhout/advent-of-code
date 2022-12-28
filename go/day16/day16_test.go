package day16_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day16"
	"github.com/stretchr/testify/assert"
)

const input = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`

func Test_Command(t *testing.T) {
	cmd := day16.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 1651, result1)
	assert.Equal(t, 1707, result2)
}
