package day24_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day24"
	"github.com/stretchr/testify/assert"
)

const input = `#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#`

func Test_Command(t *testing.T) {
	cmd := day24.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 10, result1)
	assert.Equal(t, 30, result2)
}
