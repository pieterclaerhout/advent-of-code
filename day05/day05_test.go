package day05_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day05"
	"github.com/stretchr/testify/assert"
)

const input = `    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestDay4(t *testing.T) {
	cmd := day05.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, "CMZ", result1)
	assert.Equal(t, "MCD", result2)
}
