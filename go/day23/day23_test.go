package day23_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day23"
	"github.com/stretchr/testify/assert"
)

const input = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

func Test_Command(t *testing.T) {
	cmd := day23.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 110, result1)
	assert.Equal(t, 20, result2)
}
