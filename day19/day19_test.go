package day19_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day19"
	"github.com/stretchr/testify/assert"
)

const input = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

func TestDay4(t *testing.T) {
	cmd := day19.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 33, result1)
	assert.Equal(t, 62, result2)
}
